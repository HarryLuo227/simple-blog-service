package model

import (
	"fmt"
	"time"

	"github.com/HarryLuo227/simple-blog-service/global"
	"github.com/HarryLuo227/simple-blog-service/pkg/setting"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 所有資料表裡的共同欄位
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedOn  uint32 `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn uint32 `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	IsDel      uint8  `json:"is_del"`
}

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%v&loc=Local"
	dsn := fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	db, err := gorm.Open(databaseSetting.DBType, dsn)
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)
	otgorm.AddGormCallbacks(db)
	return db, nil
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifiedTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifiedTimeField.IsBlank {
				_ = modifiedTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"Update %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}

	return ""
}

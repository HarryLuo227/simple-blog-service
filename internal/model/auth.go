package model

import "github.com/jinzhu/gorm"

type Auth struct {
	User     string `json:"user"`
	Password string `json:"password"`
	*Model
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("user = ? AND password = ? AND is_del = ?", a.User, a.Password, 0)
	if err := db.First(&auth).Error; err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}

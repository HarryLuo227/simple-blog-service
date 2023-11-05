package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	ArticleID uint32 `json:"article_id`
	TagID     uint32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}

	return articleTag, nil
}

func (a ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("tag_id = ? AND is_del = ?", a.TagID, 0).Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
}

func (a ArticleTag) ListByAIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id IN (?) AND is_del = ?", articleIDs, 0).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articleTags, nil
}

// Insert `ArticleTag` value into the table.
func (at ArticleTag) Create(db *gorm.DB) error {
	if err := db.Create(&at).Error; err != nil {
		return err
	}

	return nil
}

// Delete `ArticleTag` rows by `Article` ID.
func (a ArticleTag) Delete(db *gorm.DB) error {
	if err := db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Delete(&a).Limit(1).Error; err != nil {
		return err
	}

	return nil
}

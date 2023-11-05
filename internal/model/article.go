package model

import "github.com/jinzhu/gorm"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type ArticleRow struct {
	ArticleID     uint32
	TagID         uint32
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
}

func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article_title", "ar.desc AS article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImageUrl, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}

		articles = append(articles, r)
	}

	return articles, nil
}

func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Get `Article` row  by its `id` columns.
func (a Article) Get(db *gorm.DB) (*Article, error) {
	var article Article

	err := db.Where("id = ? and state = ? and is_del = ?", a.ID, a.State, 0).First(&article).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

// Insert `Article` value into the table.
func (a Article) Create(db *gorm.DB) error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}

	return nil
}

// Update the columns of `Article` row in the table.
func (a Article) Update(db *gorm.DB, values map[string]interface{}) error {
	err := db.Model(&Article{}).Where("id = ? and is_del = ?", a.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete `Article` row.
func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", a.Model.ID, 0).Delete(&a).Error
}

// Verify if `Article` value exists.
func (a Article) ArticleExists(db *gorm.DB) (*Article, bool) {
	var target Article

	err := db.Where("title = ? AND created_by = ?", a.Title, a.CreatedBy).First(&target).Error
	if err != nil {
		return nil, false
	}

	return &target, true
}

// Get `Article` row by its `title` and `createdBy` columns.
func (a Article) QueryArticle(db *gorm.DB) (*Article, error) {
	err := db.Where("title = ? AND created_by = ?", a.Title, a.CreatedBy).First(&a).Error
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// Get `Tag` rows which associate to the `Article`.
func (a Article) GetTagsByAID(db *gorm.DB) ([]Tag, error) {
	var tags []Tag

	// SELECT t.* FROM blog_article_tag AS at LEFT JOIN blog_article AS a ON at.article_id = a.id LEFT JOIN blog_tag as t ON at.tag_id = t.id WHERE a.id = ? AND a.state = 0 AND a.is_del = 0 AND t.is_del = 0;
	db.Select("t.*").Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS a ON at.article_id = a.id").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Where("a.id = ? AND a.state = ? AND a.is_del = ? AND t.is_del = ?", a.ID, a.State, 0, 0).Find(&tags)

	if err := db.Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// Get `ArticleTag` values which associate to the `Article`.
func (a Article) GetArticleTagsByAID(db *gorm.DB) ([]ArticleTag, error) {
	var articleTags []ArticleTag

	// SELECT at.* FROM blog_article AS a LEFT JOIN blog_article_tag AS at ON a.id = at.article_id WHERE a.id = ? AND a.state = 1 AND a.is_del = 0;
	db.Select("at.*").Table(Article{}.TableName()+" AS a").
		Joins("LEFT JOIN `"+ArticleTag{}.TableName()+"` AS at ON a.id = at.article_id").
		Where("a.id = ? AND a.state = ? AND a.is_del = ?", a.ID, a.State, 0).Find(&articleTags)

	if err := db.Error; err != nil {
		return nil, err
	}
	return articleTags, nil
}

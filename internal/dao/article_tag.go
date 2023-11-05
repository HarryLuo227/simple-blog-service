package dao

import "github.com/HarryLuo227/simple-blog-service/internal/model"

func (d *Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.GetByAID(d.engine)
}

func (d *Dao) GetArticleTagListByTID(tagID uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagID: tagID}
	return articleTag.ListByTID(d.engine)
}

func (d *Dao) GetArticleTagListByAIDs(articleIDs []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByAIDs(d.engine, articleIDs)
}

func (d *Dao) CreateArticleTag(articleID, tagID uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		ArticleID: articleID,
		TagID:     tagID,
	}

	return articleTag.Create(d.engine)
}

func (d *Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{ArticleID: articleID}

	return articleTag.Delete(d.engine)
}

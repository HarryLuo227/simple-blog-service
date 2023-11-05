package dao

import (
	"errors"

	"github.com/HarryLuo227/simple-blog-service/internal/model"
	"github.com/HarryLuo227/simple-blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	ID            uint32 `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) CountArticleListByTagID(id uint32, state uint8) (int, error) {
	article := model.Article{State: state}
	return article.CountByTagID(d.engine, id)
}

func (d *Dao) GetArticleListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}
	return article.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}

func (d *Dao) GetArticle(id uint32, state uint8) (*model.Article, error) {
	article := model.Article{
		Model: &model.Model{ID: id},
		State: state,
	}

	return article.Get(d.engine)
}

func (d *Dao) CreateArticle(title, desc, content, coverImageUrl string, state uint8, createdBy string) error {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		State:         state,
		Model:         &model.Model{CreatedBy: createdBy},
	}

	if target, isExist := article.ArticleExists(d.engine); isExist {
		values := map[string]interface{}{}
		if target.IsDel == 1 {
			values["is_del"] = 0
			values["state"] = 1
			values["modified_by"] = article.CreatedBy
			if article.Desc != "" {
				values["desc"] = article.Desc
			}
			if article.Content != "" {
				values["content"] = article.Content
			}
			if article.CoverImageUrl != "" {
				values["cover_image_url"] = article.CoverImageUrl
			}
		} else if target.State == 0 {
			values["state"] = 1
			values["modified_by"] = article.CreatedBy
			if article.Desc != "" {
				values["desc"] = article.Desc
			}
			if article.Content != "" {
				values["content"] = article.Content
			}
			if article.CoverImageUrl != "" {
				values["cover_image_url"] = article.CoverImageUrl
			}
		}

		return target.Update(d.engine, values)
	} else {
		return article.Create(d.engine)
	}

}

func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{
		Model: &model.Model{ID: param.ID},
		State: param.State,
	}

	dest, gErr := article.Get(d.engine)
	if gErr != nil {
		return gErr
	}

	query := model.Article{
		Model: &model.Model{CreatedBy: dest.CreatedBy},
		Title: param.Title,
	}
	target, qErr := query.QueryArticle(d.engine)
	if qErr != nil && qErr != gorm.ErrRecordNotFound {
		return qErr
	}
	if target != nil && target.Title == param.Title && target.CreatedBy == dest.CreatedBy {
		return errors.New("Article title after update is already exists same author.")
	}

	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}

	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}

	return article.Delete(d.engine)
}

func (d *Dao) QueryArticle(title, createdBy string) (*model.Article, error) {
	article := model.Article{
		Title: title,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return article.QueryArticle(d.engine)
}

func (d *Dao) GetTagListAssociateToArticle(id uint32, state uint8) ([]model.Tag, error) {
	article := model.Article{
		Model: &model.Model{ID: id},
		State: state,
	}

	return article.GetTagsByAID(d.engine)
}

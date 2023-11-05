package service

import (
	"github.com/HarryLuo227/simple-blog-service/internal/dao"
	"github.com/HarryLuo227/simple-blog-service/internal/model"
	"github.com/HarryLuo227/simple-blog-service/pkg/app"
)

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagID         []uint32 `form:"tag_id" binding:"required,dive,gte=1"`
	Title         string   `form:"title" binding:"required,min=1,max=100"`
	Desc          string   `form:"desc" binding:"max=255"`
	Content       string   `form:"content" binding:"required,min=1,max=4294967295"`
	CoverImageUrl string   `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string   `form:"created_by" binding:"required,min=1,max=100"`
	State         uint8    `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32   `form:"id" binding:"required,gte=1"`
	TagID         []uint32 `form:"tag_id" binding:"omitempty,gte=1"`
	Title         string   `form:"title" binding:"omitempty,min=1,max=100"`
	Desc          string   `form:"desc" binding:"max=255"`
	Content       string   `form:"content" binding:"omitempty,min=1,max=4294967295"`
	CoverImageUrl string   `form:"cover_image_url" binding:"omitempty,url"`
	ModifiedBy    string   `form:"modified_by" binding:"required,min=1,max=100"`
	State         uint8    `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type QueryArticleIDRequest struct {
	Title     string `form:"title" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=1,max=100"`
}

type ArticleWithTags struct {
	Article *model.Article
	Tags    *[]model.Tag
}

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*Article, int, error) {
	articleCount, err := svc.dao.CountArticleListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}

	articles, err := svc.dao.GetArticleListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			ID:            article.ArticleID,
			Title:         article.ArticleTitle,
			Desc:          article.ArticleDesc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			Tag:           &model.Tag{Model: &model.Model{ID: article.TagID}, Name: article.TagName},
		})
	}

	return articleList, articleCount, nil
}

func (svc *Service) GetArticle(param *ArticleRequest) (*ArticleWithTags, error) {
	article, err := svc.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}
	respArticleTags, err := svc.dao.GetTagListAssociateToArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	resp := ArticleWithTags{
		Article: article,
		Tags:    &respArticleTags,
	}

	return &resp, nil
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) (*ArticleWithTags, error) {
	for _, v := range param.TagID {
		if _, err := svc.dao.GetTag(v, 1); err != nil {
			return nil, err
		}
	}

	err := svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.CreatedBy)
	if err != nil {
		return nil, err
	}

	article, err := svc.dao.QueryArticle(param.Title, param.CreatedBy)
	if err != nil {
		return nil, err
	}
	if err := svc.dao.DeleteArticleTag(article.ID); err != nil {
		return nil, err
	}
	for _, v := range param.TagID {
		err := svc.dao.CreateArticleTag(article.ID, v, param.CreatedBy)
		if err != nil {
			return nil, err
		}
	}

	respArticleTags, err := svc.dao.GetTagListAssociateToArticle(article.ID, param.State)
	if err != nil {
		return nil, err
	}

	resp := ArticleWithTags{
		Article: article,
		Tags:    &respArticleTags,
	}

	return &resp, nil
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) (*ArticleWithTags, error) {
	err := svc.dao.UpdateArticle(&dao.Article{
		ID:            param.ID,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		ModifiedBy:    param.ModifiedBy,
	})
	if err != nil {
		return nil, err
	}

	respArticle, err := svc.dao.GetArticle(param.ID, 1)
	if err != nil {
		return nil, err
	}

	if param.TagID == nil {
		respArticleTags, err := svc.dao.GetTagListAssociateToArticle(param.ID, 1)
		if err != nil {
			return nil, err
		}
		resp := ArticleWithTags{
			Article: respArticle,
			Tags:    &respArticleTags,
		}
		return &resp, nil
	}

	if err := svc.dao.DeleteArticleTag(param.ID); err != nil {
		return nil, err
	}
	for _, v := range param.TagID {
		if err := svc.dao.CreateArticleTag(param.ID, v, param.ModifiedBy); err != nil {
			return nil, err
		}
	}
	respArticleTags, err := svc.dao.GetTagListAssociateToArticle(param.ID, 1)
	if err != nil {
		return nil, err
	}
	resp := ArticleWithTags{
		Article: respArticle,
		Tags:    &respArticleTags,
	}

	return &resp, nil
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	err := svc.dao.DeleteArticle(param.ID)
	if err != nil {
		return err
	}

	err = svc.dao.DeleteArticleTag(param.ID)
	if err != nil {
		return err
	}

	return nil
}

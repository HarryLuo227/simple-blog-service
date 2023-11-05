package service

import (
	"github.com/HarryLuo227/simple-blog-service/internal/model"
	"github.com/HarryLuo227/simple-blog-service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetTagRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=1,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"omitempty,min=1,max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type GetTagIDRequest struct {
	ID        uint32 `form:"id" binding:"omitempty, gte=1"`
	Name      string `form:"name" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=1,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) GetTag(param *GetTagRequest) (*model.Tag, error) {
	return svc.dao.GetTag(param.ID, param.State)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}

func (svc *Service) QueryTag(param *GetTagIDRequest) (*model.Tag, error) {
	return svc.dao.QueryTag(param.Name, param.CreatedBy)
}

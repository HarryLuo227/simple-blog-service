package dao

import (
	"github.com/HarryLuo227/simple-blog-service/internal/model"
	"github.com/HarryLuo227/simple-blog-service/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}

	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)

	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetTagListByIDs(ids []uint32, state uint8) ([]*model.Tag, error) {
	tag := model.Tag{State: state}
	return tag.ListByIDs(d.engine, ids)
}

func (d *Dao) GetTag(id uint32, state uint8) (*model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id}, State: state}
	return tag.Get(d.engine)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	if target, isExist := tag.TagExists(d.engine); isExist {
		values := map[string]interface{}{}
		if target.IsDel == 1 {
			values["is_del"] = 0
			values["state"] = 1
			values["modified_by"] = tag.CreatedBy
		} else if target.State == 0 {
			values["state"] = 1
			values["modified_by"] = tag.CreatedBy
		}

		return target.Update(d.engine, values)
	} else {
		return tag.Create(d.engine)
	}
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}

	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}

	return tag.Delete(d.engine)
}

func (d *Dao) QueryTag(name, createdBy string) (*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return tag.QueryTag(d.engine)
}

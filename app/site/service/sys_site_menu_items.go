package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/site/models"
	"go-admin/app/site/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysSiteMenuItems struct {
	service.Service
}

// GetPage 获取SysSiteMenuItems列表
func (e *SysSiteMenuItems) GetPage(c *dto.SysSiteMenuItemsGetPageReq, p *actions.DataPermission, list *[]models.SysSiteMenuItems, count *int64) error {
	var err error
	var data models.SysSiteMenuItems

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysSiteMenuItemsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysSiteMenuItems对象
func (e *SysSiteMenuItems) Get(d *dto.SysSiteMenuItemsGetReq, p *actions.DataPermission, model *models.SysSiteMenuItems) error {
	var data models.SysSiteMenuItems

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysSiteMenuItems error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysSiteMenuItems对象
func (e *SysSiteMenuItems) Insert(c *dto.SysSiteMenuItemsInsertReq) error {
    var err error
    var data models.SysSiteMenuItems
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysSiteMenuItemsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysSiteMenuItems对象
func (e *SysSiteMenuItems) Update(c *dto.SysSiteMenuItemsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.SysSiteMenuItems{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SysSiteMenuItemsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除SysSiteMenuItems
func (e *SysSiteMenuItems) Remove(d *dto.SysSiteMenuItemsDeleteReq, p *actions.DataPermission) error {
	var data models.SysSiteMenuItems

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSysSiteMenuItems error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

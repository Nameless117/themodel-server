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

type SysSiteSubscriptionLinks struct {
	service.Service
}

// GetPage 获取SysSiteSubscriptionLinks列表
func (e *SysSiteSubscriptionLinks) GetPage(c *dto.SysSiteSubscriptionLinksGetPageReq, p *actions.DataPermission, list *[]models.SysSiteSubscriptionLinks, count *int64) error {
	var err error
	var data models.SysSiteSubscriptionLinks

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysSiteSubscriptionLinksService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysSiteSubscriptionLinks对象
func (e *SysSiteSubscriptionLinks) Get(d *dto.SysSiteSubscriptionLinksGetReq, p *actions.DataPermission, model *models.SysSiteSubscriptionLinks) error {
	var data models.SysSiteSubscriptionLinks

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysSiteSubscriptionLinks error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysSiteSubscriptionLinks对象
func (e *SysSiteSubscriptionLinks) Insert(c *dto.SysSiteSubscriptionLinksInsertReq) error {
    var err error
    var data models.SysSiteSubscriptionLinks
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysSiteSubscriptionLinksService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysSiteSubscriptionLinks对象
func (e *SysSiteSubscriptionLinks) Update(c *dto.SysSiteSubscriptionLinksUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.SysSiteSubscriptionLinks{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SysSiteSubscriptionLinksService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除SysSiteSubscriptionLinks
func (e *SysSiteSubscriptionLinks) Remove(d *dto.SysSiteSubscriptionLinksDeleteReq, p *actions.DataPermission) error {
	var data models.SysSiteSubscriptionLinks

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSysSiteSubscriptionLinks error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

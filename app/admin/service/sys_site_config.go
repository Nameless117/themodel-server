package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysSiteConfig struct {
	service.Service
}

// GetPage 获取SysSiteConfig列表
func (e *SysSiteConfig) GetPage(c *dto.SysSiteConfigGetPageReq, p *actions.DataPermission, list *[]models.SysSiteConfig, count *int64) error {
	var err error
	var data models.SysSiteConfig

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysSiteConfigService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysSiteConfig对象
func (e *SysSiteConfig) Get(d *dto.SysSiteConfigGetReq, p *actions.DataPermission, model *models.SysSiteConfig) error {
	var data models.SysSiteConfig

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysSiteConfig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysSiteConfig对象
func (e *SysSiteConfig) Insert(c *dto.SysSiteConfigInsertReq) error {
    var err error
    var data models.SysSiteConfig
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysSiteConfigService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysSiteConfig对象
func (e *SysSiteConfig) Update(c *dto.SysSiteConfigUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.SysSiteConfig{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SysSiteConfigService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除SysSiteConfig
func (e *SysSiteConfig) Remove(d *dto.SysSiteConfigDeleteReq, p *actions.DataPermission) error {
	var data models.SysSiteConfig

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSysSiteConfig error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

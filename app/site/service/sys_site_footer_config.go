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

type SysSiteFooterConfig struct {
	service.Service
}

// GetPage 获取SysSiteFooterConfig列表
func (e *SysSiteFooterConfig) GetPage(c *dto.SysSiteFooterConfigGetPageReq, p *actions.DataPermission, list *[]models.SysSiteFooterConfig, count *int64) error {
	var err error
	var data models.SysSiteFooterConfig

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysSiteFooterConfigService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysSiteFooterConfig对象
func (e *SysSiteFooterConfig) Get(d *dto.SysSiteFooterConfigGetReq, p *actions.DataPermission, model *models.SysSiteFooterConfig) error {
	var data models.SysSiteFooterConfig

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysSiteFooterConfig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysSiteFooterConfig对象
func (e *SysSiteFooterConfig) Insert(c *dto.SysSiteFooterConfigInsertReq) error {
    var err error
    var data models.SysSiteFooterConfig
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysSiteFooterConfigService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysSiteFooterConfig对象
func (e *SysSiteFooterConfig) Update(c *dto.SysSiteFooterConfigUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.SysSiteFooterConfig{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SysSiteFooterConfigService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除SysSiteFooterConfig
func (e *SysSiteFooterConfig) Remove(d *dto.SysSiteFooterConfigDeleteReq, p *actions.DataPermission) error {
	var data models.SysSiteFooterConfig

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSysSiteFooterConfig error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

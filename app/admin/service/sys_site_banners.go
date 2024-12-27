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

type SysSiteBanners struct {
	service.Service
}

// GetPage 获取SysSiteBanners列表
func (e *SysSiteBanners) GetPage(c *dto.SysSiteBannersGetPageReq, p *actions.DataPermission, list *[]models.SysSiteBanners, count *int64) error {
	var err error
	var data models.SysSiteBanners

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysSiteBannersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysSiteBanners对象
func (e *SysSiteBanners) Get(d *dto.SysSiteBannersGetReq, p *actions.DataPermission, model *models.SysSiteBanners) error {
	var data models.SysSiteBanners

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysSiteBanners error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysSiteBanners对象
func (e *SysSiteBanners) Insert(c *dto.SysSiteBannersInsertReq) error {
    var err error
    var data models.SysSiteBanners
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysSiteBannersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysSiteBanners对象
func (e *SysSiteBanners) Update(c *dto.SysSiteBannersUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.SysSiteBanners{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SysSiteBannersService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除SysSiteBanners
func (e *SysSiteBanners) Remove(d *dto.SysSiteBannersDeleteReq, p *actions.DataPermission) error {
	var data models.SysSiteBanners

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSysSiteBanners error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

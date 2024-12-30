package models

import (
	"database/sql/driver"
	"errors"
	"github.com/goccy/go-json"
	"go-admin/common/models"
)

type SysSiteContent struct {
	models.Model

	Type        string         `json:"type" gorm:"type:enum('home','about','audio');comment:内容类型：首页/关于/有声读物"`
	Title       string         `json:"title" gorm:"type:varchar(100);comment:内容名称"`
	Content     string         `json:"content" gorm:"type:text;comment:富文本内容"`
	SubContent  SubContentJSON `json:"subContent" gorm:"type:text;comment:附加内容"`
	HtmlContent string         `json:"htmlContent" gorm:"type:text;htmlContent:附加内容"`
	models.ModelTime
	models.ControlBy
}

// 自定义 JSON 类型，用于处理 GORM JSON 字段
type SubContentJSON []Quote

type Quote struct {
	Saying string `json:"saying"`
	Author string `json:"author"`
}

func (j *SubContentJSON) Scan(value interface{}) error {
	if value == nil {
		*j = SubContentJSON{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, j)
}

func (j SubContentJSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return "[]", nil
	}
	return json.Marshal(j)
}

func (SysSiteContent) TableName() string {
	return "sys_site_content"
}

func (e *SysSiteContent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysSiteContent) GetId() interface{} {
	return e.Id
}

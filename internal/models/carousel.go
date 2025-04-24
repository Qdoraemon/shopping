package models

import (
	"time"
)

type Carousel struct {
	ID          int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"column:title;type:varchar(50);not null;comment:标题" json:"title"`
	Order       int       `gorm:"column:order;type:int;not null;default:0;comment:排序" json:"order"`
	ImageURL    string    `gorm:"column:image_url;type:varchar(255);not null;comment:图片URL地址"  json:"imageUrl"`
	RedirectURL string    `gorm:"column:redirect_url;type:varchar(255);not null;comment:图片URL地址" json:"redirectUrl"`
	IsEnabled   bool      `gorm:"column:is_enabled;type:tinyint(1);not null;default:0;comment:是否启用 0：未启用 1：启用" json:"isEnabled"`
	IsDeleted   bool      `gorm:"column:is_deleted;type:tinyint(1);not null;default:1;comment:是否删除 0：删除 1：正常" json:"createTime"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null;comment:创建时间" json:"updateTime"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime;not null;comment:更新时间" json:"isDeleted"`
}

func (Carousel) TableName() string {
	return "carousel"
}

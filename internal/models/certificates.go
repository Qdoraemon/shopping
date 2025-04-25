package models

import "time"

type Certificates struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(50);not null;comment:证书名称" json:"name"`
	Description string    `gorm:"type:varchar(255);comment:证书描述" json:"description"`
	ImageURL    string    `gorm:"type:varchar(255);comment:证书图片" json:"imageUrl"`
	IsDeleted   bool      `gorm:"type:tinyint(1);not null;default:1;comment:是否删除 0：删除 1：正常" json:"is_deleted"`
	IsAvailable bool      `gorm:"type:tinyint(1);not null;default:1;comment:上下架状态 0:下架 1:上架" json:"is_available"`
	CreateTime  time.Time `gorm:"not null;comment:创建时间" json:"create_time"`
	UpdateTime  time.Time `gorm:"not null;comment:更新时间" json:"update_time"`
}

func (Certificates) TableName() string {
	return "certificates"
}

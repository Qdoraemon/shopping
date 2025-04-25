package models

import "time"

type Brands struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(50);not null;comment:品牌名称" json:"name"`
	Description string    `gorm:"type:varchar(255);comment:品牌描述" json:"description"`
	Logo        string    `gorm:"type:varchar(255);comment:品牌logo" json:"logo"`
	UpdateTime  time.Time `gorm:"not null;comment:更新时间" json:"update_time"`
	CreateTime  time.Time `gorm:"not null;comment:创建时间" json:"create_time"`
	IsDeleted   bool      `gorm:"type:tinyint(1);not null;default:1;comment:是否删除 0：删除 1：正常" json:"is_deleted"`
	IsEnabled   bool      `gorm:"type:tinyint(1);not null;default:1;comment:是否启用 0：未启用 1：启用" json:"is_enabled"`
}

func (Brands) TableName() string {
	return "brands"
}

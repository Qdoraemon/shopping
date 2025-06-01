package models

import "time"

type Categories struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(50);not null;comment:类型名称" json:"name"`
	NameEn      string    `gorm:"type:varchar(50);not null;comment:类型名称" json:"name_en"`
	Description string    `gorm:"type:varchar(255);comment:类型描述" json:"description"`
	UpdateTime  time.Time `gorm:"not null;comment:更新时间" json:"update_time"`
	CreateTime  time.Time `gorm:"not null;comment:创建时间" json:"create_time"`
	IsDeleted   int       `gorm:"type:tinyint(1);not null;default:1;comment:是否删除 0：删除 1：正常" json:"is_deleted"`
	IsEnabled   int       `gorm:"type:tinyint(1);not null;default:1;comment:是否启用 0：未启用 1：启用" json:"is_enabled"`
}

// SearchBrandsRequest 定义搜索请求参数
type SearchCategoriesRequest struct {
	Name      string `json:"name"`       // 分類名称
	IsEnabled string `json:"is_enabled"` // 分類状态（启用/禁用）
	Page      int    `json:"page"`       // 当前页
	PageSize  int    `json:"page_size"`  // 每页大小
}

type SearchCategoriesResponse struct {
	Categories []*Categories `json:"categories"` // 分類列表
	Total      int64         `json:"total"`      // 总记录数
}

func (Categories) TableName() string {
	return "categories"
}

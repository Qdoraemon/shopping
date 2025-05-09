package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// StringSlice 是一个自定义类型，用于处理字符串切片
type StringSlice []string

// Scan 实现了 sql.Scanner 接口，用于将数据库中的数据转换为 StringSlice
func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*s = []string{}
		return nil
	}
	// 将数据库中的 []byte 转换为字符串
	str, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to convert %v to []byte", value)
	}
	// 将字符串按逗号分隔并转换为 []string
	*s = strings.Split(string(str), ",")
	return nil
}

// Value 实现了 driver.Valuer 接口，用于将 StringSlice 转换为数据库中的数据
func (s StringSlice) Value() (driver.Value, error) {
	return strings.Join(s, ","), nil
}

type Product struct {
	ID            int64       `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	Name          string      `gorm:"column:name;type:varchar(50);not null;comment:商品名称" json:"name"`
	CoverImage    string      `gorm:"column:cover_image;type:varchar(200);not null;comment:商品封面图"  json:"coverImage"`
	DetailImages  StringSlice `gorm:"column:detail_images;type:varchar(1000);comment:商品详情图片" json:"detailImages"`
	Description   string      `gorm:"column:description;type:varchar(255);not null;comment:商品描述" json:"description"`
	SalePrice     float64     `gorm:"column:sale_price;type:float(10,2);not null;default:0.00;comment:销售价格" json:"salePrice"`
	CostPrice     float64     `gorm:"column:cost_price;type:float(10,2);not null;default:0.00;comment:成本价格" json:"costPrice"`
	StockQuantity int         `gorm:"column:stock_quantity;type:int;not null;default:0;comment:库存数量" json:"stockQuantity"`
	Brand         *string     `gorm:"column:brand;type:varchar(50);comment:商品品牌" json:"brand"`
	Color         StringSlice `gorm:"column:color;type:varchar(300);comment:商品颜色" json:"color"`
	Storage       StringSlice `gorm:"column:storage;type:varchar(300);comment:商品存储" json:"storage"`
	CategoryID    int         `gorm:"column:category_id;type:int;not null;default:0;comment:商品分类ID" json:"categoryId"`
	UpdateTime    time.Time   `gorm:"column:update_time;type:datetime;not null;comment:更新时间" json:"updateTime"`
	CreateTime    time.Time   `gorm:"column:create_time;type:datetime;not null;comment:创建时间" json:"createTime"`
	IsAvailable   bool        `gorm:"column:is_available;type:tinyint(1);not null;default:1;comment:上下架状态 0:下架 1:上架" json:"isAvailable"`
	IsDeleted     bool        `gorm:"column:is_deleted;type:tinyint(1);not null;default:1;comment:是否删除 0：删除 1：正常" json:"isDeleted"`
	Type          int         `gorm:"not null;default:0;comment:默认类型：0" json:"type"`
}

func (Product) TableName() string {
	return "product"
}

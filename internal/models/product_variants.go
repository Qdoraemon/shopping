package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// IntSlice 是一个自定义类型，用于处理整数切片
type IntSlice []int

// Scan 实现了 sql.Scanner 接口，用于将数据库中的数据转换为 IntSlice
func (s *IntSlice) Scan(value interface{}) error {
	if value == nil {
		*s = []int{}
		return nil
	}
	// 将数据库中的 []byte 转换为字符串
	str, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to convert %v to []byte", value)
	}
	// 将字符串转换为 []int
	return json.Unmarshal(str, s)
}

// Value 实现了 driver.Valuer 接口，用于将 IntSlice 转换为数据库中的数据
func (s IntSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type ProductVariant struct {
	ID        int      `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	ProductID int64    `gorm:"not null;comment:商品ID" json:"productId"`
	Options   IntSlice `gorm:"type:json;not null;comment:选项" json:"options"`
	Price     float64  `gorm:"type:float;not null;comment:价格" json:"price"`
	Stock     int      `gorm:"not null;comment:库存" json:"stock"`
}

func (ProductVariant) TableName() string {
	return "product_variants"
}

package models

// Option 表结构定义
type Option struct {
	ID        int64       `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	ProductID int64       `gorm:"column:product_id;type:bigint;not null;comment:商品ID" json:"product_id"`
	Name      string      `gorm:"column:name;type:varchar(50);not null;comment:选项名称" json:"name"`
	Values    StringSlice `gorm:"column:values;type:varchar(1000);comment:选项值" json:"values"`
}

func (Option) TableName() string {
	return "options"
}

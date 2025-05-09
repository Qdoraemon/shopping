package models

type Specification struct {
	ID        int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	ProductID int64  `gorm:"column:product_id;type:bigint;not null;comment:商品ID" json:"product_id"`
	Category  string `gorm:"column:category;type:varchar(255);comment:类别" json:"category"`
	Name      string `gorm:"column:name;type:varchar(255);not null;comment:名称" json:"name"`
	Value     string `gorm:"column:value;type:varchar(255);not null;comment:值" json:"value"`
}

func (Specification) TableName() string {
	return "specifications"
}

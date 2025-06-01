package models

type BaseInformation struct {
	ID            int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	Phone         string `gorm:"column:phone;type:varchar(20);not null;comment:手机号码" json:"phone"`
	Whatsapp      string `gorm:"column:whatsapp;type:varchar(20);not null;comment:WhatsApp号码" json:"whatsApp"`
	Email         string `gorm:"column:email;type:varchar(50);not null;comment:邮箱地址" json:"email"`
	Address       string `gorm:"column:address;type:varchar(255);not null;comment:地址" json:"address"`
	WechatImage   string `gorm:"column:wechat_image;type:varchar(255);comment:微信图片地址" json:"weChatImage"`
	FaceBookImage string `gorm:"column:facebook_image;type:varchar(255);comment:Facebook图片地址" json:"faceBookImage"`
}

// TableName 方法用于返回表名
func (u BaseInformation) TableName() string {
	return "base_information"
}

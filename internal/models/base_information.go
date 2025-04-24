package models

type BaseInformation struct {
	ID              int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`
	HomeTitle       string `gorm:"column:home_title;type:varchar(100);not null;comment:首页标题" json:"homeTitle"`
	HomeDescription string `gorm:"column:home_description;type:varchar(255);not null;comment:首页描述" json:"homeDescription"`
	Phone           string `gorm:"column:phone;type:varchar(20);not null;comment:手机号码" json:"phone"`
	Email           string `gorm:"column:email;type:varchar(50);not null;comment:邮箱地址" json:"email"`
	Address         string `gorm:"column:address;type:varchar(255);not null;comment:地址" json:"address"`
	WechatImage     string `gorm:"column:wechat_image;type:varchar(255);comment:微信图片地址" json:"weChatImage"`
	IsDeleted       bool   `gorm:"column:is_deleted;type:tinyint(1);not null;default:0;comment:是否刪除" json:"isDeleted"`
	IsEnable        bool   `gorm:"column:is_enable;type:tinyint(1);not null;default:0;comment:是否启用" json:"isEnable"`
}

// TableName 方法用于返回表名
func (u BaseInformation) TableName() string {
	return "base_information"
}

package models

import "time"

type User struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Email      *string   `gorm:"column:email;type:varchar(50);comment:邮箱"`
	Username   string    `gorm:"column:username;type:varchar(20);not null;comment:用户名;uniqueIndex:ones_username"`
	Password   string    `gorm:"column:password;type:varchar(64);not null;comment:密码"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;autoUpdateTime;comment:创建时间"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;autoUpdateTime;comment:更新时间"`
	Role       int       `gorm:"column:role;type:int;not null;default:1;comment:角色 1：普通用户"`
	Nickname   *string   `gorm:"column:nickname;type:varchar(50);comment:昵称"`
}

// TableName 方法用于返回表名
func (u User) TableName() string {
	return "user"
}

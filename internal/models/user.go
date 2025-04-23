package models

type User struct {
	Id          int64  `gorm:"pk autoincr 'id'"`
	UserID      int64  `gorm:"not null 'user_id'"`
	Password    string `gorm:"varchar(50) not null 'password'"`
	UserName    string `gorm:"varchar(30) 'user_name'"`
	Email       string `gorm:"varchar(50) 'email'"`
	PhoneNumber int64  `gorm:"'phone_number'"`
	Sex         string `gorm:"char(1) 'sex'"`
	Remark      string `gorm:"varchar(500) 'remark'"`
}

// TableName 方法用于返回表名
func (u User) TableName() string {
	return "user"
}

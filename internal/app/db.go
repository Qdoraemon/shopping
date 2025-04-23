package app

import (
	"fmt"
	"shopping/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var Engine *xorm.Engine
var Engine *gorm.DB

func NewMysqlDb() error {
	if Engine == nil {
		initializeMySQL()
	}
	return nil
}

// InitializeMySQL 数据库初始化
func initializeMySQL() error {

	var err error
	Engine, err = gorm.Open(mysql.Open(config.Conf.Database.Source), &gorm.Config{})

	fmt.Println("db = ", Engine)
	fmt.Println("err = ", err)

	// 连接池
	sqlDB, err := Engine.DB()
	if err != nil {
		panic("failed to connect database")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟

	// 测试数据库连接
	if err = sqlDB.Ping(); err != nil {
		panic(fmt.Sprintf("数据库连接失败: %v", err))
	}

	return nil
}

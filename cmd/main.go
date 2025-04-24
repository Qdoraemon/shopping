package main

import (
	"fmt"

	"shopping/config"
	v1 "shopping/internal/api/v1"
	"shopping/internal/app"

	"github.com/gin-gonic/gin"
)

func main() { // 主函数，程序的入口点

	// 加载配置文件
	config.LoadConfig()

	// 初始化所有模块
	app.InitializeAll()

	r := gin.Default()
	v1.RegisterApiRouter(r, app.Engine)

	err := r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		fmt.Printf("服务启动错误: %v\n", err)
		return
	}

}

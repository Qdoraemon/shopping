package main

import (
	"fmt"

	"shopping/config"
	v1 "shopping/internal/api/v1"
	"shopping/internal/app"

	"github.com/gin-gonic/gin"
)

func main() { // 主函数，程序的入口点

	// r := gin.Default()          // 创建一个默认的Gin引擎实例，包含默认的中间件（如日志和恢复）
	// router.RegisterApiRouter(r) // 调用router包中的RegisterApiRouter函数，注册API路由到Gin引擎实例

	// r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

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

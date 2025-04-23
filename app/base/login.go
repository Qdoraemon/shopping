package base

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	// 1. 获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 2. 验证参数
	if username == "" || password == "" {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "用户名或密码不能为空",
		})
		return
	}
	// 3. 返回结果
	c.JSON(200, gin.H{
		"code":    200,
		"message": "登录成功",
	})

}

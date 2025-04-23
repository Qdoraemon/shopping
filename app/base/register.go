package base

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// 验证邮箱格式
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// 验证手机号格式，这里假设是中国大陆手机号
func isValidPhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return phoneRegex.MatchString(phone)
}

// Register 处理用户注册请求
func Register(c *gin.Context) {
	// 1. 获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 2. 验证参数
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户名和密码不能为空",
		})
		return
	}

	if !isValidEmail(username) && !isValidPhone(username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请输入有效的邮箱或手机号",
		})
		return
	}

	// 这里可以添加将用户信息保存到数据库的逻辑
	// 示例代码暂时省略数据库操作

	// 3. 返回成功的 JSON 格式
	c.JSON(http.StatusCreated, gin.H{
		"message": "注册成功",
		"status":  "success",
	})
}

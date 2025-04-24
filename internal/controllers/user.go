package controllers

import (
	"net/http"
	"regexp"
	"shopping/internal/middleware"
	"shopping/internal/services"
	"shopping/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

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

func NewUserController(UserService *services.UserService) *UserController {
	return &UserController{UserService: UserService}
}

type LoginRequest struct {
	Username string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *UserController) Login(c *gin.Context) {

	var data LoginRequest
	// 绑定 JSON 数据到结构体
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//  2.需要查询数据库
	// TODO 修改规定提示码
	user, err := l.UserService.LoginUser(data.Username, data.Password)
	if err != nil {
		c.JSON(200, utils.Error(400, "用户名或密码错误"))
		return
	}
	// TODO 3. 返回结果 （默认先返回随机的）
	jwt, err := middleware.GenToken(user.Username)
	if err != nil {
		c.JSON(200, utils.Error(500, "生成token失败"))
		return
	}
	c.JSON(200, utils.Success(gin.H{
		"token": jwt,
	}, "登录成功"))

}

// TODO Register 处理用户注册请求 (未完成和未测试)
func (l *UserController) Register(c *gin.Context) {
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

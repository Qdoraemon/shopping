package main

import (
	"net/http"
	"net/http/httptest"
	"shopping/internal/middleware"
	"shopping/internal/utils"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// 初始化一个测试用的 Gin 引擎
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RateLimitMiddleware(time.Second, 10, 9))
	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
		})
		return
	})
	// r.GET("/protected", ProtectedHandler)
	return r
}

// 测试 LoginHandler
func TestLoginHandler(t *testing.T) {
	r := setupRouter()
	var s sync.WaitGroup
	number := 20
	ss, _ := utils.HashPass("123456")
	t.Log(ss)
	t.Log("sdsd")
	s.Add(number)
	for i := 0; i < number; i++ {
		if i%10 == 0 {
			time.Sleep(time.Second)
		}
		go func() {
			defer s.Done()
			// 创建一个测试请求
			req, err := http.NewRequest("POST", "/login", nil)
			if err != nil {
				t.Fatal(err)
			}

			// 创建一个响应记录器
			w := httptest.NewRecorder()

			// 调用路由处理函数
			r.ServeHTTP(w, req)

			// 断言响应状态码
			assert.Equal(t, http.StatusOK, w.Code)
			if w.Code != http.StatusOK {
				t.Log("LoginHandler 测试失败")
				return
			}
			t.Log("LoginHandler 测试通过")

			// 调用路由处理函数
		}()
	}
	s.Wait()
	t.Log("LoginHandler 测试完成")
	// assert.Equal(t, http.StatusOK, 22)
	/*
		// 创建一个测试请求
		req, err := http.NewRequest("POST", "/login", nil)
		if err != nil {
			t.Fatal(err)
		}

		// 创建一个响应记录器
		w := httptest.NewRecorder()

		// 调用路由处理函数
		r.ServeHTTP(w, req)

		// 断言响应状态码
		assert.Equal(t, http.StatusOK, w.Code)
		t.Log("LoginHandler 测试通过")
		assert.Equal(t, http.StatusBadRequest, w.Code)
	*/
}

package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 初始cap个，每fillInterval放出quantum个
func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2005,
				"msg":  "操作频繁",
			})
			fmt.Println("操作频繁")
			c.Abort()
			return
		}
		c.Next()
	}
}

package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// type RateLimiter struct {
// 	Requests map[string]int
// 	mu       sync.Mutex
// 	Limit    int
// }

var Limiter = &RateLimiter{
	requests: make(map[string]int),
	limit:    5,
}

//init is a special go function that runs automatically when app starts//

func init() {
	go func() {
		time.Sleep(time.Minute)
		limiter.mu.Lock()
		limiter.requests = make(map[string]int)
		limiter.mu.Unlock()
	}()
}

func FixedWindowRateLimiting(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter.mu.Lock()
		//schedule cleanup:release the lock before returning.//
		defer limiter.mu.Unlock()
		limiter.requests[ip]++
		if limiter.requests[ip] > limiter.limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

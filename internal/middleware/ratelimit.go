package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	Requests map[string]int
	mu       sync.Mutex
	Limit    int
}

var Limiter = &RateLimiter{
	Requests: make(map[string]int),
	Limit:    5,
}

// for a less strict instance of limiter//
var lessFreeLimit = &RateLimiter{
	Requests: make(map[string]int),
	Limit:    30,
}

//init is a special go function that runs automatically when app starts//

func init() {
	go func() {
		time.Sleep(time.Minute)
		Limiter.mu.Lock()
		Limiter.Requests = make(map[string]int)
		Limiter.mu.Unlock()
	}()
}

// for the middleware for the less strict limiter//
func LessStrictLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		lessFreeLimit.mu.Lock()
		defer lessFreeLimit.mu.Unlock()
		lessFreeLimit.Requests[ip]++
		if lessFreeLimit.Requests[ip] > lessFreeLimit.Limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			c.Abort()
			return
		}
	}
}

func FixedWindowRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		Limiter.mu.Lock()
		//schedule cleanup:release the lock before returning.//
		defer Limiter.mu.Unlock()
		Limiter.Requests[ip]++
		if Limiter.Requests[ip] > Limiter.Limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

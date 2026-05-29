package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

//the big problem with the fix window counter Algorithm is that:
//requests exceeding your limit could be exceeded before at the boundary so say:
//5 requests per IP at time 2:50s and another 5 requests for same IP after reset at say 3:20s//
//that's still a breach of 10 requests in 50s against our limit of 5 requests per minutes
//=====and that's where token bucket algorithm comes in to fix=====//

type RateLimiter struct {
	requests map[string]int
	mu       sync.Mutex
	limit    int
}

var limiter = &RateLimiter{
	requests: make(map[string]int),
	limit:    5, // 5 requests per minute
}

// this is to reset the map every single minute //
func init() {
	go func() {
		for {
			time.Sleep(time.Minute)
			limiter.mu.Lock()
			//after sleeping clear the map//
			limiter.requests = make(map[string]int)
			limiter.mu.Unlock()
		}
	}()
}

func FixedWindowRateLimit() gin.HandlerFunc {
	//the actual middleware running on every requests//
	return func(c *gin.Context) {
		ip := c.ClientIP() //get users IP
		limiter.mu.Lock()
		defer limiter.mu.Unlock() //after middleware finishes the mutex unlock automatically//
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

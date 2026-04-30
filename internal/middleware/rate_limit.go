package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var clients = make(map[string]int)
var mu sync.Mutex

func RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {

	// reset counter tiap window
	go func() {
		for {
			time.Sleep(window)
			mu.Lock()
			clients = make(map[string]int)
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {

		ip := c.ClientIP()

		mu.Lock()
		clients[ip]++
		count := clients[ip]
		mu.Unlock()

		if count > limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}

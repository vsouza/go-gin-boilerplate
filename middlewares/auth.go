package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vsouza/go-gin-boilerplate/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.GetConfig()
		reqKey := c.Request.Header.Get("X-Auth-Key")
		reqSecret := c.Request.Header.Get("X-Auth-Secret")
		key := config.GetString("http.auth.key")
		secret := config.GetString("http.auth.secret")
		if reqKey == "" || reqSecret == "" {
			c.AbortWithStatus(401)
			return
		}
		if key == "" || secret == "" {
			c.AbortWithStatus(500)
			log.Println("key and secret credentials not found on config file")
			return
		}
		if key != reqKey || secret != reqSecret {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}

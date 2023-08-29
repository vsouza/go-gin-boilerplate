package middlewares

import (
	"strings"

	"crypto/sha256"
	"crypto/subtle"

	"github.com/gin-gonic/gin"
	"github.com/vsouza/go-gin-boilerplate/config"
)

func sha256Sum(s string) []byte {
	sum := sha256.Sum256([]byte(s))
	arr := make([]byte, len(sum))
	copy(arr, sum[:])

	return arr
}

// secureCompare calculates sha256 hash of parameters a and b and does constant time comparison
// to avoid time based attacks.
func secureCompare(a, b string) int {
	aSum := sha256Sum(a)
	bSum := sha256Sum(b)

	return subtle.ConstantTimeCompare(aSum, bSum)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.GetConfig()
		reqKey := c.Request.Header.Get("X-Auth-Key")
		reqSecret := c.Request.Header.Get("X-Auth-Secret")

		var key string
		var secret string
		if key = config.GetString("http.auth.key"); len(strings.TrimSpace(key)) == 0 {
			c.AbortWithStatus(500)
		}
		if secret = config.GetString("http.auth.secret"); len(strings.TrimSpace(secret)) == 0 {
			c.AbortWithStatus(401)
		}

		isKeysEqual := secureCompare(key, reqKey) == 1
		isSecretsEqual := secureCompare(secret, reqSecret) == 1
		if !isKeysEqual || !isSecretsEqual {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}

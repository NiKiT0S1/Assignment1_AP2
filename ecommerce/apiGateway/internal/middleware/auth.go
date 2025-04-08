package middleware

import (
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[len("Basic "):])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || pair[0] != "admin" || pair[1] != "1234" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

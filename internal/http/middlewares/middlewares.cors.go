package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"book-store/internal/utils"

	"github.com/gin-gonic/gin"
)

const (
	AllowOrigin     = "http://localhost:8080"
	AllowCredential = "true"
	AllowHeader     = "Content-Type, Content-Length, Accept-Encoding, Postman-Token, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, User-Agent, Accept" // separate with ", "
	AllowMethods    = "POST, GET, PUT, DELETE, PATCH"
	MaxAge          = "43200" // for 12 hour
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", AllowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", AllowCredential)
		c.Writer.Header().Set("Access-Control-Allow-Headers", AllowHeader)
		c.Writer.Header().Set("Access-Control-Allow-Methods", AllowMethods)
		c.Writer.Header().Set("Access-Control-Max-Age", MaxAge)

		if !utils.IsArrayContains(strings.Split(AllowMethods, ", "), c.Request.Method) {
			fmt.Printf("method %s is not allowed\n", c.Request.Method)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("forbidden with CORS policy, method %s is not allowed\n", c.Request.Method)})
			return
		}

		for key, value := range c.Request.Header {
			if !utils.IsArrayContains(strings.Split(AllowHeader, ", "), key) {
				fmt.Printf("ini header %s: %s\n", key, value)
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("forbidden with CORS policy, in header %s: %s\n", key, value)})
				return
			}
		}

		if AllowOrigin != "*" {
			if !utils.IsArrayContains(strings.Split(AllowOrigin, ", "), c.Request.Host) {
				fmt.Printf("host '%s' is not part of '%v'\n", c.Request.Host, AllowOrigin)
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("forbidden with CORS policy, host '%s' is not part of allowed origin\n", c.Request.Host)})
				return
			}
		}

		c.Next()
	}
}

package JWT

import (
	"Notes-go-project/utility/middleware/JWT/tools"
	"Notes-go-project/utility/returnBody"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var UserId int

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		code := false
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = true
		} else {
			// 解析token
			newToken := strings.TrimSpace(strings.Trim(token, "Bearer"))
			claims, err := tools.ParseToken(newToken)

			if err != nil {
				code = true
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = true
			} else {
				UserId = claims.UserId
			}
		}
		if code {
			c.JSON(200, returnBody.ErrSignParam)
			c.Abort()
			return
		}
		c.Next()
	}
}

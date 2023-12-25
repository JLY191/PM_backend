package middleware

import (
	"PM_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("PM_backend")
		if err != nil {
			logrus.Info("No token")
			c.Abort()
		}
		name := utils.ParseToken(cookie)
		c.Set("username", name)
		c.Next()
	}
}

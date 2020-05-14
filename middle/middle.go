/*
@Time : 2020/5/14 11:56
@Author : mj
@File : model
@Software: GoLand
*/
package middle

import (
	"github.com/gin-gonic/gin"
	"github.com/soouCc/go-logger/logger"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Debug(c.ClientIP(),c.Request.URL.Path)
		c.Next()
	}
}

/*
@Time : 2020/5/14 11:25
@Author : mj
@File : api
@Software: GoLand
*/
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smt/conf"
	"smt/service/trace_service"
)

func Ready(c *gin.Context) {
	BackSuccess(c, nil)
}

func SetParamter(c *gin.Context) {
	conf.DataPort = c.Param("dataport")
	c.JSON(http.StatusOK, gin.H{})
	c.Abort()

	// 开始拉去数据源
	trace_service.GetData()
}

func BackSuccess(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, gin.H{})

	c.Abort()
}

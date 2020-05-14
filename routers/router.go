/*
@Time : 2020/5/14 11:55
@Author : mj
@File : router
@Software: GoLand
*/
package routers

import (
	"smt/middle"
	"smt/routers/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRounters() *gin.Engine {

	router := gin.Default()
	router.Use(middle.Middleware())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("X-AppUser", "X-AppToken", "X-AppSign")
	config.AddAllowMethods("DELETE")
	router.Use(cors.New(config))
	//获取订单价格
	router.GET("/ready", api.Ready)
	router.GET("/setParamter", api.SetParamter)
	router.GET("/trace1.data", api.Trace1Data)
	router.GET("/trace2.data", api.Trace2Data)

	return router
}

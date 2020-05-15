/*
@Time : 2020/5/14 11:55
@Author : mj
@File : router
@Software: GoLand
*/
package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"smt/middle"
	"smt/routers/api"
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

	return router
}

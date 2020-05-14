/*
@Time : 2020/5/14 11:58
@Author : mj
@File : run
@Software: GoLand
*/
package run

import (
	"fmt"
	"github.com/soouCc/go-logger/logger"
	"net/http"
	"smt/routers"
)

func Run(prot int) {
	logger.Debug("------启动服务----")
	//启动http服务器
	router := routers.InitRounters()
	_ = http.ListenAndServe(fmt.Sprintf(":%d", prot), router)
}

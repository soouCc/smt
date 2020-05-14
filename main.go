/*
@Time : 2020/5/14 15:10
@Author : mj
@File : utils
@Software: GoLand
*/
package main

import (
	"github.com/soouCc/go-logger/logger"
	"runtime"
	"smt/run"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logger.SetRollingDaily("./log", "muzhi.log", 7)
	logger.SetLevel(logger.ALL)
	run.Run(8990)
}

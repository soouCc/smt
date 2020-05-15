/*
@Time : 2020/5/14 15:50
@Author : mj
@File : trace
@Software: GoLand
*/
package trace_service

import (
	"fmt"
	"github.com/soouCc/go-logger/logger"
	"io/ioutil"
	"net/http"
	"smt/conf"
	"smt/models"
	"strings"
)

var Tmap *models.TraceMap

func init() {
	Tmap = models.NewTraceMap()
}
func GetData() {
	url1 := fmt.Sprintf("localhost:%s/trace1.data", conf.DataPort)
	url2 := fmt.Sprintf("localhost:%s/trace2.data", conf.DataPort)
	res1, err := http.Get(url1)
	if res1 == nil || res1.Body == nil {
		return
	}

	if err != nil {
		logger.Error(err.Error())
	}

	defer res1.Body.Close()

	body1, err := ioutil.ReadAll(res1.Body)
	logger.Debug(string(body1))

	//todo
	res2, err := http.Get(url2)
	if res2 == nil || res2.Body == nil {
		return
	}

	if err != nil {
		logger.Error(err.Error())
	}

	defer res2.Body.Close()

	body2, err := ioutil.ReadAll(res1.Body)
	logger.Debug(string(body2))
}

func GoGetData(url string) string {
	res1, err := http.Get(url)
	if res1 == nil || res1.Body == nil {
		return ""
	}

	if err != nil {
		logger.Error(err.Error())
	}

	defer res1.Body.Close()

	body1, err := ioutil.ReadAll(res1.Body)

	return string(body1)
}

func SetDate(data string) []string {
	data_arr := strings.Split(data, "\n")
	return data_arr
}

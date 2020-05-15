/*
@Time : 2020/5/14 17:52
@Author : mj
@File : SetDate_test
@Software: GoLand
*/
package trace_service

import (
	"encoding/json"
	"fmt"
	"github.com/soouCc/go-logger/logger"
	"io/ioutil"
	"runtime"
	"runtime/debug"
	"smt/models"
	"sync"
	"testing"
	"time"
)

func init() {

	b, err := ioutil.ReadFile("C:\\Users\\56305\\Desktop\\data\\trace1.data") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str = string(b)

	b1, err := ioutil.ReadFile("C:\\Users\\56305\\Desktop\\data\\trace2.data") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str1 = string(b1)

	da, err := ioutil.ReadFile("C:\\Users\\56305\\Desktop\\data\\checkSum.data") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	da_map = map[string]string{}
	_ = json.Unmarshal(da, &da_map)

}

var str string
var str1 string
var da_map map[string]string

func rest() {
	Tmap = models.NewTraceMap()
	Tmap2 = models.NewTraceMap()
	smp = new(models.SMap)
	smp.Data = map[string]int{}
}

func TestSetDate(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			logger.Error(string(debug.Stack()))
		}
	}()
	runtime.GOMAXPROCS(runtime.NumCPU())
	t1 := time.Now()
	rest()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		SetDate(str, Tmap)
		logger.Debug("装载数据11111:", time.Now().Sub(t1))
		wg.Done()
	}()
	go func() {
		SetDate(str1, Tmap2)
		logger.Debug("装载数据222222:", time.Now().Sub(t1))
		wg.Done()
	}()
	wg.Wait()
	//t2:=time.Now()
	//source_map := Doaction()
	_ = Doaction()
	//logger.Debug("计算时间:",time.Now().Sub(t2),len(source_map.Data))
}

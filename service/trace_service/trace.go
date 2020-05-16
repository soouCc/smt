/*
@Time : 2020/5/14 15:50
@Author : mj
@File : trace
@Software: GoLand
*/
package trace_service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/soouCc/go-logger/logger"
	"io/ioutil"
	"net/http"
	"smt/conf"
	. "smt/models"
	"strings"
	"sync"
)

var Tmap *TraceMap
var Tmap2 *TraceMap
var smp *SMap

func init() {
	Tmap = NewTraceMap()
	Tmap2 = NewTraceMap()
	smp = new(SMap)
	smp.Data = map[string]int{}
}

//拉取数据
func GetData() {
	url1 := fmt.Sprintf("localhost:%s/trace1.data", conf.DataPort)
	url2 := fmt.Sprintf("localhost:%s/trace2.data", conf.DataPort)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		data1 := GoGetData(url1)
		SetDate(data1, Tmap)
		wg.Done()
	}()
	go func() {
		data2 := GoGetData(url2)
		SetDate(data2, Tmap2)
		wg.Done()
	}()
	wg.Wait()

	Doaction()
}

func Doaction() *SSMap {
	source_map := new(SSMap)
	source_map.Data = map[string]string{}
	wg := sync.WaitGroup{}
	wg.Add(len(smp.Data))

	for k, _ := range smp.Data {
		go func() {
			_db := Tmap.Get(k)
			_db2 := Tmap2.Get(k)
			if len(_db) > 0 || len(_db2) > 0 {
				dy := sortArr(_db, _db2)
				soudata := ""
				index := 0
				for _, v := range dy {
					index++
					soudata += v.Data + "\n"
				}
				md5str := GetMD5Encode([]byte(soudata))
				source_map.Add(dy[0].TraceId, md5str)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return source_map
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

func SetDate(data string, tmp *TraceMap) {
	arr := strings.Split(data, "\n")

	//开通道 限制最大协程数量
	var td = 1000
	var ch = make(chan string, td)

	var w = sync.WaitGroup{}
	var dones = make(chan string, 1)
	for i := 0; i < td; i++ {
		go func() {
			for {
				select {
				case v := <-ch:
					//t1:=time.Now()
					var t = &Trace{}
					if t.Analysis(v) {
						tmp.Add(t)
					}
					if t.Error {
						smp.Add(t.TraceId)
					}
					w.Done()
				case <-dones:
					return
				}
			}
		}()
	}
	for _, v := range arr {
		w.Add(1)
		ch <- v
	}
	w.Wait()
	for i:=0;i<td;i++ {
		dones <- "ok"
	}
}

func GetMD5Encode(str []byte) string {
	h := md5.New()
	h.Write(str)
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

//2个有序的 合并成1个
func sortArr(a, b []*Trace) []*Trace {

	//判断数组的长度
	al := len(a)
	bl := len(b)
	cl := al + bl

	//var c [cl]int // non-constant array bound cl
	c := make([]*Trace, cl)

	ai := 0
	bi := 0
	ci := 0

	for ai < al && bi < bl {

		if a[ai].StartTime < b[bi].StartTime {
			c[ci] = a[ai]
			ci++
			ai++
		} else {
			c[ci] = b[bi]
			ci++
			bi++
		}
	}

	for ai < al {
		c[ci] = a[ai]
		ci++
		ai++
	}
	for bi < bl {
		c[ci] = b[bi]
		ci++
		bi++
	}
	return c
}

func quickSort(arr []*Trace, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2].StartTime
		for i <= j {
			for arr[i].StartTime < key {
				i++
			}
			for arr[j].StartTime > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

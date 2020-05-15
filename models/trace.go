/*
@Time : 2020/5/14 11:30
@Author : mj
@File : model
@Software: GoLand
*/
package models

import (
	"smt/utils"
	"strings"
	"sync"
)

type Trace struct {
	TraceId   string
	StartTime int64
	//SpanId       string                 `json:"spanId"`
	//ParentSpanId string                 `json:"parentSpanId"`
	//Duration     int64                  `json:"duration"`
	//ServiceName  string                 `json:"serviceName"`
	//SpanName     string                 `json:"spanName"`
	//Host         string                 `json:"host"`
	Tags map[string]interface{}
	Data string `json:"data"`
}

func (t *Trace) Analysis(str string) bool {

	data_arr := strings.Split(str, "|")

	if len(data_arr) != 9 {
		return false
	}
	t.TraceId = data_arr[0]
	t.StartTime = utils.Json2int(data_arr[1])
	//t.SpanId = data_arr[2]
	//t.ParentSpanId = data_arr[3]
	//t.Duration = utils.Json2int(data_arr[4])
	//t.ServiceName = data_arr[5]
	//t.SpanName = data_arr[6]
	//t.Host = data_arr[7]
	t.Tags = map[string]interface{}{}
	t.Data = str

	pars := strings.Split(data_arr[8], "&")
	for _, par := range pars {
		parkv := strings.Split(par, "=")
		if len(parkv) == 2 && (parkv[0] == "http.status_code" || parkv[0] == "error") {
			t.Tags[parkv[0]] = parkv[1]
		}
	}
	return true
}

func (t *Trace) Check() bool {
	if t == nil {
		return false
	}

	for k, v := range t.Tags {
		if k == "http.status_code" && v != "200" {
			return true
		}
		if k == "error" && v == "1" {
			return true
		}
	}
	return false
}

type TraceMap struct {
	sync.RWMutex
	Data map[string][]*Trace //key:traceId  value:调用链路
}

func NewTraceMap() *TraceMap {
	tm := new(TraceMap)
	tm.Data = make(map[string][]*Trace, 0)
	return tm
}

func (tm *TraceMap) Get(TraceId string) []*Trace {
	tm.RLock()
	defer tm.RUnlock()
	return tm.Data[TraceId]
}

func (tm *TraceMap) Add(t *Trace) {
	tm.Lock()
	defer tm.Unlock()
	if tm.Data[t.TraceId] == nil {
		tm.Data[t.TraceId] = []*Trace{}
	}
	tm.Data[t.TraceId] = append(tm.Data[t.TraceId], t)
}

//type TraceSlice []*Trace

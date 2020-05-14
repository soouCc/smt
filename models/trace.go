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
)

type Trace struct {
	TraceId      string                 `json:"traceId"`
	StartTime    int64                  `json:"startTime"`
	SpanId       string                 `json:"spanId"`
	ParentSpanId string                 `json:"parentSpanId"`
	Duration     int64                  `json:"duration"`
	ServiceName  string                 `json:"serviceName"`
	SpanName     string                 `json:"spanName"`
	Host         string                 `json:"host"`
	Tags         map[string]interface{} `json:"tags"`
}

func (t *Trace)Analysis(str string)  {
	data_arr := strings.Split(str,"|")

	if len(data_arr)!=9{
		return
	}
	t.TraceId = data_arr[0]
	t.StartTime = utils.Json2int(data_arr[1])

	t.SpanId = data_arr[2]
	t.ParentSpanId = data_arr[3]
	t.Duration = utils.Json2int(data_arr[4])
	t.ServiceName = data_arr[5]
	t.SpanName = data_arr[6]
	t.Host = data_arr[7]
	t.Tags = map[string]interface{}{}

	pars := strings.Split(data_arr[8], "&")
	for _, par := range pars {
		parkv := strings.Split(par, "=")
		t.Tags[parkv[0]] = parkv[1]
	}
}

func (t *Trace)Check()bool  {
	if t==nil{
		return false
	}

	for k,v:=range t.Tags{
		if k=="http.status_code"&&v!=200{
			return true
		}
		if k=="error"&&v==1{
			return true
		}
	}
	return false
}

type TraceMap struct {
	Data map[string] TraceSlice  //key:traceId  value:调用链路
}

type TraceSlice []Trace
func (s TraceSlice) Len() int { return len(s) }

func (s TraceSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }

func (s TraceSlice) Less(i, j int) bool { return s[i].StartTime < s[j].StartTime }

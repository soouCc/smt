/*
@Time : 2020/5/15 10:58
@Author : mj
@File : models
@Software: GoLand
*/
package models

import "sync"

type SMap struct {
	sync.Mutex
	Data map[string]int
}

func (sm *SMap) Add(k string) {
	sm.Lock()
	defer sm.Unlock()
	sm.Data[k]++
}

type SSMap struct {
	sync.Mutex
	Data map[string]string
}

func (sm *SSMap) Add(k, v string) {
	sm.Lock()
	defer sm.Unlock()
	sm.Data[k] = v
}

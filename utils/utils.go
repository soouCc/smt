/*
@Time : 2020/5/14 15:56
@Author : mj
@File : utils
@Software: GoLand
*/
package utils

import (
	"github.com/soouCc/go-logger/logger"
	"strconv"
)

func Json2int(hh interface{}) int64 {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	if hh == nil {
		return 0
	}

	heifan := 0
	switch hh.(type) {
	case float64:
		heifan = int(hh.(float64))
		return int64(heifan)
	case int32:
		heifan = int(hh.(int32))
		return int64(heifan)
	case int64:
		heifan = int(hh.(int64))
		return int64(heifan)
	case string:
		heifan, _ = strconv.Atoi(hh.(string))
		return int64(heifan)
	}
	return int64(hh.(int))
}


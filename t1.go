/*
@Time : 2020/5/15 11:52
@Author : mj
@File : t1
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/soouCc/go-logger/logger"
	"smt/models"
	"time"
)

func main() {

	_db := []int{}
	for i := 0; i < 1000000; i++ {
		_db = append(_db, i)
	}
	_db2 := []int{}
	for i := 5000; i < 1500000; i++ {
		_db2 = append(_db2, i)
	}

	t1 := time.Now()
	new_arr := sortArr(_db, _db2)

	logger.Debug(time.Now().Sub(t1))

	fmt.Println(new_arr)
}

func sortArr(a, b []int) []int {

	//判断数组的长度
	al := len(a)
	bl := len(b)
	cl := al + bl

	fmt.Println(cl)
	//var c [cl]int // non-constant array bound cl
	c := make([]int, cl)

	fmt.Println(len(c))
	fmt.Println(cap(c))
	ai := 0
	bi := 0
	ci := 0

	for ai < al && bi < bl {

		if a[ai] < b[bi] {
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

	/*	for i, v := range c {
			fmt.Println(i, ":", v)
		}
	*/
	return c
}
func sortArr2(a, b []*models.Trace) []*models.Trace {

	//判断数组的长度
	al := len(a)
	bl := len(b)
	cl := al + bl

	fmt.Println(cl)
	//var c [cl]int // non-constant array bound cl
	c := make([]*models.Trace, cl)

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

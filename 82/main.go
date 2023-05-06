package main

import (
	"fmt"
	"time"
)

func main() {
	/*
			test request example
			  starttime : 1614308303  2021-02-26 10:58:23
			  endtime:    1614653903  2021-03-02 10:58:23
		      timeFormat: "2006-01-02"   //go格式化日期 必须用2006年01月02日 15:04:05这个时间

		    response
		    allDateArray
	*/
	allDateArray := make([]string, 0)
	var start, end int64
	start, end = 1614308303, 1614653903
	timeFormat := "2006-01-02"
	startTime := time.Unix(start, 0)
	endTime := time.Unix(end, 0)
	//After方法 a.After(b) a,b Time类型 如果a时间在b时间之后，则返回true
	for endTime.After(startTime) {
		allDateArray = append(allDateArray, startTime.Format(timeFormat))
		startTime = startTime.AddDate(0, 0, 1)
	}
	allDateArray = append(allDateArray, endTime.Format(timeFormat))
	for _, v := range allDateArray {
		fmt.Println(v)
	}
}

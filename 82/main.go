package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(splitDate("2022-02-09", "2022-03-06", "2006-01-02"))
}

// 输入两个日期，返回当前区间内的所有日期
func splitDate(beginDate, endDate, format string) []string {
	bDate, _ := time.ParseInLocation(format, beginDate, time.Local)
	eDate, _ := time.ParseInLocation(format, endDate, time.Local)
	day := int(eDate.Sub(bDate).Hours() / 24)
	dlist := make([]string, 0)
	dlist = append(dlist, beginDate)
	for i := 1; i < day; i++ {
		result := bDate.AddDate(0, 0, i)
		dlist = append(dlist, result.Format(format))
	}
	dlist = append(dlist, endDate)
	return dlist
}

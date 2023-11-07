package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	endTimeStr := "2023-11-02 13:51:42"
	startTimeStr := "2023-11-02 00:51:00"
	// startTimeStr := time.Now().Format(layout)
	startTime, err := time.Parse(layout, startTimeStr)
	if err != nil {
		fmt.Println("解析开始时间错误:", err)
		return
	}

	endTime, err := time.Parse(layout, endTimeStr)
	if err != nil {
		fmt.Println("解析结束时间错误:", err)
		return
	}

	diff := endTime.Sub(startTime)
	// days := int(diff.Hours() / 24.0)
	days := int32(diff.Hours() / 24.0)
	fmt.Println(days)
}

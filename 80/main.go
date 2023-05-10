package main

import (
	"fmt"
	"time"
)

// 计算两个日期的相差天数
func main() {
	start, _ := time.Parse("2006-01-02", "2016-02-27")
	end, _ := time.Parse("2006-01-02", "2016-03-01")
	d := end.Sub(start)
	fmt.Println(d.Hours() / 24)
}

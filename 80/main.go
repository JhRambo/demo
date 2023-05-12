package main

import (
	"fmt"
	"time"
)

// 计算两个日期的相差天数
func main() {
	start, _ := time.Parse("2006-01-02", "2023-02-15")
	end, _ := time.Parse("2006-01-02", "2023-02-28")
	d := end.Sub(start)
	fmt.Println(d.Hours() / 24)
}

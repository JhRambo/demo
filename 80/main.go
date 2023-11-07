package main

import (
	"fmt"
	"math"
	"time"
)

// 计算两个日期的相差天数
func main() {
	layout := "2006-01-02"
	start, _ := time.Parse(layout, "2023-02-15")
	end, _ := time.Parse(layout, "2023-02-28")
	diff := end.Sub(start)
	// days :=int(diff.Hours() / 24)
	days := int(math.Ceil(diff.Hours() / 24.0)) //向上取整
	fmt.Println(days)
}

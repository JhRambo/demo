package main

import (
	"fmt"
	"time"
)

func main() {
	r := GetDurationFormatBySecond(4004)
	fmt.Println(r)
}

// 秒数转时分秒，不够两位补零0
func GetDurationFormatBySecond(sec int64) (formatString string) {
	duration := time.Duration(sec) * time.Second
	h := int64(duration.Hours())
	m := int64(duration.Minutes()) % 60
	s := int64(duration.Seconds()) % 60
	formatString = fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	return
}

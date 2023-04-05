package main

import (
	"fmt"
	"time"
)

// 时间
func main() {
	var curtime = time.Now()
	fmt.Println(curtime)                       //2022-11-23 15:27:38.505336 +0800 CST m=+0.018048301
	a := curtime.Format("2006-01-02 15:04:05") //2022-11-23 15:27:38
	fmt.Println(a)
	b := curtime.Year() //2022
	fmt.Println(b)
	c := curtime.Unix() //1669188677 时间戳
	fmt.Println(c)
	d := time.Unix(c, 0).Format("2006-01-02 15:04:05") //时间戳转时间日期格式
	fmt.Println(d)
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	e, _ := time.ParseInLocation("2006-01-02 15:04:05", d, loc)
	fmt.Println(e.Unix()) //时间日期格式转时间戳
}

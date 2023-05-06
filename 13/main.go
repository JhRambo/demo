package main

import (
	"fmt"
	"time"
)

// 时间
func main() {
	var curtime = time.Now()
	fmt.Println(curtime)                       //2023-05-05 12:57:29.3930994 +0800 CST m=+0.001567801
	a := curtime.Format("2006-01-02 15:04:05") //2023-05-05 12:57:29	年月日时分秒
	fmt.Println(a)
	b := curtime.Year() //2023	年
	fmt.Println(b)
	c := curtime.Unix() //1683262649 当前时间戳
	fmt.Println(c)
	d := time.Unix(c, 0).Format("2006-01-02 15:04:05") //时间戳转时间日期格式	2023-05-05 12:57:29
	fmt.Println(d)
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	e, _ := time.ParseInLocation("2006-01-02 15:04:05", d, loc)
	fmt.Println(e.Unix())        //时间日期格式转时间戳	1683262649
	f1, f2, f3 := curtime.Date() //返回年月日
	fmt.Println(f1)              //2023
	fmt.Println(f2)              //May
	fmt.Println(f3)              //5
}

package main

import (
	"fmt"
	"strconv"
)

// 类型转换
func main() {
	// 1、int转成string
	a := 10
	b := strconv.Itoa(a)
	fmt.Printf("%v--%T--%p--%p\n", b, b, &a, &b) //10--string--0xc00000e088--0xc000042050

	// 2、string转成int：
	c := "123456"
	d, _ := strconv.Atoi(c)
	fmt.Printf("%v--%T--%p--%p\n", d, d, &c, &d) //123456--int--0xc000042080--0xc00000e0a8

	// 3、数值类型转换
	var g = 123
	var h = 456.78
	i := float64(g) + h
	fmt.Printf("值：%.2f -- 类型：%T\n", i, i)

	// 4、其他类型转换为字符类型 Sprintf 参考demo 24
	var aa = 123    //int
	var bb = true   //bool
	var ee = 11.235 //float
	var ff = 'f'    //byte
	str1 := fmt.Sprintf("%d", aa)
	str2 := fmt.Sprintf("%t", bb)
	str3 := fmt.Sprintf("%.2f", ee)
	str4 := fmt.Sprintf("%c", ff)
	fmt.Printf("值：%v--类型：%T\n", str1, str1)
	fmt.Printf("值：%v--类型：%T\n", str2, str2)
	fmt.Printf("值：%v--类型：%T\n", str3, str3)
	fmt.Printf("值：%v--类型：%T\n", str4, str4)

	// 5、strconv转类型
	// 5.1 int转string
	// var i = 100
	// str := strconv.FormatInt(int64(i), 10)
	// fmt.Printf("值：%v--类型：%T", str, str)
	// 5.2 float转string
	// var i = 100.22
	// str := strconv.FormatFloat(float64(i), 'f', 2, 64)
	// fmt.Printf("值：%v--类型：%T", str, str)
	// 5.3 字符属于int类型
	// var s = 'g'
	// i := strconv.FormatUint(uint64(s), 10)
	// fmt.Printf("值：%v--类型：%T", i, i) //输出ascii值

	// 6、string转int
	// var s = "123456"
	// i, _ := strconv.ParseInt(s, 10, 64)
	// fmt.Printf("值：%v--类型：%T", i, i)

	// 7、string转float
	var s = "123456.88"
	n, _ := strconv.ParseFloat(s, 64)
	fmt.Printf("值：%v--类型：%T", n, n)
}

package main

import "fmt"

// [a:b] 左闭右开
func main() {
	//非中文字符截取
	str := "hello golang"
	substr := []byte(str)[2:7]
	fmt.Println(substr)         //[108 108 111 32 103] ASCII码值
	fmt.Println(string(substr)) //llo g

	//中文字符截取
	str2 := "你好golang"
	substr2 := []rune(str2)[1:5]
	fmt.Println(substr2)         //[22909 103 111 108]	utf8或unicode编码
	fmt.Println(string(substr2)) //好gol

	str3 := "你好golang"
	for i := 0; i < len(str3); i++ { //byte类型 ASCII码值
		fmt.Printf("%v(%c)", str3[i], str3[i]) //228(ä)189(½)160( )229(å)165(¥)189(½)103(g)111(o)108(l)97(a)110(n)103(g)20320
	}
	for _, v := range str3 { //rune类型 unicode或utf8类型
		fmt.Printf("%v(%c)", v, v) //20320(你)22909(好)103(g)111(o)108(l)97(a)110(n)103(g)
	}

	//非中文字符字符替换
	str4 := "hello golang"
	str4_ := []byte(str4)
	str4_[2] = 'g'
	fmt.Println(string(str4_)) //heglo golang

	//中文字符字符替换
	str5 := "你好 golang"
	str5_ := []rune(str5)
	str5_[1] = '嗨'
	fmt.Println(string(str5_)) //你嗨 golang
}

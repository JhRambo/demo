package main

import (
	"fmt"
)

func main() {
	var arr1 []int
	arr1 = append(arr1, 2)
	fmt.Printf("值：%#v，类型：%T\n", arr1, arr1) //值：[]int{2}，类型：[]int

	var arr2 = [2]int{1}
	fmt.Printf("值：%#v，类型：%T\n", arr2, arr2) //值：[2]int{1, 0}，类型：[2]int

	var arr3 = [...]int{1, 2, 3, 4}
	arr3[3] = 55
	fmt.Printf("值：%#v，类型：%T，长度：%v，容量：%v\n", arr3, arr3, len(arr3), cap(arr3)) //值：[4]int{1, 2, 3, 55}，类型：[4]int，长度：4，容量：4

	var arr4 = [...]int{3: 44}
	arr4[3] = 55
	fmt.Printf("值：%#v，类型：%T，长度：%v，容量：%v\n", arr4, arr4, len(arr4), cap(arr4)) //值：[4]int{0, 0, 0, 55}，类型：[4]int，长度：4，容量：4

	// 1.数组是值类型
	var arr5 = [...]string{"name", "zhsna"}
	fmt.Printf("值：%#v，类型：%T\n", arr5, arr5) //值：[2]string{"name", "zhsna"}，类型：[2]string
	var arr6 = [...]string{3: "name", "zhsna"}
	fmt.Printf("值：%#v，类型：%T，长度：%v，容量：%v\n", arr6, arr6, len(arr6), cap(arr6)) //值：[5]string{"", "", "", "name", "zhsna"}，类型：[5]string，长度：5，容量：5
	// 2.切片是引用类型
	var arr7 = []string{"name", "zhsna"}
	fmt.Printf("值：%#v，类型：%T\n", arr7, arr7) //值：[]string{"name", "zhsna"}，类型：[]string
	var arr8 = []string{3: "name", "zhsna"}
	fmt.Printf("值：%#v，类型：%T，长度：%v，容量：%v\n", arr8, arr8, len(arr8), cap(arr8)) //值：[]string{"", "", "", "name", "zhsna"}，类型：[]string，长度：5，容量：5
	// 3.map是引用类型
	var arr9 = map[string]string{
		"name": "张三",
		"age":  "20",
	}
	fmt.Printf("值：%#v，类型：%T，长度：%v\n", arr9, arr9, len(arr9)) //值：map[string]string{"age":"20", "name":"张三"}，类型：map[string]string，长度：2
	arr10 := arr9
	arr10["name"] = "李四"
	fmt.Println(arr9, arr10) //map[age:20 name:李四] map[age:20 name:李四]
}

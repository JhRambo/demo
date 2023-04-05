package main

import (
	"fmt"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}

// 该 方法 属于 Circle 结构体中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func main() {
	//第一种方式实例化结构体
	// var cc Circle
	// cc.radius = 10.00
	// 第二种方式实例化结构体
	cc := Circle{
		radius: 10.00,
	}
	fmt.Println("圆的面积 = ", cc.getArea())
}

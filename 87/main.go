package main

type S struct{}

// 定义A接口包含两个方法
type A interface {
	f1() string
	f2() string
	// f3() string	结构体S没有这个方法，报错
}

func (S) f1() string {
	return "f1"
}

func (S) f2() string {
	return "f2"
}

// 这里返回接口A，但实际上返回了S这个结构体却不会报错，因为f1 f2 两个方法属于 S 这个结构体的
func ff() A {
	return &S{}
}

func main() {

}

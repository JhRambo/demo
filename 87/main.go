package main

type S struct{}

// 定义A接口包含两个方法
// f1 f2 两个方法属于 S 这个结构体的
type A interface {
	f1() string
	f2() string
}

func (S) f1() string {
	return "f1"
}

func (S) f2() string {
	return "f2"
}

// 这里返回接口A，但实际上返回了S这个结构体
func ff() A {
	return &S{}
}

func main() {

}

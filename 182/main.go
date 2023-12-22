package main

type MyInterface interface {
	Method1() string
	Method2() int
}
type B struct {
	C MyInterface
}

type MyStruct struct{}

func (m *MyStruct) Method1() string {
	return "abc"
}

func (m *MyStruct) Method2() int {
	return 123
}

func (m *MyStruct) Method3() string {
	return "abc"
}

func main() {
	var d MyInterface
	d = &MyStruct{}
	d.Method1()
	d.Method2()
	e := &MyStruct{}
	e.Method3()
	f := &B{
		C: &MyStruct{},
	}
	f.C.Method1()
	f.C.Method2()
}

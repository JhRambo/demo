package main

import (
	"fmt"
)

// Interface 定义了一个操作方法
type Interface interface {
	DoSomething()
}

// Implementation 是 Interface 的实现
type Implementation struct {
	message string
}

// DoSomething 方法实现了 Interface 中的操作方法
func (i *Implementation) DoSomething() {
	fmt.Printf("%s\n", i.message)
}

// NewImplementation 是 Implementation 的构造函数
func NewImplementation(message string) Interface {
	return &Implementation{
		message: message,
	}
}

// Service 是依赖于 Interface 的服务
type Service struct {
	implementation Interface
}

// NewService 是 Service 的构造函数，注入实现了 Interface 的对象
func NewService(implementation Interface) *Service {
	return &Service{
		implementation: implementation,
	}
}

// DoSomething 是 Service 提供给外部使用的方法，它实际上是调用了依赖的 Implementation 的方法
func (s *Service) DoSomething() {
	s.implementation.DoSomething()
}

func main() {
	implementation := NewImplementation("Hello, World!")
	service := NewService(implementation)
	service.DoSomething()
}

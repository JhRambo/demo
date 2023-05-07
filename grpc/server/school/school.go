package main

import (
	"demo/grpc/proto/school"
	"fmt"
)

func main3() {
	//创建学生信息
	var students []*school.Student
	for i := 0; i < 3; i++ {
		var sex = (school.Sex)(i % 3)
		student := &school.Student{
			Name: fmt.Sprintf("Student_%d", i),
			Age:  10,
			Sex:  sex,
		}
		students = append(students, student)
	}
	//创建班级信息
	var myClass school.Class
	myClass.Name = "我的班级"
	myClass.Students = students
	fmt.Println("data==============", &myClass)
}

package main

import (
	pb "demo/grpc/proto/school"
	"fmt"
	"log"
)

func main() {
	//创建学生信息
	var students []*pb.Student
	count := 5
	for i := 1; i <= count; i++ {
		sex := (pb.Sex)(0)
		student := &pb.Student{
			Name: fmt.Sprintf("Student_%d", i),
			Age:  int32(i),
			Sex:  sex,
		}
		log.Println("==============", student)
		students = append(students, student)
	}
	//创建班级信息
	var myClass pb.Class
	myClass.Name = "我的班级"
	myClass.Students = students
}

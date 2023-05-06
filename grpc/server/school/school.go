package main

import (
	"demo/grpc/proto/school"
	"fmt"

	"google.golang.org/protobuf/proto"
)

// func writeProto(filename string) (err error) {
// 	//创建学生信息
// 	var students []*school.Student
// 	for i := 0; i < 10; i++ {
// 		var sex = (school.Sex)(i % 10)
// 		student := &school.Student{
// 			Name: fmt.Sprintf("Student_%d", i),
// 			Age:  10,
// 			Sex:  sex,
// 		}
// 		students = append(students, student)
// 	}

// 	//创建班级信息
// 	var myClass school.Class
// 	myClass.Name = "我的班级"
// 	myClass.Students = students

// 	data, err := proto.Marshal(&myClass)
// 	if err != nil {
// 		fmt.Printf("marshal proto buf failed, err:%v\n", err)
// 		return
// 	}

// 	err = ioutil.WriteFile(filename, data, 0666)
// 	if err != nil {
// 		fmt.Printf("write file failed, err:%v\n", err)
// 		return
// 	}
// 	return
// }

func writeProto() {
	//创建学生信息
	var students []*school.Student
	for i := 0; i < 5; i++ {
		var sex = (school.Sex)(i % 10)
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

	data, err := proto.Marshal(&myClass)
	if err != nil {
		fmt.Printf("marshal proto buf failed, err:%v\n", err)
		return
	}
	fmt.Println("data1==============", data)
	proto.Unmarshal(data, &myClass)
	fmt.Println("data2==============", &myClass)
	// err = ioutil.WriteFile(filename, data, 0666)
	// if err != nil {
	// 	fmt.Printf("write file failed, err:%v\n", err)
	// 	return
	// }
	// return
}

func main() {
	writeProto()
}

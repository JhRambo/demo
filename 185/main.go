package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"demo/utils/proto/person"

	"github.com/golang/protobuf/proto"
)

func main() {
	person := &person.Person{
		Name: "Alice",
		Address: &person.Address{
			City:   "Beijing",
			Detail: "1001街道",
		},
		NodeList: []string{"{\"id\":\"111\"}"},
	}

	fmt.Println("原始结构：", person)

	//1.根据path更新value
	err := UpdateFieldValue(person, "Name", "Bob")
	if err != nil {
		fmt.Println("更新 Name 字段失败：", err)
	}

	err = UpdateFieldValue(person, "Address.City", "Shanghai")
	if err != nil {
		fmt.Println("更新 Address.City 字段失败：", err)
	}

	err = UpdateFieldValue(person, "Job.Work", "农民")
	if err != nil {
		fmt.Println("更新 Job.Work 字段失败：", err)
	}

	nodeList := []string{"{\"id\":\"222\"}", "{\"id\":\"333\"}"}
	err = UpdateFieldValue(person, "NodeList", nodeList)
	if err != nil {
		fmt.Println("更新 NodeList 字段失败：", err)
	}

	fmt.Println("更新后的结构：", person)

	//2.根据path获取value
	value, err := GetFieldValue(person, "Address")
	if err != nil {
		fmt.Println("获取字段值失败：", err)
		return
	}
	fmt.Println("Address字段值：", value)

	//3.根据path删除value
	err = DeleteFieldValue(person, "Address.City")
	if err != nil {
		fmt.Println("删除字段失败：", err)
		return
	}
	fmt.Println("删除后的结构：", person)
	fmt.Println(person.Address)
}

// UpdateFieldValue 更新 Protocol Buffers 消息的字段值
func UpdateFieldValue(msg proto.Message, fieldPath string, newValue interface{}) error {
	v := reflect.ValueOf(msg).Elem()

	// 拆分字段路径
	fieldNames := strings.Split(fieldPath, ".")

	// 遍历字段路径
	for i, fieldName := range fieldNames {
		// 获取字段值的 reflect.Value
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				// 如果是 nil 指针，则创建新的实例
				v.Set(reflect.New(v.Type().Elem()))
			}
			v = v.Elem()
		}

		// 检查字段类型
		if v.Kind() != reflect.Struct {
			return fmt.Errorf("字段路径错误：%s 不是结构体类型", strings.Join(fieldNames[:i+1], "."))
		}

		// 获取字段的 reflect.Value
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			// 如果字段不存在，创建新字段
			newFieldValue := reflect.New(field.Type()).Elem()
			v.FieldByName(fieldName).Set(newFieldValue)
			v = newFieldValue
		} else {
			v = field
		}
	}

	// 设置新的字段值
	switch v.Interface().(type) {
	case []string:
		newValueSlice, ok := newValue.([]string)
		if !ok {
			return errors.New("新值类型不匹配")
		}
		v.Set(reflect.AppendSlice(v, reflect.ValueOf(newValueSlice)))
	default:
		newValueReflect := reflect.ValueOf(newValue)
		if newValueReflect.Type() != v.Type() {
			return fmt.Errorf("字段类型不匹配")
		}
		v.Set(newValueReflect)
	}

	return nil
}

// GetFieldValue 根据字段路径获取 Protocol Buffers 消息的字段值
func GetFieldValue(msg proto.Message, fieldPath string) (interface{}, error) {
	v := reflect.ValueOf(msg).Elem()

	// 拆分字段路径
	fieldNames := strings.Split(fieldPath, ".")

	// 遍历字段路径
	for _, fieldName := range fieldNames {
		// 获取字段值的 reflect.Value
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return nil, fmt.Errorf("字段路径错误：%s 为空指针", fieldPath)
			}
			v = v.Elem()
		}

		// 检查字段类型
		if v.Kind() != reflect.Struct {
			return nil, fmt.Errorf("字段路径错误：%s 不是结构体类型", fieldPath)
		}

		// 获取字段的 reflect.Value
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			return nil, fmt.Errorf("字段 %s 不存在", fieldName)
		}
		v = field
	}

	// 返回字段值
	return v.Interface(), nil
}

// DeleteFieldValue 根据字段路径删除 Protocol Buffers 消息的字段
func DeleteFieldValue(msg proto.Message, fieldPath string) error {
	v := reflect.ValueOf(msg).Elem()

	// 拆分字段路径
	fieldNames := strings.Split(fieldPath, ".")

	// 遍历字段路径
	for i, fieldName := range fieldNames {
		// 获取字段值的 reflect.Value
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return fmt.Errorf("字段路径错误：%s 为空指针", fieldPath)
			}
			v = v.Elem()
		}

		// 检查字段类型
		if v.Kind() != reflect.Struct {
			return fmt.Errorf("字段路径错误：%s 不是结构体类型", fieldPath)
		}

		// 获取字段的 reflect.Value
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			return fmt.Errorf("字段 %s 不存在", fieldName)
		}

		// 如果是最后一个字段，则删除该字段
		if i == len(fieldNames)-1 {
			field.Set(reflect.Zero(field.Type()))
			return nil
		}

		v = field
	}

	return nil
}

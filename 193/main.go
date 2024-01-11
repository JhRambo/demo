package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	pb "demo/utils/proto/person"
)

func main() {
	as := []*pb.X{}
	as = append(as, &pb.X{
		A1: 1,
		A2: "X-detail",
	}, &pb.X{
		A1: 2,
		A2: "X-detail",
	})
	info1 := &pb.AreaInfo{
		A: as,
		B: "one",
	}
	info2 := &pb.AreaInfo{
		A: as,
		B: "two",
	}

	areaInfos := []*pb.AreaInfo{}
	areaInfos = append(areaInfos, info1, info2)

	p := &pb.Person{
		Status: pb.Status_P_NORMAL,
		Name:   "Alice",
		Age:    25,
		Address: &pb.AddressInfo{
			Email: "123@example.com",
			Code:  "666",
			Area:  areaInfos,
		},
	}

	// bys, _ := json.Marshal(p)
	// fmt.Println(string(bys))

	person := &pb.Person{}
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	err = proto.Unmarshal(data, person)
	if err != nil {
		log.Fatal(err)
	}

	pr := person.ProtoReflect() //protoreflect.Message

	// fmt.Println(">>>>>>>>>>>>>>>>>> Before update >>>>>>>>>>>>>>>>>>")
	// GetProtoMessage("person", pr)

	// // // 1.修改 path=name和path=age  首先获取 name和age 字段的描述符 protoreflect.FieldDescriptor，然后再修改 name和age 字段的值
	// // nameField := pr.Descriptor().Fields().ByName(protoreflect.Name("name"))
	// // pr.Set(nameField, protoreflect.ValueOf("Bob"))
	// // ageField := pr.Descriptor().Fields().ByName("age")
	// // pr.Set(ageField, protoreflect.ValueOf(int32(30)))
	// fmt.Println(">>>>>>>>>>>>>>>>>> After update1 >>>>>>>>>>>>>>>>>>")
	// UpdateProtoMessage(pr, "name", "jack")
	// UpdateProtoMessage(pr, "age", int32(30))
	// GetProtoMessage("person", pr)
	// return

	// // // 2.修改 path=address.email  首先获取 address 消息的 protoreflect.Message，再获取 email 字段的描述符 protoreflect.FieldDescriptor，最后修改 email 字段的值
	// // addressField := pr.Descriptor().Fields().ByName("address")
	// // addressValue := pr.Get(addressField)
	// // if addressValue.IsValid() && addressValue.Message().IsValid() {
	// // 	emailField := addressValue.Message().Descriptor().Fields().ByName("email")
	// // 	if emailField != nil {
	// // 		addressValue.Message().Set(emailField, protoreflect.ValueOfString("666@example.com"))
	// // 	}
	// // }
	// fmt.Println(">>>>>>>>>>>>>>>>>> After update2 >>>>>>>>>>>>>>>>>>")
	// UpdateProtoMessage(pr, "address.email", "666@example.com")
	// GetProtoMessage("person", pr)
	// return

	// // 3.修改 path=address.area[0].b  首先获取 address.area 消息的 protoreflect.List，再获取 [0] 消息的 protoreflect.Message，再获取 b 字段的描述符 protoreflect.FieldDescriptor，最后修改 b 字段的值
	// // path = "address.area[0].b"
	// infoListField := addressValue.Message().Descriptor().Fields().ByName("area")
	// infoListValue := addressValue.Message().Get(infoListField)
	// if infoListValue.List().Len() > 0 { //数组
	// 	infoMessage := infoListValue.List().Get(0).Message() //用索引查找
	// 	aField := infoMessage.Descriptor().Fields().ByName("b")
	// 	infoMessage.Set(aField, protoreflect.ValueOfString("oneoneone"))
	// }
	// fmt.Println(">>>>>>>>>>>>>>>>>> After update3 >>>>>>>>>>>>>>>>>>")
	// UpdateProtoMessage(pr, "address.area[0].b", "onetwothreego")
	// GetProtoMessage("person", pr)
	// return

	// 4.修改 path=address.area[1].a.a1
	fmt.Println(">>>>>>>>>>>>>>>>>> After update4 >>>>>>>>>>>>>>>>>>")
	UpdateProtoMessage(pr, "address.area[1].a[1].a1", int32(999))
	GetProtoMessage("person", pr)

	// // 5.删除 protoreflect.Message 的字段
	// pr.Clear(pr.Descriptor().Fields().ByName("age"))
	// fmt.Println(">>>>>>>>>>>>>>>>>> After update5 >>>>>>>>>>>>>>>>>>")
	// GetProtoMessage("person", pr)

	// // 6.删除 path=address.info[1]  TODO
	// if infoListValue.List().Len() > 1 {
	// }
	// fmt.Println(">>>>>>>>>>>>>>>>>> After update6 >>>>>>>>>>>>>>>>>>")
	// GetProtoMessage("person", pr)

	// // 6.查询 path=name 的值
	// nameValue := pr.Get(nameField)
	// fmt.Println("person.name:", nameValue.Interface())

	// // 7.查询 path=address.info[1] 的值

	// // 8.查询 path=不存在 的值
	// xField := pr.Descriptor().Fields().ByName("x")
	// if xField == nil {
	// 	fmt.Printf("%s field not found", "x")
	// 	return
	// }
	// xValue := pr.Get(xField)
	// fmt.Println("person.x:", xValue.Interface())
}

// 获取message
func GetProtoMessage(prefix string, message protoreflect.Message) map[string]interface{} {
	result := make(map[string]interface{})

	message.Range(func(fd protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		fieldName := fmt.Sprintf("%s.%s", prefix, fd.Name())
		if fd.IsList() { // 数组
			list := value.List()
			for i := 0; i < list.Len(); i++ {
				element := list.Get(i)
				elementFieldName := fmt.Sprintf("%s[%d]", fieldName, i)
				elementResult := GetProtoMessage(elementFieldName, element.Message())
				result[elementFieldName] = elementResult
			}
		} else {
			switch fd.Kind() { // 字段类型
			case protoreflect.MessageKind: // 消息类型
				if value.Message().IsValid() {
					messageResult := GetProtoMessage(fieldName, value.Message())
					result[fieldName] = messageResult
				}
			case protoreflect.EnumKind: // 枚举类型 TODO
				enumValue := int32(value.Enum())
				enumNumber := protoreflect.EnumNumber(enumValue)
				enumDescriptor := fd.Enum().Values().ByNumber(enumNumber)
				enumName := enumDescriptor.Name()
				result[fieldName] = map[string]interface{}{
					"value": enumValue,
					"name":  enumName,
				}
				fmt.Printf("%s: %d (%s)\n", fieldName, enumValue, enumName)
			default:
				result[fieldName] = value.Interface()
				fmt.Printf("%s: %v\n", fieldName, value.Interface())
			}
		}
		return true
	})

	return result
}

// 更新message
func UpdateProtoMessage(message protoreflect.Message, path string, data interface{}) error {
	paths := ParsePath1(path)
	targetFieldName := paths[len(paths)-1] //最后一个元素即目标字段

	// 遍历每个路径
	for i := 0; i < len(paths); i++ {
		fieldName := paths[i]

		var field protoreflect.FieldDescriptor
		var fieldNames []string
		if strings.Contains(fieldName, "[") {
			fieldNames = ParsePath2(fieldName)
			// 获取当前字段描述符
			field = message.Descriptor().Fields().ByName(protoreflect.Name(fieldNames[0]))
		} else {
			// 获取当前字段描述符
			field = message.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		}
		if field == nil {
			return fmt.Errorf("field %s not found", fieldName)
		}

		if field.IsList() { // 数组
			index, err := strconv.Atoi(fieldNames[1])
			if err != nil {
				return fmt.Errorf("Failed to convert string to int:%v", err)
			}
			list := message.Get(field).List()
			if index >= list.Len() {
				return fmt.Errorf("array index out of bounds: %d", index)
			}
			if i == len(paths)-1 {
				// 最后一个字段为数组元素，直接设置
				listMessage := list.Get(index).Message()
				aField := listMessage.Descriptor().Fields().ByName(protoreflect.Name(targetFieldName))
				listMessage.Set(aField, protoreflect.ValueOf(data))
			} else {
				// 最后一个字段为数组元素的子字段，递归处理
				subMessage := list.Get(index).Message()
				if !subMessage.IsValid() {
					return fmt.Errorf("invalid sub-message at %s[%d]", fieldName, index)
				}
				if err := UpdateProtoMessage(subMessage, strings.Join(paths[i+1:], "."), data); err != nil {
					return err
				}
			}
		} else {
			switch field.Kind() { // 字段类型
			case protoreflect.MessageKind:
				// 处理嵌套的消息类型
				if i == len(paths)-1 {
					// 最后一个字段为子消息的字段，递归处理
					subMessage := message.Get(field).Message()
					if !subMessage.IsValid() {
						return fmt.Errorf("invalid sub-message at %s", fieldName)
					}
					if err := UpdateProtoMessage(subMessage, "", data); err != nil {
						return err
					}
				} else {
					// 非最后一个字段为子消息字段，继续遍历
					message = message.Get(field).Message()
					if !message.IsValid() {
						return fmt.Errorf("invalid message at %s", fieldName)
					}
				}
			default:
				// 处理简单的字段类型
				if i == len(paths)-1 {
					// 最后一个字段为简单类型字段，直接设置
					aField := message.Descriptor().Fields().ByName(protoreflect.Name(targetFieldName))
					message.Set(aField, protoreflect.ValueOf(data))
				} else {
					// 非最后一个字段为简单类型的子字段，无法继续遍历
					return fmt.Errorf("%s is a simple type field, cannot set sub-field", fieldName)
				}
			}
		}
	}

	return nil
}

// 返回[x[1]]
func ParsePath1(path string) []string {
	return strings.Split(path, ".")
}

// 返回[x 1]
func ParsePath2(path string) []string {
	var result []string
	// 利用正则表达式匹配路径中的各级字段名和数组索引
	re := regexp.MustCompile(`([^\[\].]+)|(\d+)`)
	matches := re.FindAllStringSubmatch(path, -1)
	// 遍历匹配结果，组装各级字段名和数组索引
	for _, match := range matches {
		if match[1] != "" {
			result = append(result, match[1])
		} else {
			result = append(result, match[2])
		}
	}

	return result
}

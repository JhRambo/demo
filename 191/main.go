package main

import (
	"fmt"
	"log"
	"strings"

	"demo/utils/proto/p"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// 内存中保存的数据结构 全局变量
var MemoryData protoreflect.Message

// 更新内存中的值
func UpdateMemoryData(data protoreflect.Message, path string, fieldMask *field_mask.FieldMask, value interface{}) error {
	// 根据path解析出各级字段名
	fieldNames, err := ParsePath(path)
	if err != nil {
		return err
	}

	// 获取要更新的目标字段名
	targetFieldName := fieldNames[len(fieldNames)-1]

	// 根据路径逐级遍历查找需要更新的字段
	currentData := data
	for _, fieldName := range fieldNames[:len(fieldNames)-1] {
		// 获取当前字段对应的消息类型
		fieldDescriptor := currentData.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		if fieldDescriptor == nil {
			// 如果字段不存在，则创建一个空的子数据
			subData := currentData.NewField(fieldDescriptor)
			currentData.Set(fieldDescriptor, subData)
			currentData = subData.Message()
		} else {
			switch fieldDescriptor.Kind() {
			case protoreflect.MessageKind:
				// 获取当前字段的值
				fieldValue := currentData.Get(fieldDescriptor).Message()

				if fieldValue.IsValid() {
					currentData = fieldValue
				} else {
					// 如果字段不存在，则创建一个空的子数据
					subData := currentData.NewField(fieldDescriptor)
					currentData.Set(fieldDescriptor, subData)
					currentData = subData.Message()
				}
			default:
				return fmt.Errorf("field %s is not a valid sub-data", fieldName)
			}
		}
	}

	// 更新目标字段的值
	fieldDescriptor := currentData.Descriptor().Fields().ByName(protoreflect.Name(targetFieldName))
	currentData.Set(fieldDescriptor, protoreflect.ValueOf(value))

	// 根据fieldMask进一步过滤需要更新的字段
	filteredData := FilterData(data, fieldMask)
	MergeData(data, filteredData)

	return nil
}

// 合并源数据和过滤后的数据
func MergeData(source protoreflect.Message, filtered protoreflect.Message) {
	sourceDescriptor := source.Descriptor()
	filteredDescriptor := filtered.Descriptor()

	// 遍历源数据的字段
	for i := 0; i < sourceDescriptor.Fields().Len(); i++ {
		sourceFieldDescriptor := sourceDescriptor.Fields().Get(i)

		// 获取对应字段名的过滤后的字段描述符
		filteredFieldDescriptor := filteredDescriptor.Fields().ByName(sourceFieldDescriptor.Name())
		if filteredFieldDescriptor == nil {
			continue // 继续遍历下一个字段
		}

		// 获取源数据和过滤后数据字段的值
		sourceFieldValue := source.Get(sourceFieldDescriptor)
		filteredFieldValue := filtered.Get(filteredFieldDescriptor)

		// 如果过滤后的字段值有效，则将源数据中的字段值替换为过滤后的值
		if filteredFieldValue.IsValid() {
			sourceFieldValue = filteredFieldValue
		}

		// 将过滤后的值设置到源数据中
		source.Set(sourceFieldDescriptor, sourceFieldValue)
	}
}

// 根据path获取内存中的值
func GetMemoryData(data protoreflect.Message, path string) ([]byte, error) {
	// 根据path解析出各级字段名
	fieldNames, err := ParsePath(path)
	if err != nil {
		return nil, err
	}

	// 根据路径逐级遍历查找目标字段的值
	currentData := data
	for _, fieldName := range fieldNames {
		fieldDescriptor := currentData.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		if fieldDescriptor == nil {
			return nil, fmt.Errorf("field %s not found", fieldName)
		}
		if fieldDescriptor.Kind() == protoreflect.MessageKind {
			currentData = currentData.Get(fieldDescriptor).Message()
			if currentData == nil {
				return nil, fmt.Errorf("field %s not found", fieldName)
			}
		} else {
			// 将消息序列化为字节流返回
			result, err := proto.Marshal(currentData.Interface())
			if err != nil {
				return nil, fmt.Errorf("failed to marshal message: %v", err)
			}
			return result, nil
		}
	}

	// 将消息序列化为字节流返回
	result, err := proto.Marshal(currentData.Interface())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %v", err)
	}
	return result, nil
}

// 根据fieldMask过滤数据
func FilterData(data protoreflect.Message, fieldMask *field_mask.FieldMask) protoreflect.Message {
	filteredData := data.New()

	for _, path := range fieldMask.Paths {
		fieldNames, err := ParsePath(path)
		if err != nil {
			continue
		}

		currentData := data
		currentFilteredData := filteredData
		for _, fieldName := range fieldNames {
			fieldDescriptor := currentData.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
			if fieldDescriptor == nil {
				break
			}
			switch fieldDescriptor.Kind() {
			case protoreflect.MessageKind:
				fieldValue := currentData.Get(fieldDescriptor).Message()
				if !fieldValue.IsValid() {
					break
				}
				subFilteredData := currentFilteredData.Mutable(fieldDescriptor).Message().New()
				currentFilteredData.Set(fieldDescriptor, protoreflect.ValueOf(subFilteredData))
				currentData = fieldValue
				currentFilteredData = subFilteredData
			default:
				currentFilteredData.Set(fieldDescriptor, currentData.Get(fieldDescriptor))
				return filteredData
			}
		}
	}

	return filteredData
}

// 解析path，获取各级字段名
func ParsePath(path string) ([]string, error) {
	// 假设path的格式为：Scene.NodeList[0].FileInfo[1].name
	// 先按点号分割字段名
	fieldNames := strings.Split(path, ".")

	// 逐个处理字段名
	var result []string
	for _, fieldName := range fieldNames {
		// 判断是否包含索引
		if idx := strings.Index(fieldName, "["); idx != -1 {
			fieldName = fieldName[:idx] // 截取方括号之前的部分
		}
		result = append(result, fieldName)
	}

	return result, nil
}

func main() {
	// 创建一个示例的WsData对象
	wsData := &p.WsData{
		MethodId: p.MethodId_DATA_UPDATE,
		Path:     "Scene.NodeList[0].FileInfo[1].name",
		FieldMask: &field_mask.FieldMask{
			Paths: []string{},
		},
		Data: &p.WsData_StringData{StringData: "new value"},
	}

	// 将WsData对象序列化为字节流
	dataBytes, err := proto.Marshal(wsData)
	if err != nil {
		log.Fatal(err)
	}

	// 反序列化字节流为WsData对象
	receivedWsData := &p.WsData{}
	err = proto.Unmarshal(dataBytes, receivedWsData)
	if err != nil {
		log.Fatal(err)
	}

	// // 根据路径获取对应的消息类型
	// messageType, err := getMessageTypeByName(receivedWsData.Path)
	// if err != nil {
	// 	log.Fatalf("failed to find message type: %v", err)
	// }
	// memoryData := messageType.New().Interface().(protoreflect.Message)

	// 根据接收到的WsData更新内存中的数据
	path := receivedWsData.Path
	fieldMask := receivedWsData.FieldMask
	value := receivedWsData.Data
	err = UpdateMemoryData(wsData.ProtoReflect(), path, fieldMask, value)
	if err != nil {
		log.Fatal(err)
	}

	// 获取更新后的值
	result, err := GetMemoryData(wsData.ProtoReflect(), path)
	if err != nil {
		log.Fatal(err)
	}

	// 打印结果
	fmt.Println(string(result))
}

// 根据路径获取对应的消息类型
func getMessageTypeByName(path string) (protoreflect.MessageType, error) {
	descriptor, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(path))
	if err != nil {
		return nil, fmt.Errorf("message type %s not found", path)
	}

	messageType := descriptor.(protoreflect.MessageType)
	return messageType, nil
}

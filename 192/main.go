package main

import (
	"fmt"
	"log"
	"strings"

	"demo/utils/proto/p"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DataStore struct {
	data protoreflect.Message
}

func NewDataStore(data protoreflect.Message) *DataStore {
	return &DataStore{
		data: data,
	}
}

// 内存中保存的数据结构 全局变量
var MemoryData map[string]*DataStore

// 获取message结构
func GetProtoMessage(prefix string, message protoreflect.Message) map[string]interface{} {
	result := make(map[string]interface{})

	message.Range(func(fd protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		fieldName := fmt.Sprintf("%s.%s", prefix, fd.Name())
		if fd.IsList() { // array
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
			case protoreflect.EnumKind: // 枚举类型
				enumValue := int32(value.Enum())
				result[fieldName] = enumValue
			default:
				result[fieldName] = value.Interface()
				fmt.Printf("%s: %v\n", fieldName, value.Interface())
			}
		}
		return true
	})

	return result
}

// 更新内存中的值
func (ds *DataStore) UpdateMemoryData(data protoreflect.Message, fieldNames []string, targetFieldName string, fieldMask *field_mask.FieldMask, value interface{}) error {
	currentData := data
	for _, fieldName := range fieldNames[:len(fieldNames)-1] {
		// 获取当前字段对应的消息类型
		fieldDescriptor := currentData.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		if fieldDescriptor == nil {
			return fmt.Errorf("field %s not found", fieldName)
		} else {
			switch fieldDescriptor.Kind() { //字段类型
			case protoreflect.MessageKind: //消息类型
				// 获取当前字段的message
				fieldValue := currentData.Get(fieldDescriptor).Message()
				if fieldValue.IsValid() {
					currentData = fieldValue
				} else {
					return fmt.Errorf("field %s not found", fieldName)
				}
			default:
				return fmt.Errorf("field %s is not a valid sub-data", fieldName)
			}
		}
	}
	// 更新目标字段的值
	fieldDescriptor := currentData.Descriptor().Fields().ByName(protoreflect.Name(targetFieldName))
	if fieldDescriptor == nil {
		return fmt.Errorf("targetField %s not found", targetFieldName)
	} else {
		// 检查新值与目标字段的类型是否匹配
		newValue := protoreflect.ValueOf(value)
		// // TODO 类型校验
		// if !newValue.IsValid() {
		// 	return fmt.Errorf("new value is not compatible with field %s", targetFieldName)
		// }
		// msg, ok := newValue.Interface().(proto.Message)
		// if !ok {
		// 	return fmt.Errorf("new value is not compatible with field %s", targetFieldName)
		// }
		// if msg.ProtoReflect().Descriptor().FullName() != fieldDescriptor.Message().FullName() {
		// 	return fmt.Errorf("new value is not compatible with field %s", targetFieldName)
		// }
		currentData.Set(fieldDescriptor, newValue)
	}

	// 根据fieldMask进一步过滤需要更新的字段
	if fieldMask != nil && len(fieldMask.Paths) > 0 {
		filteredData := FilterData(data, fieldMask)
		if filteredData != nil {
			MergeData(data, filteredData)
		}
	}

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
func (ds *DataStore) GetMemoryData(data protoreflect.Message, fieldNames []string) ([]byte, error) {
	// 根据路径逐级遍历查找目标字段的值
	currentData := data
	for _, fieldName := range fieldNames {
		fieldDescriptor := currentData.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		if fieldDescriptor == nil {
			return nil, fmt.Errorf("field %s not found", fieldName)
		}
		if fieldDescriptor.Kind() == protoreflect.MessageKind { //字段类型是消息类型时
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
	// 假设path的格式为：nodeList[0].fileInfo[1].name
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

// 获取oneof的data
func GetOneOfDataValue(d *p.WsData) string {
	var dataValue string
	switch data := d.Data.(type) {
	case *p.WsData_StringData:
		dataValue = data.StringData
	case *p.WsData_Int32Data:
		// 处理 int32 类型的数据
	case *p.WsData_FloatData:
		// 处理 float 类型的数据
	case *p.WsData_BoolData:
		// 处理 bool 类型的数据
	case *p.WsData_BytesData:
		// 处理 bytes 类型的数据
	}
	return dataValue
}

func main() {
	// p.WsData 是一个消息类型 可以被视为 protoreflect.Message 类型的对象
	// 1.INIT
	wsDataInit := &p.WsData{
		MethodId: p.MethodId_DATA_INIT,
		Path:     "",
		FieldMask: &field_mask.FieldMask{
			Paths: []string{},
		},
		// Data: &p.WsData_BytesData{
		// 	BytesData: []byte(`{"nodeList": [{"fileInfo": [{"name": "1-0001"}, {"name": "1-0002"}]}, {"fileInfo": [{"name": "2-0001"}, {"name": "2-0002"}]}]}`),
		// },
	}
	ds := NewDataStore(wsDataInit.ProtoReflect())
	MemoryData = map[string]*DataStore{}
	MemoryData["space_1"] = ds //存储的是指向DataStore对象的指针
	// GetProtoMessage("", ds.data)

	// 2.UPDATE
	wsDataUpdate := &p.WsData{
		MethodId: p.MethodId_DATA_UPDATE,
		// Path:     "nodeList[0].fileInfo[1]",
		Path: "x",
		FieldMask: &field_mask.FieldMask{
			Paths: []string{"name"},
		},
		Data: &p.WsData_StringData{StringData: "new value 666"},
	}
	// 将对象序列化为字节流
	dataBytes, err := proto.Marshal(wsDataUpdate)
	if err != nil {
		log.Fatal(err)
	}
	// 反序列化字节流为对象
	receivedWsData := &p.WsData{}
	err = proto.Unmarshal(dataBytes, receivedWsData)
	if err != nil {
		log.Fatal(err)
	}

	var fieldNames []string
	var targetFieldName string
	path := receivedWsData.Path
	// 根据path解析出各级字段名
	fieldNames, err = ParsePath(path)
	if err != nil {
		log.Fatal(err)
	}
	// 获取要更新的目标字段名
	targetFieldName = fieldNames[len(fieldNames)-1]
	// 分配新的FieldMask对象
	if wsDataUpdate.FieldMask == nil {
		wsDataUpdate.FieldMask = &field_mask.FieldMask{}
	}

	value := GetOneOfDataValue(wsDataUpdate)
	err = MemoryData["space_1"].UpdateMemoryData(ds.data, fieldNames, targetFieldName, wsDataUpdate.FieldMask, value)
	if err != nil {
		log.Fatal(err)
	}

	// 获取更新后的值
	result, err := MemoryData["space_1"].GetMemoryData(ds.data, fieldNames)
	if err != nil {
		log.Fatal(err)
	}

	// 打印结果
	fmt.Println(string(result))
}

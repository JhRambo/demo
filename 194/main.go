package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	pb "demo/utils/proto/p2"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

type DataStore struct {
	data protoreflect.ProtoMessage
}

func NewDataStore(data protoreflect.ProtoMessage) *DataStore {
	return &DataStore{
		data: data,
	}
}

// 内存中保存的数据结构 全局变量
var MemoryData map[string]*DataStore

// 获取message
func (ds *DataStore) GetProtoMessage(key string) []byte {
	bys, _ := proto.Marshal(ds.data)
	return bys
}

// 新增message
func (ds *DataStore) AddProtoMessage(message protoreflect.Message, path string, data protoreflect.Message) error {
	paths := ParsePath1(path)
	for i := 0; i < len(paths); i++ {
		fieldName := paths[i]
		var field protoreflect.FieldDescriptor
		var fieldNames []string
		if strings.Contains(fieldName, "[") { // 数组
			fieldNames = ParsePath2(fieldName)
			field = message.Descriptor().Fields().ByName(protoreflect.Name(fieldNames[0]))
			if field == nil {
				return fmt.Errorf("field %s not found 1", fieldNames[0])
			}
		} else {
			field = message.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
			if field == nil {
				return fmt.Errorf("field %s not found 2", fieldName)
			}
		}

		if field.IsList() { // 数组
			list := message.Mutable(field).List() // 数组列表
			if list == nil {
				return fmt.Errorf("failed to create list for field %s", fieldName)
			}
			if len(fieldNames) > 0 {
				index, err := strconv.Atoi(fieldNames[1]) // 数组索引
				if err != nil {
					return fmt.Errorf("Failed to convert string to int: %v", err)
				}
				if index >= list.Len() {
					return fmt.Errorf("array index out of bounds: %d", index)
				}
				listIndexMessage := list.Get(index).Message() // 获取对应数组索引[x]的 protoreflect.Message
				if i == len(paths)-1 {
					// 最后一个字段为数组元素，直接设置
					msg := data.(protoreflect.Message)
					listIndexMessage.Set(listIndexMessage.Descriptor().Fields().ByName(protoreflect.Name(fieldNames[0])), protoreflect.ValueOf(msg))
				} else {
					// 最后一个字段为数组元素的子字段，递归处理
					if !listIndexMessage.IsValid() {
						return fmt.Errorf("invalid sub-message at %s[%d]", fieldName, index)
					}
					if err := ds.AddProtoMessage(listIndexMessage, strings.Join(paths[i+1:], "."), data); err != nil {
						return err
					}
				}
			} else {
				list.Append(protoreflect.ValueOf(data))
			}
		} else {
			return fmt.Errorf("invalid path %s", fieldName)
		}
	}

	return nil
}

// 更新message
func (ds *DataStore) UpdateProtoMessage(message protoreflect.Message, path string, data interface{}) error {
	paths := ParsePath1(path)
	targetFieldName := paths[len(paths)-1] //最后一个元素即目标字段
	// 遍历每个路径
	for i := 0; i < len(paths); i++ {
		fieldName := paths[i]
		var field protoreflect.FieldDescriptor
		var fieldNames []string
		if strings.Contains(fieldName, "[") { //数组
			// 先获取数组字段名称
			fieldNames = ParsePath2(fieldName)
			// 再获取当前字段描述符
			field = message.Descriptor().Fields().ByName(protoreflect.Name(fieldNames[0]))
			if field == nil {
				return fmt.Errorf("field %s not found 1", fieldNames[0])
			}
		} else {
			// 直接获取当前字段描述符
			field = message.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
			if field == nil {
				return fmt.Errorf("field %s not found 2", fieldName)
			}
		}

		if field.IsList() { // 数组
			index, err := strconv.Atoi(fieldNames[1]) //数组索引
			if err != nil {
				return fmt.Errorf("Failed to convert string to int: %v", err)
			}
			list := message.Get(field).List() //数组列表
			if index >= list.Len() {
				return fmt.Errorf("array index out of bounds: %d", index)
			}
			listIndexMessage := list.Get(index).Message() //获取对应数组索引[x]的protoreflect.Message
			if i == len(paths)-1 {
				// 最后一个字段为数组元素，直接设置
				t := reflect.TypeOf(data)
				if t.Kind() == reflect.Ptr {
					list.Set(index, protoreflect.ValueOf(data)) //直接替换当前索引值
				} else {
					aField := listIndexMessage.Descriptor().Fields().ByName(protoreflect.Name(fieldNames[0]))
					if aField == nil {
						return fmt.Errorf("target field %s not found", fieldNames[0])
					} else {
						listIndexMessage.Set(aField, protoreflect.ValueOf(data))
					}
				}
			} else {
				// 最后一个字段为数组元素的子字段，递归处理
				if !listIndexMessage.IsValid() {
					return fmt.Errorf("invalid sub-message at %s[%d]", fieldName, index)
				}
				if err := ds.UpdateProtoMessage(listIndexMessage, strings.Join(paths[i+1:], "."), data); err != nil {
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
					if err := ds.UpdateProtoMessage(subMessage, "", data); err != nil {
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

func main() {
	// 1.INIT
	info1 := []*pb.ExtendInfo{}
	info2 := []*pb.ExtendInfo{}
	info1 = append(info1, &pb.ExtendInfo{
		Name: "拓展信息1",
	}, &pb.ExtendInfo{
		Name: "拓展信息2",
	})
	info2 = append(info2, &pb.ExtendInfo{
		Name: "拓展信息1111",
	}, &pb.ExtendInfo{
		Name: "拓展信息2222",
	})
	list := []*pb.List{}
	// list = append(list, &pb.List{
	// 	Name:       "节点1",
	// 	Desc:       "节点1描述",
	// 	ExtendInfo: info1,
	// }, &pb.List{
	// 	Name:       "节点2",
	// 	Desc:       "节点2描述",
	// 	ExtendInfo: info2,
	// })
	// dataInit := &pb.Node{
	// 	Data: list,
	// }

	list = append(list, &pb.List{
		Name:       "节点1",
		Desc:       "节点1描述",
		ExtendInfo: info1,
	})
	dataInit := &pb.Node{
		// Data: list,
	}
	bys, _ := proto.Marshal(dataInit)
	wsDataInit := &pb.WsData{
		Path: "",
		FieldMask: &field_mask.FieldMask{
			Paths: []string{},
		},
		Data: bys,
	}

	// 获取Message描述符
	desc := dataInit.ProtoReflect().Descriptor()
	newMsg, err := CreateNewMessage(wsDataInit.Data, desc)
	if err != nil {
		// 处理错误
		fmt.Printf("err: %v", err)
		return
	}

	ds := NewDataStore(newMsg)
	MemoryData = map[string]*DataStore{}
	MemoryData["space_1"] = ds
	// // 更新message
	// // 1.值类型
	// ds.UpdateProtoMessage(MemoryData["space_1"].data.ProtoReflect(), "data[0].extendInfo[1].name", "xxxxxx")
	// // 2.引用类型
	// extendInfo := &pb.ExtendInfo{
	// 	Name: "1234566666666666666666666",
	// }
	// ds.UpdateProtoMessage(MemoryData["space_1"].data.ProtoReflect(), "data[0].extendInfo[0]", extendInfo.ProtoReflect())

	// // 3.动态新增message
	// extendInfo := &pb.ExtendInfo{
	// 	Name: "1234566666666666666666666",
	// }
	// ds.AddProtoMessage(MemoryData["space_1"].data.ProtoReflect(), "data[0].extendInfo", extendInfo.ProtoReflect())

	// 4.动态新增message
	nodeInfo1 := []*pb.ExtendInfo{}
	nodeInfo2 := []*pb.ExtendInfo{}
	nodeInfo1 = append(nodeInfo1, &pb.ExtendInfo{
		Name: "动态node拓展信息1",
	}, &pb.ExtendInfo{
		Name: "动态node拓展信息2",
	})
	node1 := &pb.List{
		Name:       "动态node",
		Desc:       "动态node描述666666",
		ExtendInfo: nodeInfo1,
	}
	ds.AddProtoMessage(MemoryData["space_1"].data.ProtoReflect(), "data", node1.ProtoReflect())
	node2 := &pb.List{
		Name:       "动态node",
		Desc:       "动态node描述777777",
		ExtendInfo: nodeInfo2,
	}
	ds.AddProtoMessage(MemoryData["space_1"].data.ProtoReflect(), "data", node2.ProtoReflect())
	// 获取message
	proto.Unmarshal(ds.GetProtoMessage("space_1"), newMsg)
	fmt.Println(newMsg)
}

// func CreateNewMessage(data []byte, messageType reflect.Type) (proto.Message, error) {
// 	// 创建新的空消息对象
// 	newMsg := reflect.New(messageType).Interface().(proto.Message)

// 	// 反序列化数据到新的消息对象
// 	err := proto.Unmarshal(data, newMsg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return newMsg, nil
// }

func CreateNewMessage(data []byte, desc protoreflect.MessageDescriptor) (proto.Message, error) {
	if desc == nil {
		return nil, fmt.Errorf("message descriptor is nil")
	}
	newMsg := dynamicpb.NewMessage(desc)

	// 反序列化数据到新的消息对象
	err := proto.Unmarshal(data, newMsg)
	if err != nil {
		return nil, err
	}

	return newMsg, nil
}

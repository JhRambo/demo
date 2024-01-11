package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"demo/utils/proto/p"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DataStore struct {
	dataMap map[string]interface{}
}

func NewDataStore() *DataStore {
	return &DataStore{
		dataMap: make(map[string]interface{}),
	}
}

var MemoryData map[string]interface{}

func (ds *DataStore) UpdateOrInsertData(path string, fieldMask *field_mask.FieldMask, data interface{}) {
	value := reflectFieldValue(data)
	if value == nil {
		log.Println("Invalid data type")
		return
	}

	ds.updateDataAtPath(path, fieldMask, value)
}

func (ds *DataStore) GetData(path string) interface{} {
	return ds.getDataAtPath(path)
}

func (ds *DataStore) updateDataAtPath(path string, fieldMask *field_mask.FieldMask, value interface{}) {
	currentData, ok := ds.dataMap[path]
	if !ok {
		ds.dataMap[path] = value
	} else {
		protoMsg, ok := currentData.(proto.Message)
		if !ok {
			log.Println("Invalid data type")
			return
		}

		msgDesc := protoMsg.ProtoReflect().Descriptor()
		msg := protoMsg.ProtoReflect()
		for _, fieldPath := range fieldMask.Paths {
			fieldDesc := msgDesc.Fields().ByName(protoreflect.Name(fieldPath))
			if fieldDesc == nil {
				log.Printf("Field '%s' not found\n", fieldPath)
				continue
			}

			fieldValue := reflectFieldValue(value)
			if fieldValue == nil {
				log.Printf("Invalid value type for field '%s'\n", fieldPath)
				continue
			}

			fieldValueProto := protoreflect.ValueOf(fieldValue)
			if fieldValueProto.Interface() == nil {
				log.Printf("Invalid value type for field '%s'\n", fieldPath)
				continue
			}

			switch fieldDesc.Kind() {
			case protoreflect.BoolKind:
				if v, ok := fieldValueProto.Interface().(bool); ok {
					msg.Set(fieldDesc, protoreflect.ValueOfBool(v))
				}
			case protoreflect.EnumKind:
				if v, ok := fieldValueProto.Interface().(protoreflect.EnumNumber); ok {
					msg.Set(fieldDesc, protoreflect.ValueOfEnum(v))
				}
				// 处理其他字段类型...
			}
		}
	}
}

func (ds *DataStore) getDataAtPath(path string) interface{} {
	// return ds.dataMap[path]
	return MemoryData["1_1"]
}

func reflectFieldValue(data interface{}) interface{} {
	switch v := data.(type) {
	case int32, int64, float32, float64, string, bool, []byte:
		return v
	case *int32, *int64, *float32, *float64, *string, *bool, *[]byte:
		return reflectFieldValue(reflectValue(v))
	default:
		if protoMsg, ok := v.(proto.Message); ok {
			return protoMsg.ProtoReflect().Interface()
		}
	}
	return nil
}

func reflectValue(value interface{}) reflect.Value {
	if value == nil {
		return reflect.ValueOf(nil)
	}

	return reflect.ValueOf(value).Elem()
}

func main() {
	// 示例数据
	data := &p.WsData{
		MethodId: p.MethodId_DATA_UPDATE,
		Path:     "Scene.NodeList[0].FileInfo[1].name",
		FieldMask: &field_mask.FieldMask{
			Paths: []string{},
		},
		Data: &p.WsData_StringData{
			StringData: "zhangs",
		},
	}

	// 反序列化示例数据
	dataBytes, err := proto.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	wsData := &p.WsData{}
	if err := proto.Unmarshal(dataBytes, wsData); err != nil {
		log.Fatal(err)
	}

	// 创建数据存储对象
	store := NewDataStore()

	// 更新或新增数据
	store.UpdateOrInsertData(wsData.Path, wsData.FieldMask, wsData)

	MemoryData = map[string]interface{}{}
	MemoryData["1_1"] = wsData.Data

	for {
		time.Sleep(5 * time.Second)
		// 获取数据
		result := store.GetData(wsData.Path)
		fmt.Println("Result:", result)
	}

}

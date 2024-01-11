package main

import (
	"fmt"
	"log"
	"strings"

	"demo/utils/proto/p"

	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/protobuf/field_mask"
)

// 内存中保存的数据结构 全局变量
var MemoryData map[string]interface{}

// 更新内存中的值
func UpdateMemoryData(data map[string]interface{}, path string, fieldMask *field_mask.FieldMask, value interface{}) error {
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
		fieldValue, ok := currentData[fieldName]
		if !ok {
			// 如果字段不存在，则创建一个空的子数据
			subData := make(map[string]interface{})
			currentData[fieldName] = subData
			currentData = subData
		} else {
			switch fieldValue := fieldValue.(type) {
			case map[string]interface{}:
				currentData = fieldValue
			case []interface{}:
				// 遍历切片查找目标字段
				for _, subData := range fieldValue {
					subDataMap, ok := subData.(map[string]interface{})
					if !ok {
						return fmt.Errorf("field %s is not a valid sub-data", fieldName)
					}
					subFieldValue, ok := subDataMap[targetFieldName]
					if ok {
						subFieldValue = value
					}
					subDataMap[targetFieldName] = subFieldValue
				}
				return nil
			default:
				return fmt.Errorf("field %s is not a valid sub-data", fieldName)
			}
		}
	}

	// 更新目标字段的值
	currentData[targetFieldName] = value

	// 根据fieldMask进一步过滤需要更新的字段
	filteredData := FilterData(data, fieldMask)
	data = filteredData

	return nil
}

// 根据path获取内存中的值
func GetMemoryData(data map[string]interface{}, path string) (interface{}, error) {
	// 根据path解析出各级字段名
	fieldNames, err := ParsePath(path)
	if err != nil {
		return nil, err
	}

	// 根据路径逐级遍历查找目标字段的值
	currentData := data
	for _, fieldName := range fieldNames {
		fieldValue, ok := currentData[fieldName]
		if !ok {
			return nil, fmt.Errorf("field %s not found", fieldName)
		}
		subData, ok := fieldValue.(map[string]interface{})
		if !ok {
			return fieldValue, nil
		}
		currentData = subData
	}

	return currentData, nil
}

// 根据fieldMask过滤数据
func FilterData(data map[string]interface{}, fieldMask *field_mask.FieldMask) map[string]interface{} {
	filteredData := make(map[string]interface{})

	for _, path := range fieldMask.Paths {
		fieldNames, err := ParsePath(path)
		if err != nil {
			continue
		}

		currentData := data
		for _, fieldName := range fieldNames {
			fieldValue, ok := currentData[fieldName]
			if !ok {
				break
			}
			subData, ok := fieldValue.(map[string]interface{})
			if !ok {
				filteredData[fieldName] = fieldValue
				break
			}
			currentData = subData
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

	// 模拟接收到WsData的字节流

	// 反序列化字节流为WsData对象
	receivedWsData := &p.WsData{}
	err = proto.Unmarshal(dataBytes, receivedWsData)
	if err != nil {
		log.Fatal(err)
	}

	MemoryData = map[string]interface{}{}

	// 根据接收到的WsData更新内存中的数据
	path := receivedWsData.Path
	fieldMask := receivedWsData.FieldMask
	value := receivedWsData.Data
	err = UpdateMemoryData(MemoryData, path, fieldMask, value)
	if err != nil {
		log.Fatal(err)
	}

	// 获取更新后的值
	result, err := GetMemoryData(MemoryData, path)
	if err != nil {
		log.Fatal(err)
	}

	// 打印结果
	fmt.Println(result)
	fmt.Println(string(result.([]byte)))
}

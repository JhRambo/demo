package main

import (
	"demo/logs"
	"fmt"
	"reflect"
	"strings"

	"demo/utils/proto/nodestruct"
	"demo/utils/proto/person"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var messageRegistry map[nodestruct.MessageId]reflect.Type
var SceneMgr map[string]*Scene

type Scene struct {
	NodeList protoreflect.Message
}

func main() {
	update()
	// req := &nodestruct.WsData{
	// 	Data: &nodestruct.Payload{
	// 		Path: "nodeList.transformInfo.position.x",
	// 	},
	// }
	// message, _ := proto.Marshal(req)
	// sceneMgr := GetSpaceNodeManager(1, int32(1))
	// wsData := &nodestruct.WsData{}
	// err := proto.Unmarshal(message, wsData)
	// if err != nil {
	// 	logs.Error("解析二进制数据时出错:", err)
	// 	return
	// }
	// if wsData.Data == nil {
	// 	return
	// }
	// res, err := sceneMgr.GetData(wsData.Data.Path)
	// if err != nil {
	// 	logs.Error(err)
	// 	return
	// }
	// if res == nil {
	// 	logs.Info("data is nil")
	// 	return
	// }
	// _, err = proto.Marshal(res.Interface())
	// if err != nil {
	// 	logs.Error(err)
	// 	return
	// }
}

func update() {
	sceneMgr := GetSpaceNodeManager(1, 1)
	p := &person.Person{
		Name: "张三",
	}
	bys, _ := proto.Marshal(p)
	wsData := &nodestruct.WsData{
		Data: &nodestruct.Payload{
			Path: "nodeList.transformInfo.position.x",
			Data: bys,
		},
	}
	tmpData := ConvertDataToValue(wsData.Data)
	err := sceneMgr.UpdateData(wsData.Data.Path, tmpData, nodestruct.MethodId_SCENE_UPDATE)
	if err != nil {
		logs.Error(err)
		return
	}
}

func init() {
	logs.Init("", "", 5, false)
	// 初始化消息ID和消息结构的映射关系
	messageRegistry = make(map[nodestruct.MessageId]reflect.Type)
	registerMessage(nodestruct.MessageId_MSG_SCENE, &nodestruct.Scene{})
	registerMessage(nodestruct.MessageId_MSG_MODEL_MODULE, &nodestruct.ModelModule{})
	registerMessage(nodestruct.MessageId_MSG_TRANSFORM, &nodestruct.Transform{})
	registerMessage(nodestruct.MessageId_MSG_VECTOR, &nodestruct.Vector{})
	registerMessage(nodestruct.MessageId_MSG_FILE_INFO, &nodestruct.FileInfo{})
}

func GetSpaceNodeManager(configId, sceneType int32) *Scene {
	key := fmt.Sprintf("%d_%d", configId, sceneType)
	if obj, ok := SceneMgr[key]; ok {
		return obj
	} else {
		return InitSpace(configId, sceneType)
	}
}

// 初始化/创建场景资源
func InitSpace(configId, sceneType int32) *Scene {
	SceneModel := new(Scene)
	initData := nodestruct.Scene{NodeList: nil}
	initData.NodeList = make(map[int32]*nodestruct.Payload)
	SceneModel.NodeList = initData.ProtoReflect()
	SceneMgr = make(map[string]*Scene)
	SceneMgr[fmt.Sprintf("%d_%d", configId, sceneType)] = SceneModel
	return SceneModel
}

// 递归函数 getFieldValue，用于获取指定字段的值
func getFieldValue(current protoreflect.Message, fields []protoreflect.Name) protoreflect.Message {
	end := len(fields) - 1
	for index, field := range fields {
		if index > 0 {
			//如果索引大于 0，表示当前字段不是根字段，则返回 nil
			return nil
		}
		key := int32(0)
		file1 := current.Descriptor().Fields().ByName(field) //通过字段名从当前消息对象的描述符中获取字段的描述符
		//判断当前消息对象的类型是否为 "Payload"
		if current.Type().Descriptor().FullName() == "Payload" {
			fieldPayloadData := current.Descriptor().Fields().ByName("msgId")
			msgId := current.Get(fieldPayloadData).Interface().(protoreflect.EnumNumber)
			fieldPayloadMsgId := current.Descriptor().Fields().ByName("data")
			data := current.Get(fieldPayloadMsgId).Bytes()
			tmpData, err := CreateMessage(nodestruct.MessageId(msgId))
			if err != nil {
				logs.Error(err)
				return nil
			}
			proto.Unmarshal(data, tmpData)
			current = tmpData.ProtoReflect()
			file1 = current.Descriptor().Fields().ByName(field)
		}
		if current.Has(file1) {
			if index == end {
				if file1.IsMap() {
					res := current.Get(file1).Map().Get(protoreflect.MapKey(protoreflect.ValueOf(key))).Message()
					return res
				} else {
					tmCommon := &nodestruct.CommonData{}
					switch file1.Kind() {
					case protoreflect.Int32Kind, protoreflect.Int64Kind:
						tmCommon.Data = &nodestruct.CommonData_Int32Data{Int32Data: int32(current.Get(file1).Int())}
					case protoreflect.FloatKind:
						tmCommon.Data = &nodestruct.CommonData_FloatData{FloatData: float32(current.Get(file1).Float())}
					case protoreflect.StringKind:
						tmCommon.Data = &nodestruct.CommonData_StringData{StringData: current.Get(file1).String()}
					case protoreflect.BytesKind:
						tmCommon.Data = &nodestruct.CommonData_BytesData{BytesData: current.Get(file1).Bytes()}
					default:
						return current.Get(file1).Message()
					}
					return tmCommon.ProtoReflect()
				}
			} else {
				if file1.IsMap() {
					tmp := current.Get(file1).Map().Get(protoreflect.MapKey(protoreflect.ValueOf(key))).Message()
					return getFieldValue(tmp, fields[index+1:])
				} else {
					return getFieldValue(current.Get(file1).Message(), fields[index+1:])
				}
			}
		} else {
			return nil
		}
	}
	return nil
}

func (s *Scene) GetData(path string) (protoreflect.Message, error) {
	fields := getPathFields(path)
	res := getFieldValue(s.NodeList, fields)
	if res != nil && res.Interface() != nil {
		logs.Debug(">>>>>>>>:", res.Interface())
	}
	return res, nil
}

func (s *Scene) UpdateData(path string, data protoreflect.Value, methodId nodestruct.MethodId) error {
	fields := getPathFields(path)
	updateFieldValue(&s.NodeList, fields, data, methodId)
	return nil
}

// 递归算法遍历消息对象的字段，并根据字段路径和需要执行的操作类型来更新字段值
func updateFieldValue(msgData *protoreflect.Message, fields []protoreflect.Name, newValue protoreflect.Value, methodId nodestruct.MethodId) {
	logs.Info("methodId：", methodId)
	var current protoreflect.Message
	current = *msgData
	end := len(fields) - 1
	for index, field := range fields {
		if index > 0 {
			return
		}
		key := int32(0)
		file1 := current.Descriptor().Fields().ByName(field)
		var playLoay protoreflect.Message
		newMsg := false
		if current.Type().Descriptor().FullName() == "Payload" {
			newMsg = true
			fieldPayloadMsgId := current.Descriptor().Fields().ByName("msgId")
			msgId := current.Get(fieldPayloadMsgId).Interface().(protoreflect.EnumNumber)
			fieldPayloadData := current.Descriptor().Fields().ByName("data")
			data := current.Get(fieldPayloadData).Bytes()
			tmpData, err := CreateMessage(nodestruct.MessageId(msgId))
			if err != nil {
				logs.Error(err)
				return
			}
			proto.Unmarshal(data, tmpData)
			playLoay = tmpData.ProtoReflect()
			file1 = playLoay.Descriptor().Fields().ByName(field)
		} else {
			playLoay = current
		}
		if playLoay.Has(file1) {
			if index == end {
				if file1.IsMap() {
					mapValue := playLoay.Get(file1).Map()
					switch methodId {
					case nodestruct.MethodId_MAP_ADD, nodestruct.MethodId_SCENE_UPDATE:
						mapValue.Set(protoreflect.MapKey(protoreflect.ValueOf(key)), newValue)
					case nodestruct.MethodId_ARRAY_DELETE, nodestruct.MethodId_MAP_DELETE:
						logs.Debug("删除：", key)
						mapValue.Clear(protoreflect.MapKey(protoreflect.ValueOf(key)))
					}
				} else {
					playLoay.Set(file1, newValue)
				}
			} else {
				if !playLoay.Get(file1).IsValid() {
					data := protoreflect.Value{}
					playLoay.Set(file1, protoreflect.ValueOfMessage(data.Message()))
				}
				if file1.IsMap() {
					var msg protoreflect.Message
					msg = playLoay.Get(file1).Map().Get(protoreflect.MapKey(protoreflect.ValueOf(key))).Message()
					updateFieldValue(&msg, fields[index+1:], newValue, methodId)
				} else {
					var msg protoreflect.Message
					msg = playLoay.Get(file1).Message()
					updateFieldValue(&msg, fields[index+1:], newValue, methodId)
				}
			}
		} else {
			if index == end {
				if file1.IsMap() {
					if methodId == nodestruct.MethodId_ARRAY_DELETE {
						return
					}
					if !current.Get(file1).Map().IsValid() {
						// 初始化
						nodeMap := current.Mutable(playLoay.Descriptor().Fields().ByName(field)).Map()
						nodeMap.Set(protoreflect.MapKey(protoreflect.ValueOf(key)), newValue)
					}
				} else {
					playLoay.Set(file1, newValue)
				}
			}
		}
		if newMsg {
			fieldPayloadData := current.Descriptor().Fields().ByName("data")
			playData, _ := proto.Marshal(playLoay.Interface())
			current.Set(fieldPayloadData, protoreflect.ValueOf(playData))
		}
	}
}

// getPathFields splits the provided path into individual field names
func getPathFields(path string) []protoreflect.Name {
	var fields []protoreflect.Name
	for _, fieldName := range splitPath(path) {
		fields = append(fields, protoreflect.Name(fieldName))
	}
	return fields
}

// splitPath splits a path into individual field names
func splitPath(path string) []string {
	return strings.Split(path, ".")
}

func CreateMessage(msgID nodestruct.MessageId) (proto.Message, error) {
	msgType, ok := messageRegistry[msgID]
	if !ok {
		return nil, fmt.Errorf("Unknown Message ID: %v", msgID)
	}
	return reflect.New(msgType.Elem()).Interface().(proto.Message), nil
}

func registerMessage(msgID nodestruct.MessageId, msgType interface{}) {
	messageRegistry[msgID] = reflect.TypeOf(msgType)
}

func ConvertDataToValue(data interface{}) protoreflect.Value {
	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Int:
		return protoreflect.ValueOf(data.(int32))
	case reflect.Float32:
		return protoreflect.ValueOf(data.(float32))
	default:
		return protoreflect.ValueOf((data.(proto.Message)).ProtoReflect())
	}
}

package main

func main() {
	// // 创建
	// resp, err := utils.Create()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段1 操作值类型或object类型 新增字段 ok
	// path := []string{"nodeList.node1.ooxx"}
	// data := []interface{}{"造型哦ing"}
	// eid := int64(1000)
	// action := []string{""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段2 操作值类型或object类型 新增字段 ok
	// path := []string{"nodeList.node1.ooxx", "nodeList.node1.xyz"}
	// data := []interface{}{"1111", "9999"}
	// eid := int64(1000)
	// action := []string{"", ""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段3 操作值类型或object类型 修改值内容 ok
	// path := []string{"nodeList.node1.ooxx"}
	// data := []interface{}{"ssssss"}
	// eid := int64(1000)
	// action := []string{""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段4 操作值类型或object类型 修改值内容 ok
	// path := []string{"nodeList.node1.ooxx", "nodeList.node1.xyz"}
	// data := []interface{}{"[11,00]", "[\"xx\",\"yy\"]"}
	// eid := int64(1000)
	// action := []string{"", ""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段5 操作值类型或object类型 删除字段 ok
	// path := []string{"nodeList.node1.ooxx"}
	// data := []interface{}{""}
	// eid := int64(1000)
	// action := []string{"d"}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段6 操作数组 修改值内容 ok
	// path := []string{"nodeList.node2.slices.0.a", "nodeList.node2.slices.0.b"}
	// data := []interface{}{"111111", "222222"}
	// action := []string{"", ""}
	// eid := int64(1000)
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段7 操作数组 修改值内容 ok
	// path := []string{"nodeList.node2.slices.1.a", "nodeList.node2.slices.2.a"}
	// data := []interface{}{"张三李四王五老六", "赵谦孙俪"}
	// eid := int64(1000)
	// action := []string{"", ""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段8 操作数组 新增字段 ok
	// path := []string{"nodeList.node2.slices.0.c", "nodeList.node2.slices.1.c", "nodeList.node2.slices.2.c"}
	// data := []interface{}{"x", "y", "z"}
	// eid := int64(1000)
	// action := []string{"", "", ""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段9 操作数组 删除字段 ok
	// path := []string{"nodeList.node2.slices.0.c", "nodeList.node2.slices.1.c", "nodeList.node2.slices.2.c"}
	// data := []interface{}{"", "", ""}
	// eid := int64(1000)
	// action := []string{"d", "d", "d"}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段10 操作数组或值类型或object类型 不同类型同时更新字段 ok
	// path := []string{"nodeList.node2.slices.0.a", "nodeList.node1.xyz"}
	// data := []interface{}{"aaaaaaaaaa", "2222222222"}
	// eid := int64(1000)
	// action := []string{"", ""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段11 操作数组 调整顺序 TODO
	// //1.先删除
	// path := []string{"nodeList.node2.slices.0.a", "nodeList.node2.slices.0.b"}
	// data := []interface{}{"", ""}
	// eid := int64(1000)
	// action := []string{"d", "d"}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)
	// //2.再新增
	// path2 := []string{"nodeList.node2.slices.0.a", "nodeList.node2.slices.0.b"}
	// data2 := []interface{}{"1111111111", "2222222222"}
	// eid2 := int64(1000)
	// action2 := []string{"", ""}
	// configId2 := int64(1)
	// spaceId2 := int64(1)
	// resp2, err2 := utils.Update(path2, data2, action2, configId2, spaceId2, eid2)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// fmt.Printf("%#v", resp2)
}

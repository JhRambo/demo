package main

import (
	"fmt"
	"sync"
)

// sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量
// sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能
// 通过 read 和 dirty 两个字段将读写分离，读的数据存在只读字段 read 上，将最新写入的数据则存在 dirty 字段上。
// 读取时会先查询 read，不存在再查询 dirty，写入时则只写入 dirty。
// 读取 read 并不需要加锁，而读或写 dirty 都需要加锁。
// 另外有 misses 字段来统计 read 被穿透的次数（被穿透指需要读 dirty 的情况），超过一定次数则将 dirty 数据同步到 read 上。
// 对于删除数据则直接通过标记来延迟删除。
func main() {

	var scene sync.Map

	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london")) //100 true

	scene.Delete("c")
	// 参数是一对key：value
	// 如果该key存在且没有被标记删除，则返回原先的value和true
	// 不存在，则store，返回该value 和false
	v, ok := scene.LoadOrStore("c", 1) //c不存在
	fmt.Println("11111", v, ok)        //11111 1 false
	// scene.Delete("c")                    //被标记删除
	vv, okk := scene.LoadOrStore("c", 2) //c已存在
	fmt.Println("22222", vv, okk)        //22222 1 true

	// 根据键删除对应的键值对
	scene.Delete("london")

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool { //随机非顺序打印
		fmt.Println("range:", k, v)
		return true
	})
}

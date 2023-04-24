package main

import (
	"context"
	"fmt"
	"time"
)

// 获取键值对的过程也是层层向上调用直到【首次】设置key的父节点
// 如果没有找到首次设置key的父节点，会向上遍历直到根节点
// 如果根节点找到了key就会返回，否则就会找到最终的emptyCtx返回nil
func main() {
	ctx := context.WithValue(context.Background(), "key", "value01")
	_ = context.WithValue(ctx, "key", "value02")
	_ = context.WithValue(ctx, "key", "value03")
	fmt.Println(ctx.Value("key"))
	time.Sleep(1 * time.Nanosecond)
}

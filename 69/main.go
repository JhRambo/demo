package main

import (
	"context"
	"fmt"
	"time"
)

// 1.Context接口
// Deadline — 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
// Done — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消之后关闭，多次调用 Done 方法会返回同一个 Channel；
// Err — 返回 context.Context 结束的原因，它只会在 Done 返回的 Channel 被关闭时才会返回非空的值；如果 context.Context 被取消，会返回 Canceled 错误；如果 context.Context 超时，会返回 DeadlineExceeded 错误；
// Value — 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；

// 2.创建context
// context包允许以下方式创建和获得context:
// context.Background()：这个函数返回一个空context。这只能用于高等级（在 main 或顶级请求处理中）。
// context.TODO()：这个函数也是创建一个空context。也只能用于高等级或当您不确定使用什么 context，或函数以后会更新以便接收一个 context 。这意味您（或维护者）计划将来要添加 context 到函数。

// HelloHandle函数并没有进入超时的select分支
// 但是main函数的select却会等待context.Context的超时并打印出Hello Handle context deadline exceeded。
func HelloHandle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err()) //context deadline exceeded
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go HelloHandle(ctx, 2000*time.Millisecond) //500 2000 处理请求时间
	select {
	case <-ctx.Done():
		fmt.Println("Hello Handle ", ctx.Err())
	}
}

package main

import (
	"context"
	"fmt"
)

type key string

func Get(ctx context.Context, k key) {
	if v, ok := ctx.Value(k).(string); ok {
		fmt.Println(v)
	}
}

func main() {
	ctx := context.WithValue(context.Background(), key("lineshen"), "tencent")
	Get(ctx, key("lineshen"))
	Get(ctx, key("line"))
}

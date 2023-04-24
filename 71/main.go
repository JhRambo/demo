package main

import (
	"context"
	"fmt"
	"time"
)

func Monitor(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err()) //context deadline exceeded
	case <-time.After(2 * time.Second):
		fmt.Println("stop monitor")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	Monitor(ctx)
}

package main

import (
	"context"
	"fmt"
	"time"
)

func Monitor(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("stop monitor")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go Monitor(ctx)
	time.Sleep(5 * time.Second)
}

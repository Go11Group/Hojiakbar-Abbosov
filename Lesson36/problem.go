package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	go commonTask(ctx, "parameter1", "parameter2")

	time.Sleep(5 * time.Second)
}

func commonTask(ctx context.Context, param1, param2 string) {
	childCtx, cancel := context.WithDeadline(ctx, time.Now().Add(2 * time.Second))
	defer cancel()

	go childTask(childCtx, param1, param2)

	select {
	case <-childCtx.Done():
		fmt.Println("Parent context canceled or deadline exceeded")
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed within time limit")
	}
}

func childTask(ctx context.Context, param1, param2 string) {
	select {
	case <-ctx.Done():
		fmt.Println("Child context canceled or deadline exceeded")
		return
	case <-time.After(1 * time.Second):
		fmt.Println("Child task completed with parameters:", param1, param2)
	}
}

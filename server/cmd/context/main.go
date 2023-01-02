package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct {
}

func main() {
	ctx := context.WithValue(context.Background(), paramKey{}, "ang")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	mainTask(ctx)
}

func mainTask(ctx context.Context) {
	fmt.Printf("mainTask started with param: %q\n", ctx.Value(paramKey{}))
	c1, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	subTask(c1, "task1", 4*time.Second)
	subTask(ctx, "task2", 2*time.Second)
}

func subTask(ctx context.Context, name string, d time.Duration) {
	fmt.Printf("subTask %s started with param: %q\n", name, ctx.Value(paramKey{}))
	select {
	case <-ctx.Done():
		fmt.Printf("subTask %s timed out with param: %q\n", name, ctx.Value(paramKey{}))
	case <-time.After(d):
		fmt.Printf("subTask %s done with param: %q\n", name, ctx.Value(paramKey{}))
	}
}

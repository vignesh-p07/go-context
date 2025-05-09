package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Base context — never cancelled
	baseCtx := context.Background()

	// 1. context.WithValue — attach value
	ctxWithValue := context.WithValue(baseCtx, "userID", 12345)

	// 2. context.WithCancel — manually cancel
	ctxCancel, cancel := context.WithCancel(ctxWithValue)

	// 3. context.WithTimeout — auto cancel after time
	ctxTimeout, timeoutCancel := context.WithTimeout(ctxCancel, 2*time.Second)
	defer timeoutCancel() // always call cancel on timeout context

	// Start a goroutine that uses the context
	go doWork(ctxTimeout)

	// Simulate some main work
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Canceling context manually...")
	cancel() // manually cancel the context
	time.Sleep(2 * time.Second)
}

func doWork(ctx context.Context) {
	// Get value from context
	userID := ctx.Value("userID")
	fmt.Println("Worker: User ID from context:", userID)

	for {
		select {
		case <-ctx.Done():
			// Context cancelled or timed out
			fmt.Println("Worker: Context canceled:", ctx.Err())
			return
		default:
			fmt.Println("Worker: Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

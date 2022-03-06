package main

import (
	"context"
	"log"
	"time"
)

func main() {
	// a_create a parent context or get it from http.Request
	ctx := context.Background()

	// a_create a context with timeout of 2 seconds
	timeout, stop := context.WithTimeout(ctx, time.Second*2)
	defer stop()

	// call the longRunningOperation
	if err := longRunningOperation(timeout); err != nil {
		log.Println(err)
	}
}

func longRunningOperation(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond * 20)
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
	log.Println("done")
	return nil
}

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	longRunningOperation(ctx)
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	longRunningOperation(ctx)
	fmt.Fprintf(w, "HeloWorld Function")
}

func longRunningOperation(ctx context.Context) {
	time.Sleep(time.Second * 10)
}

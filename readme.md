# Understanding context in go

This repository intends to explain how timeout and cancellations can be handled using`context` package in go. This repository supports the 
examples discussed in [this](https://prakharsrivastav.com/posts/golang-context-and-cancellation/) blog post.

## Creating new context

- context.Background() : Use when creating fresh default context.
- context.Todo() : For future use.

## Deriving context from an existing context

- context.WithCancel(parentCtx) : Provides a new context with a CancelFunc. CancelFunc can be used for explicit cancellation.
- context.WithDeadline(parentCtx) : Provides a new context that will automatically cancel after a durations. Also provides CancelFunc for explicit cancellation.
- context.WithTimeout(parentCtx) : Provides a new context that will expire at a given timestamp. Also provides CancelFunc for explicit cancellation.

## Cancelling a context

1. cancelling explicitly

```go
func main() {
    _,cancel := context.WithCancel(parentCtx)
    // call the cancel to cancel the context
    cancel()
}
```

2. cancelling after few seconds
```go
func main() {
    // cancel the context in 10 seconds
    _,cancel := context.WithTimeout(parentCtx, time.Second * 10) 
    
    // cancel explicitly (when function returns)
    defer cancel()  
}
```

3. cancelling at a given time
```go
func main() {
    // cancel the context at given time
    _,cancel := context.WithDeadline(parentCtx, time.Now().Add(2 * time.Second)) 

    // cancel explicitly (when function returns)
    defer cancel()
}
```
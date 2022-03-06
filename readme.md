# Understanding context in go

This repository intends to explain how timeout and cancellations can be handled using`context` package in go. This repository supports the 
examples discussed in this blog post.

## Creating new context

- context.Background() : 
- context.Todo() : 

## Deriving context from an existing context

- context.WithCancel(parentCtx)
- context.WithDeadline(parentCtx)
- context.WithTimeout(parentCtx)

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
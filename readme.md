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
    cancel() // call the cancel to cancel the context
}
```

2. cancelling after few seconds
```go
func main() {
    _,cancel := context.WithTimeout(parentCtx, time.Second * 10) // cancel the context in 10 seconds
    defer cancel() // cancel explicitly (when function returns) 
}
```

3. cancelling at a given time
```go
func main() {
    _,cancel := context.WithDeadline(parentCtx, time.Now().Add(2 * time.Second)) // cancel the context at given time
    defer cancel() // cancel explicitly (when function returns) 
}
```
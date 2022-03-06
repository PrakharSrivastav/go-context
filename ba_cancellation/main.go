package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

func init() {
	go func() {
		for {
			log.Println("NumGoroutines :", runtime.NumGoroutine())
			time.Sleep(time.Millisecond * 100)
		}
	}()
}
func main() {
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	// run the work in the background
	go func() {
		if err := work(ctx, "./example.txt"); err != nil {
			log.Println(err)
		}
	}()

	// perform some operation and that causes error
	time.Sleep(time.Millisecond * 110)
	if true {
		stop()
	}

	time.Sleep(time.Second)
}

func work(ctx context.Context, filename string) error {

	out := make(chan string)
	errChan := make(chan error)
	done := make(chan struct{})

	go do(filename, out, errChan, done)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case v := <-out:
			log.Println(v)
		case err := <-errChan:
			return err
		case <-done:
			return nil
		}
	}
}

func do(filename string, out chan string, errChan chan error, done chan struct{}) {
	defer close(out)
	defer close(errChan)
	defer close(done)

	file, err := os.Open(filename)
	if err != nil {
		errChan <- err
		return
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			errChan <- err
			return
		}
		time.Sleep(time.Millisecond * 30)
		out <- strings.TrimSpace(line)
	}
	done <- struct{}{}
	return
}

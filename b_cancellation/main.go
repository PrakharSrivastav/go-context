package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"time"
)

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
	time.Sleep(time.Millisecond * 150)
	if true {
		stop()
	}

	time.Sleep(time.Second)
}

func work(ctx context.Context, filename string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		time.Sleep(time.Millisecond * 100)
		log.Print(line) // do something with the line
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
	return nil
}

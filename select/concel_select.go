package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	// contextを生成
	ctx := context.Background()

	// 親のcontextを生成し、parentに渡す
	ctxParent, cancel := context.WithCancel(ctx)
	go func() {
		err := child(ctxParent)
		if err != nil {
			fmt.Println("*parent err*")
			cancel()
			return
		}
	}()

	// 無限ループ
	for {
		select {
		case <-ctxParent.Done():
			fmt.Println(ctxParent.Err(), "*parent done*")
			return
		default:
			fmt.Println("parent process")
		}
		time.Sleep(1 * time.Second)
	}
}

func child(ctx context.Context) error {
	// parentからcontextを生成し、childに渡す
	childCtx, cancel := context.WithCancel(ctx)
	go func() {
		if err := getErr(); err != nil {
			fmt.Println("*child err*")
			// childCtxをcancel
			cancel()
		}
	}()

	// 無限ループ
	for {
		select {
		case <-childCtx.Done():
			fmt.Println(childCtx.Err(), "*child done*")
			time.Sleep(4 * time.Second)
			return errors.New("")
		default:
			fmt.Println("child process")
		}
		time.Sleep(1 * time.Second)
	}
}

func getErr() error {
	time.Sleep(2 * time.Second)
	return errors.New("")
}

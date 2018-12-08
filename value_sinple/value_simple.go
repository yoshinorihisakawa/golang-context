package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctxValue1 := context.WithValue(ctx, "hoge", 1)
	ctxValue2 := context.WithValue(ctxValue1, "piyo", 2)
	ctxValue3 := context.WithValue(ctxValue2, "fuga", 3)
	ctxValue4 := context.WithValue(ctxValue3, "fuga", 4) // fugaを上書き
	go sayValue(ctxValue4)

	time.Sleep(2 * time.Second)
}

func sayValue(ctx context.Context) {
	for {
		fmt.Print("hoge", ctx.Value("hoge"), " : ")
		fmt.Print("piyo", ctx.Value("piyo"), " : ")
		fmt.Println("fuga", ctx.Value("fuga"))
		time.Sleep(1 * time.Second)
	}
}

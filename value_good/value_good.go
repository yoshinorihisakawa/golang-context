package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	//ctx1 := SetAuthToken(ctx, "auth-token-value")
	authToken, ok := GetAuthToken(ctx)
	if !ok {
		fmt.Println("failed")
		return
	}
	fmt.Println("success : ", authToken)
}

// keyの型情報を定義
type contextKey string

// ここにkeyを設定していく
var (
	keyAuthToken contextKey = "auth-token"
	// keyUserID    contextKey = "user-id"
)

// setするためのfunction
func SetAuthToken(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, keyAuthToken, value)
}

// getするためのfunction
func GetAuthToken(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(keyAuthToken).(string)
	return val, ok
}

package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

func printUserInfo(ctx context.Context) {
	fmt.Println("traceId: ", ctx.Value("traceId"))
}

func test1() {

	newUUID, _ := uuid.NewUUID()

	randId := strings.ReplaceAll(newUUID.String(), "-", "")
	ctx := context.WithValue(context.Background(), "traceId", randId)

	go printUserInfo(ctx)
	time.Sleep(time.Second)
}

func test2() {

	newUUID, _ := uuid.NewUUID()

	randId := strings.ReplaceAll(newUUID.String(), "-", "")
	ctx := context.WithValue(context.Background(), "traceId", randId)

	userCtx := context.WithValue(ctx, "traceId", "dct")
	go printUserInfo(userCtx)
	time.Sleep(time.Second)

}

func main() {

	//test1()
	test2()
}

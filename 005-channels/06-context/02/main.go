package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:", ctx)
	fmt.Println("context err:", ctx.Err())
	fmt.Printf("context type:%T\n", ctx)
	fmt.Println("-----------")

	ctx, _ = context.WithCancel(ctx)

	fmt.Println("context:", ctx)
	fmt.Println("context err:", ctx.Err())
	fmt.Printf("context type:%T\n", ctx)
	fmt.Println("-----------")

}

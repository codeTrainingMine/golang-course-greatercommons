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

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:", ctx)
	fmt.Println("context err:", ctx.Err())
	fmt.Printf("context type:%T\n", ctx)
	fmt.Println("cancel:", cancel)
	//fmt.Println("cancel err:", cancel.Err())
	fmt.Printf("cancel type:%T\n", cancel)
	fmt.Println("-----------")

	cancel()

	fmt.Println("context:", ctx)
	fmt.Println("context err:", ctx.Err())
	fmt.Printf("context type:%T\n", ctx)
	fmt.Println("cancel:", cancel)
	//fmt.Println("cancel err:", cancel.Err())
	fmt.Printf("cancel type:%T\n", cancel)
	fmt.Println("-----------")
}

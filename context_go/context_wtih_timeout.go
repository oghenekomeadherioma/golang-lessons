package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 0)
	// defer cancel()

	// go func() {
	// 	select {
	// 	case <- ctx.Done(): fmt.Println("Context is Done/Cancelled")
	// 	}
	// }()

	ctx1 := context.Background()
	ctx_wt, cancel := context.WithTimeout(ctx1, time.Millisecond * 300)
	defer cancel()
	fmt.Println(context.Background().Done())
	fmt.Println(ctx_wt.Done())
}
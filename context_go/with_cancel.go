package main

import (
	"fmt"
	"context"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel

	go func() {
		select {
		case <- ctx.Done: fmt.Println("Context is Done/Cancelled")
		}

	}()
}
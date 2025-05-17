package main 

import (
	"fmt"
	"context"
)

func main() {
	ctx := context.Background()
	ctx_wd, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		select {
		case <- ctx_wd.Done():
			fmt.Println("Context is Completed")
			cancel()
		}
	}()
}
package main 

import (
	"fmt"
	"context"
)

func main() {
	ctx := context.Background()
	next_ctx := context.WithValue(ctx, "key", "Value")
	fmt.Println(next_ctx.Value("Key"))
}
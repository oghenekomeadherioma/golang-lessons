package main 
import (
	"context"
	"fmt"
)

func main() {
	ctx2 := context.Background()
	ctx := context.TODO()
	ctx3 := context.WithValue(ctx2, "Name", "Mustapha")
	dosomething(ctx)
}

func dosomething(ctx context.Context){
	fmt.Println("Do something", ctx.Value())

}
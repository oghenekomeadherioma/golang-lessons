package main

import (
	"fmt"
)

func main() {
	series(1, 2, 3, 4, 5)
}

func series(num ...int){
	for  y := range num{
		fmt.Println(y)
	}
}


package main
import (
	"fmt"
	"time"
)

func main() {
	
	go sayHello("World")
	sayHello("Hello")
}

func sayHello(message string){
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(time.Second)
	}
}
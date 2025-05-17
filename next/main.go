package main

import (
	"fmt"
	"unicode/utf8"
	"os"
)

func main() {
	var input string
	fmt.Println("What you wanna count")
	fmt.Scanln(&input)
	fmt.Println(utf8.RuneCountInString(input))
	len := utf8.RuneCountInString(input)
	fmt.Printf(" %T \n", len)
	// As argument
	val1 := os.Args[1]
	//val_len := len(val)
	fmt.Printf(" %T  %v \n", val1)


}
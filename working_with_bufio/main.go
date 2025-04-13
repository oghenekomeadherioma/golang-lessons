package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	myfile, err := os.Open("test_file.txt")
	if err != nil {
		panic(err)
	}
	content := bufio.NewReader(myfile)

	fmt.Println(content.ReadString('\n'))
}
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	myfile, err := os.Create("new_file.txt")
	if err != nil {
		panic(err)
	}
	defer myfile.Close()
	the_file := bufio.NewWriter(myfile)
	//file_contet := strings.Reader(the_file)
	fmt.Println(the_file)
}
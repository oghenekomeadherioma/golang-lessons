package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	file := flag.String("file", "test.txt", "Name of File")
	content := flag.String("content", "This is a test data\n", "Content to write to file")
	flag.Parse()
	// filename := "test.txt"
	// data := "This is a test data\n"
	//createFile(filename)
	writeFile(*file, *content)
	fmt.Println("File created and data written successfully")

}

func createFile(filename string) {
	file, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Println("Creating file file  ", file)
		os.Create(filename)
	}

}

func writeFile(filename string, data string) {
	file_info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Println("File does not exist", file_info, "The follwoing error was observed:", err)
		createFile(filename)
	}

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(file, data)
	if _, err := file.WriteString(data); err != nil {
		log.Fatal(err)
	}
}

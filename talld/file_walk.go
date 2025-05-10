package main
import (
	"os"
	"log"
	"fmt"
)

func main() {}
func createFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("File created successfully")
}

func writeToFile(fileName string, data string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data written to file successfully")
}
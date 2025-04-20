package main
import (
	"time"
	"fmt"
)

func main() {
	dob := time.Date(1940, time.October, 18, 9, 27, 6, 7, time.UTC)
	fmt.Println(dob)
}
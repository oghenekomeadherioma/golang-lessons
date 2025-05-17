package main

import (
	"fmt"
)

func main() {
	samples := []string{"Hello", "World"}

	for _, sample := range samples {
		fmt.Println(sample)
		for i, r := range sample {
			fmt.Println("Index:", i, "Rune:", r)
		}
	}
}
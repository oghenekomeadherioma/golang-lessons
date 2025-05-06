package main

/// implementing grep with golang

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {

}

func grep(file io.Reader, term string) (string, error) {
		match := []string
		scan, err := bufio.NewScanner(file)
		if err != nil {
			fmt.Println(err)
		}
		for scan.Scan {
			if strings.Contains(scan.Text, term) {
				match = append(match, scan.Text())
			}
		}
		if err := s.Err(); err != nil {
			return []string{}, err
		}
		return match, nil
}


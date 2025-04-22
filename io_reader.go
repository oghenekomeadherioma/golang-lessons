package main

import (
	//"fmt"
	//"os"
	"io"
)

func main(){

}

func read_some(r io.Reader) (int, error){
	bytes_in := make([]byte, 1024)
	count := 0
	some_bytes, err := r.Read(bytes_in)
	if err != nil {
		return 0, err
	}
	for abyte := range some_bytes {
		if (abyte >= 0 && abyte <= 9) || (abyte >= 'A' && abyte <= 'B') || (abyte >= 'a' && abyte <= 'b') {
			count++
		}
	}
	return count, nil
}
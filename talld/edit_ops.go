package main 

import (
	"fmt"
)

func main() {

}

type Instance struct {
	original_list []string
	deleted_list  []string
	next_add string
	last_add string
}

func add(alist *[]string, item string) *[]string{
	add_pos := append(*alist, item)
	fmt.Println("Added item at position", len(*alist), ":", add_pos)
	return &add_pos
}

// delete removes the last element from the slice `alist` and appends it to a 
// separate slice `deleted`. It then prints the deleted item and its position.
// Note: The function as written has issues, such as the use of an undefined 
// `alist` variable and potential scoping problems. Ensure `alist` is properly 
// defined and accessible before using this function.
func delete(){
	last := alist[len(alist)-1]
	deleted := []string{}
	alist := alist[0:len(alist)-1]
	deleted = append(deleted, last)
	fmt.Println("Deleted item at position", len(alist), ":", deleted)

}
func redo(alist []string){
	deleted := []string{}
	alist = append(alist, deleted[len(deleted)-1])
	deleted = deleted[0:len(deleted)-1]
}
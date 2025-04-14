package main 

import (
	"fmt"
	"os"
	"encoding/json"
	"io"
)

func main() {
	unilag := school{City : "Lagos", Name : "University of Lagos", Short : "Unilag"}
	segun := student{ school : unilag, Name: "Olusegun", Course : "Business Admin", Year : "2016", CGPA : 3.45}

	segun_details, err := json.Marshal(segun)
	if err != nil {
		panic(err)
	}

	segun_file, err := os.Create("segun_details.json")
	if err != nil {
		fmt.Print("Err: ", err)
		return
	}
	_, err = io.WriteString(segun_file, string(segun_details))
	if err != nil {
		panic(err)
	}

}

type school struct {
	City string
	Name string
	Short string

}

type student struct {
	school
	Name string 
	Course string 
	Year string //time.Time
	CGPA float64
}
package main 

import (
	"fmt"
	"os"
	"encoding/json"
	"time"
)

func main() {
	unilag := school{City : "Lagos", Name : "University of Lagos", Short : "Unilag"}
	segun := student{ school : unilag, Name: "Olusegun", Course : "Business Admin", Year : "2016", CGPA : 3.45}

	segun_details, err := json.Marshal(segun)
	if err != nil {
		panic(err)
	}

	segun_file := os.WriteFile(segun_details)
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
package main
import (
	"time"
	"fmt"
)

func main() {
	dob := time.Date(1940, time.October, 18, 9, 27, 6, 7, time.UTC)
	fmt.Println(dob)
	fmt.Println(dob.Format("2025-05-05"))
	fmt.Println(dob.Format("2006-01-02"))
	fmt.Println(dob.Format("Mon, Jan, 02"))
	 d := 35000 * time.Millisecond
	fmt.Println(d)
	tz, _ := time.LoadLocation("America/New_York")
	fmt.Println(tz)

	// Conver t one time zone to another
	lag_time := time.Date(2025, time.April, 21, 12, 0, 0, 0, time.UTC)
	fmt.Println(lag_time)

	// 
}
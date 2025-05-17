package main

import (
	//"bytes"
	"regexp"
	"fmt"
)

// regexp.MatchString()
// regexp.MustCompile()
// regexp.Compile()
// regexp.FindString()
// regexp.FindStringIndex()
// regexp.FindAllString()

func main() {
	// Example string to search
	text := "The quick brown fox jumps over the lazy dog. 1234567890"

	// Regular expression pattern to match digits
	pattern := `\d+`

	// Compile the regular expression
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Find all matches in the text
	matches := re.FindAllString(text, -1)

	// Print the matches
	fmt.Println("Matches found:", matches)

	// Check if a string matches the pattern
	isMatch := re.MatchString("12345")
	fmt.Println("Does '12345' match the pattern?", isMatch)

	inst, _ := regexp.MatchString("p([a-z]+)ch", "preach")
	fmt.Println(inst)
	inst2 := regexp.MustCompile("([a-z]+)@([a-z]+).([a-z]+)")
	email_1 := inst2.MatchString("tom@inz.in")
	lots_email := `
	The annual Spring Harvest Festival is just around the corner! Visitors can explore over 50 local
	 vendors offering fresh produce, handmade crafts, and delicious treats. Don’t miss the live music
	  performances starting at noon. For vendor inquiries, reach out
	   to maria.fernandez@greenvalleyevents.net before April 20th. Additionally, gardening workshops
	    will be held hourly near the main stage. Bring reusable bags to receive a complimentary 
		seedling. Parking is free, but carpooling is encouraged due to limited space.
	 Let’s celebrate sustainability together!
	`
	email_2 := `The annual Spring Harvest Festival is just around the corner! 
	Visitors can explore over 50 local vendors offering fresh produce, handmade crafts, 
	and delicious treats. For volunteer sign-ups, email volunteers@springharvestfest.com by April 15th.
	 Don’t miss live music performances starting at noon. Vendor inquiries? 
	 Contact maria.fernandez@greenvalleyevents.net before April 20th. Gardening workshops 
	 
	 (hourly near the main stage) require registration—email workshops.gvf@outlook.com to 
	 secure your spot. Bring reusable bags for a complimentary seedling. Parking is free, 
	 but carpooling is encouraged due to limited space. Let’s celebrate sustainability together!
	 Rain date updates will be posted on our website. See you there! `
	fmt.Println(inst2.FindString(lots_email))
	fmt.Println(inst2.FindString(email_2))
	fmt.Println(email_1)
	f_regex := regexp.MustCompile("([a-z]+)i([a-z]+)")
	fmt.Println(f_regex.ReplaceAllString("foolish", "foolishness"))
}
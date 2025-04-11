package main

import "fmt"

func main() {
	messageStart:= "Happy birthday! you are now"
	age:= 21
	messageEnd:= "years old!"

	fmt.Println(messageStart, age, messageEnd)
}
// idea was to join the different formats together using the fmt operation and the walrus operator to produce a readable output 
// walrus operator can only be used in a function but the var can be used out of a function

package main

import "fmt"

func main() {
	fmt.Println("Starting Textio server...")
}

// this function involved the main function for printing an output unto the screen for the basics of the go language and syntax

package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
    var hasPermission bool 
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

// bringing out the variables listing out the variables and their coresponding functions

//fixing a bug for the codes 

/* Multiline comments come in this format
so now you can have multiple lines of code i fixed the bug in this sample */
package main

import "fmt"

func main() {
	/*
		We are increasing the maximum message length from 140 to 280 characters.
		Very reluctantly, I might add.
		Users actually want to write more than 140 characters?!? Madness.
	*/
	maxMessageLength := 140
	newMaxMessageLength := 280
	fmt.Println("Textio is increasing the maximum message length from", maxMessageLength, "to", newMaxMessageLength, "characters.")
}

// 11th April, 2025
// users of compiled languages you dont need access to the source code 
// fibonacci series with golang 
package main 

import "fmt"

func main(){
	var n int 
	fmt.Print("Enter the number of Fibonacci terms: ")
    fmt.Scan(&n)

    a, b := 0, 1

    fmt.Println("Fibonacci Series:")

    for i := 0; i < n; i++ {
        fmt.Print(a, " ")
        a, b = b, a+b
    }
}


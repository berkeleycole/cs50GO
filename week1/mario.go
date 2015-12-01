package main

import (
	"fmt"
	"errors"
	"strings"
)

// this is a comment

var (
	NegativeNumErr = errors.New("You can't select a negative number")
	BigNumErr = errors.New("That's more than 23! That would be too tall of a pyramid for us!")
)

func main() {
	fmt.Println("Please choose the hight of your pyramid! Any positive number no greater than 23 will do:\n")
	// take in single integer from user
	var i int
	fmt.Scan(&i)
	//make sure user input is >=0 and <=23 - if not, send error messages
	for i < 0 {
		fmt.Println(NegativeNumErr)
		fmt.Println("Try Again:")
		fmt.Scan(&i)
	}
	for i > 23{
		fmt.Println(BigNumErr)
		fmt.Println("Try Again:")
		fmt.Scan(&i)
	}
	fmt.Println("You are building a pyramid", i, "blocks high.")
	//pyramid construction
	for x:=2; x <= i+1; x++ {
		fmt.Println(strings.Repeat(" ", (i+1)-x) + strings.Repeat("#", x))
	}
}
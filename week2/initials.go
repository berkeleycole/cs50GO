package main

import (
	"bufio"
	"fmt"
	"github.com/berkeleycole/initials"
	"os"
)

func main() {
	//get full name
	fullName := bufio.NewReader(os.Stdin)
	fmt.Print("Your full name, please: ")
	text, err := fullName.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	initials, err := initials.GetInitials(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Your Initials:", initials)
}

//still need to: 1) give error for numbers in name (not actually required per the assignment)
//				 2) collect all initials in an array which is printed at the end, instead of printing each one individually as the program finds it

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	HasNumErr = errors.New("You cannot have numbers in your name.")
)

func cleanString(a string) bool {
	return strings.ContainsRune(a)
}

func main() {

	//get full name
	fullName := bufio.NewReader(os.Stdin)
	fmt.Print("Your full name, please: ")
	text, _ := fullName.ReadString('\n')

	if cleanString(text) == true {
		fmt.Println(HasNumErr)
	}

	//name to all caps
	var name string = strings.ToUpper(text)

	initials := make([]string, len(name))

	//print first initial
	fmt.Printf("%c", name[0])

	//loop through all chars
	i := 0

	for i < len(name) {
		if name[i] == 32 { //find spaces
			i++
			fmt.Printf("%c", name[i]) //take the char after the space as a new initial
		} else {
			i++ //otherwise, skip past
		}
	}

	fmt.Printf("\n") //extra space for readability
}

//still need to: 1) give error for numbers in name (not actually required per the assignment)
//				 2) collect all initials in an array which is printed at the end, instead of printing each one individually as the program finds it

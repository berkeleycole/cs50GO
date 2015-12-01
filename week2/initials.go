package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var HasNumErr = errors.New("You cannot have numbers in your name.")

func main() {
	//get full name
	fullName := bufio.NewReader(os.Stdin)
	fmt.Print("Your full name, please: ")
	text, err := fullName.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	initials, err := GetInitials(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Your Initials:", initials)
}

func GetInitials(name string) (string, error) {
	if strings.IndexAny(name, "1234567890") >= 0 {
		return "", HasNumErr
	}

	//name to all caps
	caps := strings.ToUpper(name)

	var initials []string

	initials = append(initials, string(caps[0]))

	//loop through all chars
	for i := range name {
		if caps[i] == 32 {
			initials = append(initials, string(caps[i+1]))
		}
	}

	ret := strings.Join(initials, "")

	return ret, nil
}

//still need to: 1) give error for numbers in name (not actually required per the assignment)
//				 2) collect all initials in an array which is printed at the end, instead of printing each one individually as the program finds it

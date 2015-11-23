package main

import (
	"fmt"
	"errors"
	"bufio"
	"os"
)

var (
	Err = errors.New("")
)

func main() {
	//get full name
	fullName := bufio.NewReader(os.Stdin)
	fmt.Print("Your full name, please: ")
	text, _ := fullName.ReadString('\n')
	fmt.Println(text[0])

}
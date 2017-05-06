package main

import (
	"fmt"
)

func main() {
	//ask user for input
	fmt.Println("How long is your typical shower?")

	var minutes int
	fmt.Scan(&minutes)

	fmt.Println("That's ", minutes*192/16, " bottles of water!")
}

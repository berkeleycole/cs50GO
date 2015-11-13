package main

import (
	"fmt"
	"errors"
)

var (
	NegativeNumErr = errors.New("You can't select a negative number")
)

func main() {
	//ask user for a float(amount due in change)
	fmt.Println("How much in change? \n")
	var change float32
	fmt.Scan(&change)
	
	//Continue asking for a number until a positive number is given
	for change < 0 {
		fmt.Println(NegativeNumErr)
		fmt.Println("Try Again:")
		fmt.Scan(&change)
	}
	//multiply the float by 100 and truncate(setting $ amounts to cents)
	change = change*100
	fmt.Printf("That is %.0f cents!\n", change)
	
	var cents int = int(change+.05)
	fmt.Println("That is ", cents)
	//set coin counter variable
	coins := 0
	//loop breaks the input into 25's, 10's, 5's and 1's and count coints
	for cents >= 25 {
		coins++
		cents -= 25
	}
	for cents < 25 && cents >= 10 {
		coins++
		cents -= 10
	}
	for cents < 10 && cents >= 5 {
		coins++
		cents -= 5
	}
	for cents < 5 && cents >= 1 {
		coins++
		cents -= 1
	}
	//Print number of coins used 
	fmt.Println("Coins", coins)
}
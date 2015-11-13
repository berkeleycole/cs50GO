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
	//loop will break the input into 25's, 10's, 5's and 1's
		//if i > 25
			//subtract total by 25 and increase coin count by one
		//if i < 25 && i > 10
			//subtract total by 10 and increase coin count by one
		//if i < 10 && i > 5
			//subtract total by 5 and increase coin count by one
		//if i < 5 && i > 1
			//subtract total by 1 and increase coin count by one
			
	//Print number of coins used to dispense the designated amount of change
}
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	//"reflect"
	"strconv"
)

var (
	NotNumErr = errors.New("You must provide a number which is greater than 0")
)

func main() {

	//GET USER INPUT FROM CLI (this is os.Args[1])

	Arg, _ := strconv.Atoi(os.Args[1]) // convert from string to int
	// converting types in Go is not as hard as it first seems. desiredType(variable to change)
	//In this case, key must be a byte in order to later be added to letter
	key := byte(Arg)

	//CLEAN KEY VALUE

	// To see the series of changes for key, uncomment the extra print lines.
	//fmt.Println(key, "original key")
	if key > 0 { //if key is a valid number
		if key >= 26 {
			//numbers greater than 26 are not helpful in this program and must be changed. ROTATION 26 == ROTATION 0, the message will not change. An extra clause could be added here for key == 26 which would skip straigt to printing the message given by the user, or warn of a useless encryption

			for key >= 26 { //continues to decrement a large key by 26 until a valid value is reached
				//fmt.Println(key, "decrementing key value")
				key = key - 26
			}
		}
		// fmt.Println(key, "cleaned key")

		//GET MESSAGE

		text := bufio.NewReader(os.Stdin)   //capeable of taking in multiple words
		message, _ := text.ReadString('\n') //convert to string

		//ENCRYPTION

		//first declare slice to hold the encrypted message
		var EncMessage = make([]byte, len(message))

		//encrypt message, character by character
		for i := 0; i < len(message); i++ {

			var letter byte = message[i] //declare letter

			if letter >= 65 && letter <= 91 { //if uppercase letter
				letter = letter + key
				if letter > 91 {
					letter = letter - 26
				}
			} else if letter >= 97 && letter <= 122 { //if lowercase letter
				letter = letter + key
				if letter > 122 {
					letter = letter - 26
				}
			} // else (letter is punctuation){ do nothing }
			EncMessage = append(EncMessage, letter)
		}

		//To see the the slice of bytes in the encrypted message, uncomment:
		//fmt.Println(EncMessage, "Encrypted Message Slice")
		//To understand the conversion better, play with reflect.TypeOf. Uncomment "reflect" in imports and then uncomment:
		//fmt.Println(reflect.TypeOf(EncMessage))
		NewMessage := string(EncMessage) //converts slice to string
		fmt.Println(NewMessage)

	} else { // user input was 0 or a letter - returns error and exits program
		fmt.Println(NotNumErr)
	}
}

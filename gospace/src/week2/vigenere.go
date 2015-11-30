package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	WrongArgsErr = errors.New("You must provide one key word")
)

func keySetter(s string, i int) byte {

	s = strings.ToUpper(s) //easier if key is all in same case
	key := byte(s[i])
	return key - 65
}

func main() {

	// 1. GET USER INPUT FROM CLI (this is a string in os.Args[1])

	keyWord := os.Args[1]

	// 2. SET KEY INDEX

	keyIndex := 0 //declared outside the if statement so that incrementing from any sub-statement effects the overall count

	// 3. GET MESSAGE

	fmt.Println("Enter your message:")
	text := bufio.NewReader(os.Stdin)     //capeable of taking in multiple words
	message, err := text.ReadString('\n') //convert to string
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 4. ENCRYPT

	//set the encrypted message slice outside the for loop
	var EncMessage = make([]byte, len(message))

	for i := 0; i < len(message); i++ { //find the message letter

		letter := byte(message[i])
		var key byte

		if letter >= 65 && letter <= 91 { //if uppercase letter
			key = keySetter(keyWord, keyIndex)
			letter = letter + key
			if letter > 91 {
				letter = letter - 26
			}
			keyIndex++
		} else if letter >= 97 && letter <= 122 { //if lowercase letter
			key = keySetter(keyWord, keyIndex)
			letter = letter + key
			if letter > 122 {
				letter = letter - 26
			}
			keyIndex++
		} // else (letter is punctuation){ do nothing }

		if keyIndex >= len(keyWord) {
			keyIndex = 0
		}

		EncMessage = append(EncMessage, letter)
	}

	NewMessage := string(EncMessage) //converts slice to string
	fmt.Println(NewMessage)
}

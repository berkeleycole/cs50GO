##Assignment 4: Caesar Cypher

To be honest, the amount of jumping between data types in this assignment had me just about pulling my hair out...and more so in Go than in C. But, once you get the hang of it, it really isn't completely aweful, just annoying. 

The goal of the assignment was to make a simple caesar cypher that took in a key value and a message and rotated the letters of each message by the key to make a "secret" message. The course lectures go into great detail about why the Caesar cypher is really not a secure encryption method, but it was difficult enough to create and as an assigment I found it fairly time consuming. 

The most important thing about Go that I learned during this assignment is that there are many data types with subtle differences that really do matter because they don't play nice together until you convert them. It wasn't until I had gone around in circles for awhile that I came across the simple conversion syntax that generally works:
	thingYouWant = DesiredType(variable to convert)
	
The other thing I found helpful was the reflect library that allows you to run reflect.TypeOf and see what types you are actually dealing with. I was printing TypeOf lines before and after most conversions in this program while I was working on it. 

This is the C code for comparison: 

	include <stdio.h>
	include <cs50.h>
	include <stdlib.h>
	include <string.h>
	include <ctype.h>

	int main(int argc, string argv[]){
	    if(argc == 2){ //we have the right CLI args, now convert the string argv to int to define our key
	        int key = atoi(argv[1]);
	        //normalize large key values (values greater than 26 are useless)
	        while(key > 26){
	            key -= 26;
	        }
	        //if the string can't be converted, or is 0, exit program
	        if(key == 0){
	            printf("Please be sure that your CLI input is an integer.\n");
	            return 1;
	        } 
	        //ok, we have a key, now ask for a string from the user to encrypt
	        //printf("Message to be encrypted:\n");---CS50 checker doesn't like this, but I think it adds clarity
	        string message = GetString();
        
	        //start encrypting the message, character by character
	        for(int i = 0; i < strlen(message); i++){
	            char letter = message[i];
	            if(isalpha(letter)){ //we only want to change chars, not other symbols
	                int newLetter;
	                if(isupper(letter)){ //so that chars stay upper or lower case
	                    newLetter = letter + key;
	                    if(newLetter > 90){ //makes the ascii char value run in a circle (Z + 1) = A
	                        newLetter -= 26; 
	                    }
	                }if(islower(letter)){ //repeat process for lower case chars
	                    newLetter = letter + key;
	                    if(newLetter > 122){
	                        newLetter -= 26;
	                    }
	                }
	                printf("%c", newLetter); //print each encrypted char
	            }else{
	                // if the char is not in the alphabet, it is punctuation and we want to print it unchanged
	                printf("%c", letter); 
	            }
	        }
	        printf("\n"); //end space is needed for the cs50 checker
	    } else {
	        //for incorrect/missing argv inputs
	        printf("incorrect command line input.\n");
	        return 1; 
	    }
	}
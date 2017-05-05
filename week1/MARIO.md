##Assignment 1: Mario Tower

My first actual program in Go! Found out just how wonderful all the standard library packages are, this program was smoother to write in Go than in C for sure. 

The assignment was to create a half-pyramid of sharps, aligning to the RIGHT, forcing the programmer to do some tricky stuff with adding a cascading number of spaces to the left. The user designates the height of pyramid to be built (anywhere between 0 and 23 inclusive) and the program takes in the integer and constructs a half pyramid with that height. 

I also learned some good lessons about error messaging with the error package, which (as I understand at this point) is the preffered method among Go programmers. 

For reference, here is my C code for the same program: 

include <stdio.h>
include <cs50.h>

int main(void)
{
    //ask user for an integer
    
    printf("Please choose the hight of your pyramid! Any positive number no greater than 23 will do:\n");
    int n = GetInt();
    
    while(n < 0 || n > 23){
        printf("Please choose a valid number - one that is greater than 0 and no greater than 23.\n"); 
        n = GetInt();
      }
    
    if(n > 0 && n < 24)
    {
       for(int i = 2; i <= n+1; i++){ 
            printf("%.*s", (n+1)-i, "                       ");
            printf("%.*s\n", i, "########################");
       }
    }
    else 
    {
        printf("No pyramid could be built with your input. Try again with a different number.\n");
    }
    
 }

##Assignment 2: Greedy Coin Counter

In this assignment, you must build a program that dispenses change. The goal is to dispense the right amount of change using the fewest coins possible. It requires doing some tricky conversions from floats to ints, and a good understanding of loops and counters. 

The hardest part of this assignment was the conversion of floats to ints. Making sure that the float value of $1.05 didn't become 104 cents meant that I had to automatically round up when making the transition to an int. I had never understood before how difficult it was for computers to deal with decimal values. 

For reference, here was my program in C. (It didn't get full credit in the CS50 course, and I never could quite figure out why, since I seemed to get back the correct number of coins - but I think my Go program is actually better!)

int main(void)
{
    //ask user for a float
    
    printf("How much in change?\n");
    float n = GetFloat();
    
    if(n < 0 ){
        printf("Retry:");
        n = GetFloat();
    }
    
    int cents = roundf(n*100);
    int coins = 0;
    if(cents > 0){
        for(; cents >=25; coins++){
            cents = cents-25;
        } 
        for(;cents < 25 && cents >=10; coins++){
            cents = cents-10;
        }
        for(;cents < 10 && cents >=5; coins++){
            cents = cents-5;
        }
        for(;cents < 5 && cents >=1; coins++){
            cents = cents-1;
        }
    }
    else {
        printf("Make sure you entered a valid number!");
    }
    
    printf("%d", coins);
} 
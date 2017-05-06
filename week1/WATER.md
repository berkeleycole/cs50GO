##Assignment 3: Greedy Coin Counter

#This assignment is new as of this year. I'm actually not sure if this is great Go syntax (primarily looking at the final message "That's (minutes*192/16) bottles of water!") but it does work. Below is my C code for reference.

int main(void)

{

    printf("How long is your usual shower?\n");

    int minutes = get_int() * 192 / 16;

    printf("Thats %d water bottles of water!", minutes);

}

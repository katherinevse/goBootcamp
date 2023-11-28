#include <stdio.h>
#include <string.h>



int lengthOfLastWord(char* s) {
    int i = strlen(s) - 1;
    int length = 0;
    while (i >= 0 && s[i] == ' '){
        i = i-1;
    }
    while (i >= 0 && s[i] != ' ')
    {
        length++;
        i=i-1;
    }
    
    return length;
}

int main() {
    char str[] = "Hello World";
    int result = lengthOfLastWord(str);
    printf("Length of the last word: %d\n", result);

    return 0;
}
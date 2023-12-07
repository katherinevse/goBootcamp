#include <stdio.h>
#include <string.h>
#include<stdlib.h>

char* longestCommonPrefix(char** strs, int strsSize) {
    char **ptrToStrs = strs;

    for (int i = 0; i < strsSize; i++)
    {
        if(strncmp(ptrToStrs[i], ptrToStrs[i+1], 1) != 0)
            printf("Первые символы в словах разные.\n");
            return NULL;

    }
    printf("Первые символы в каждом слове одинаковы.\n");
    char *result = malloc(n);
    result[0] = ptrToStrs[0][0];
    result[1] = '\0';

    return result;

}

int main(){
    char *strs[] = {"flower", "flow"};
    int n = 2;
    char *commonPrefix = longestCommonPrefix(strs, n);

    if (commonPrefix != NULL) {
        printf("Общий префикс: %s\n", commonPrefix);
        free(commonPrefix); 
    }

    return 0;

}


// Input: strs = ["flower","flow","flight"]
// Output: "fl"

#include <stdio.h>
#include <stdlib.h>

bool isPalindrome(int x){
    int result =0,sum=0,num=x;
    while(x>0){
        result = x % 10;  // result = 1
        sum = sum * 10 + result;   // sum = 1
        x /= 10;            // x = 12
    }

    if(sum==num)
    return true;
    else
    return false;
}
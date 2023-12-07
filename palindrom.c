#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

bool isPalindrome(int x){
int t=x,b,s; //t = 121
while(x>0){
    b=x%10; // 1
    s=s*10+b;
    x/=10;
}
if(s==t)
  return true;
else 
  return false;
}

int main(){
  int val =  121;
  isPalindrome(val);

  return 0;

}
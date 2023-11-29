
#include <stdio.h>

int findMaxConsecutiveOnes(int* nums, int numsSize) {
  int lenght = 0;
  int newLenght = 0;
  for (int i = 0; i < numsSize; i++) {
    if (nums[i] == 1) {
      newLenght++;
    }
    if (nums[i] != 1) {
      newLenght = 0;
    }
    if (newLenght > lenght) lenght = newLenght;
  }
  return lenght;
}

int main() {
  int nums[] = {1, 1, 0, 1, 1, 1};
  int nsize = 6;
  printf("%d", findMaxConsecutiveOnes(nums, nsize));

  return 0;
}
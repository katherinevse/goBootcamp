#include <stdio.h>
#include <stdlib.h>

int removeElement(int* nums, int numsSize, int val) {
  int left = 0;
  int right = 0;

  while (right < numsSize) {
    if (nums[right] == val) {
      right++;
    } else {
      nums[left] = nums[right];
      right++;
      left++;
    }
  }

  return left;
}

int main() {
  int nums[] = {3, 2, 2, 3};
  int val = 3;
  int nSize = 4;

  removeElement(nums, nSize, val);
  return 0;
}
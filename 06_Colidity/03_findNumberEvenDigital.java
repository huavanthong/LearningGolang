
/*
Reference:
    https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3237/
*/
package main;

class Solution {
    public int findNumbers(int[] nums) {
        
        int countEven;
        int count =0;
        int remainder;
        for(int i = 0; i < nums.length; i++) {
            countEven = 0;
            int temp = nums[i];
            while(temp > 0) {
                remainder = temp % 10;
                temp = temp/10;
                countEven++;
            }
            
            if(countEven % 2 == 0) {
                count++;
            }
        }
        
        return count;
    }
}
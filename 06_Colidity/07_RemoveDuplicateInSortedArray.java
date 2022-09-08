package main;

public class Solution {
    public int removeDuplicates(int[] nums) {
//         int count = 0;
//         for(int i = 0; i < nums.length - 1; i++) {
//             if(nums[i] == nums[i+1]) {
//                 for(int j = i; j < nums.length - 1; j++) {
//                     nums[j] = nums[j+1]; 
//                 }
//                 count++;
//             }
            
//         }
//         return count;

        if (nums.length == 0) return 0;
        
        int i = 0;
        for (int j = 1; j < nums.length; j++) {
            if (nums[j] != nums[i]) {
                i++;
                nums[i] = nums[j];
            }
        }
        return i + 1;
    }
}
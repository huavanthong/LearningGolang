/*
 * 
    Given a binary array, find the maximum number of consecutive 1s in this array.
 */

class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
        // Hint: Initialise and declare a variable here to 
        // keep track of how many 1's you've seen in a row.
        int max = 0;
        int count = 0;
        for (int i = 0; i < nums.length; i++) {
            // Do something with element nums[i].

            if(nums[i] == 1) {
                count++;
            } else {
                count = 0;
            }

            if (count > max) {
                max = count;
            }
        }

        return max;
    }
}
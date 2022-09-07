/*
Reference:
     https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3240/
*/
class Solution {
    public int[] sortedSquares(int[] nums) {
        
        int[] arr = new int[nums.length];
        
        for(int i =0; i < nums.length; i++) {
            arr[i] = nums[i] * nums[i];
        }
        
        int len = arr.length;
        for(int i = 0; i < len - 1; i++) {
            int min_id = i;
            for(int j = i +1; j < len; j++) {
                
                if(arr[j] < arr[min_id]) {
                    min_id = j;
                }
            }
            
            int temp = arr[min_id];
            arr[min_id] = arr[i];
            arr[i] = temp;
        }
        
        return arr;
    }
}
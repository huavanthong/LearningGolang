package 06_Colidity;

public class Solution {
    public static void main(String[] args) {
	    
	    int[] test = new int[]{0,1,0,3,12};
	    
	    moveZeroes(test);

        for(int i = 0; i < test.length; i++) {
            System.out.println("Value: " + test[i]);
        }
		
	}

    public static void moveZeroes(int[] nums) {
        
        if(nums.length == 0) return;
        
        // Count zeros array
        int count = 0;
        for(int i = 0; i < nums.length; i++) {
            if(nums[i] == 0) {
                count++;
            }
        }
        
        // Check value not equal zeros in array
        // Then assign value to arrays
        int size = nums.length - count;
        for(int i = 0, pos = 0; i < nums.length; i++) {
            if(nums[i] != 0) {
                nums[pos] = nums[i];
                pos++;
            }
        }
        
        // Set zeros to other the rest of array.
        for(int i =size; i < nums.length; i++) {
            nums[size] = 0; 
            size++;
        }
    }
}
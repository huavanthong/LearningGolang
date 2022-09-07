package main;

import java.util.Arrays; 
class Main {   
public static void main(String[] args) 
    {   
        // define original array 
        int[] tensArray = { 10,20,30,40,50,60}; 
 
        // Print the original array 
        System.out.println("Original Array: " + Arrays.toString(tensArray)); 
 
        // the index at which the element in the array is to be removed 
        int rm_index = 2; 
 
        // display index 
        System.out.println("Element to be removed at index: " + rm_index); 
 
        // if array is empty or index is out of bounds, removal is not possible 
          if (tensArray == null
            || rm_index&lt; 0
            || rm_index&gt;= tensArray.length) { 
 
            System.out.println("No removal operation can be performed!!");
        } 
        // Create a proxy array of size one less than original array
        int[] proxyArray = new int[tensArray.length - 1]; 
 
        // copy all the elements in the original to proxy array except the one at index 
        for (int i = 0, k = 0; i &lt;tensArray.length; i++) { 
 
            // check if index is crossed, continue without copying 
            if (i == rm_index) { 
                continue; 
            } 
 
            // else copy the element
            proxyArray[k++] = tensArray[i]; 
        } 
 
       // Print the copied proxy array 
       System.out.println("Array after removal operation: " + Arrays.toString(proxyArray)); 
    } 
}

public class FinalSolution
{
	public static void main(String[] args) {
	    
	    int[] test = new int[]{3,2,2,3};
	    
	    removeElement(test, 3);
		System.out.println("Hello World");
		
	}
	
    public static int removeElement(int[] nums, int val) {
        
        int count = 0;
        int[] temp = new int[nums.length];
        for(int i = 0; i < nums.length; i++) {
            if(nums[i] == val) {
                temp[count] = i;
                count++;
            }
        }
        
        System.out.println("Count: " + count);

        for(int i = 0; i < temp.length; i++) {
            System.out.println("Index: " + temp[i]);
        }
        
        int[] newArr = new int[nums.length - count];
        int index = 0;
        for(int i = 0, k=0; i < nums.length; i++) {
            

            if(temp[index] == i) {

                index++;
                continue;
            }
            
            newArr[k++] = nums[i]; 
            
            if(index == count) {
                break;
            }
            
        }
        
        for(int i = 0; i < newArr.length; i++) {
            System.out.println("VAl: " + newArr[i]);
        }

        return count;
    }
}
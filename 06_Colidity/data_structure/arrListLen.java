
/*
Reference:
	- https://app.codility.com/programmers/trainings/7/arr_list_len/


Case 1: 	[1, 4, -1, 3, 2]


Suy luan logic nhu sau:

A[0] = 1 => count = 1
A[1] = 4 => count = 2
A[4] = 2 => count = 3
A[2] = -1 redirect to A[4] => count = 4

Vay nen ta se set count = 1 tu luc dau cho truong hop co gia tri la -1 => Refer: <-------
*/
package main;

class Solution {
    public static void main(String []Args) {
  
        int []arr = new int[] {1, 4, -1, 3, 2};
//        int []arr = new int[] {2,-1,1,0,0,0};

        System.out.println(solution(arr));
    }

    public int solution(int[] A) {
        // write your code in Java SE 8
        int next = 0;
        int count = 1; //  <-------
        
        while(A[next] != -1) {
            next = A[next];
            count++;
            if(next == -1)
                break;
        }

        return count;
    }
}

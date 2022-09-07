/* Reference:
 * 
 * Time complexity: O(N^2)
 */
class Solution {
    public int solution(int K, int[] A) {
        // write your code in Java SE 8
        int count = 0;
        for(int i = 0; i < A.length; i++) {
            int max = A[i];
            int min = A[i];
            for(int j = i; j < A.length; j++) {
                
                if(A[j] > max) {
                    max = A[j];
                }

                if(A[j] < min) {
                    min = A[j];
                }

                int result = max - min;
                if(result <= K) {
                    count++;
                    if (count == 1000000000) {
                        return count;
                    }
                } else {
                    break;
                }
            }
        }
        return count;
    }
}
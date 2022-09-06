class Solution {
    public int solution(int[] A) {
        // write your code in Java SE 8
        int max = 0;
        for(int i = 0; i< A.length; i++) {
            if(A[i] > max) {
                max = A[i];
            }
        }
        if (max >= A.length -1 ) {
            return max;
        } else {
            return A.length;
        }
    }
}

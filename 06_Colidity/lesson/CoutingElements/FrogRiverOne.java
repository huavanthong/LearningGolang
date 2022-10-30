public class Main
{
	public static void main(String[] args) {
	    int[] arr = new int[]{1};
	    solution(1, arr);
	}
	
	 public int solution(int X, int[] A) {
        
        int[] B = new int[X];

        int step = 0;
        for(int i =0; i < A.length; i++) {
            if(B[A[i] - 1] == 0 && A[i] <= X) {
                step++;
                B[A[i] - 1] = 1;
            }
            if(step == X) {
                return i;
            }
        }
        
        return -1;
    }

    // Case fail: [1]
	public static int solution_not_pass_all_cases(int X, int[] A) {
        // write your code in Java SE 8
        if(A.length < 1 && A.length > 100000) {
            return -1;
        }
        if(X > A.length) {
            return -1;
        }
        
        int[] arr = new int[A.length];
        for(int i =0; i < arr.length; i++) {
            System.out.println("Check val1: " + arr[i]);
        }


        int index = -1;
        for(int i = 0; i < A.length; i++) {
            arr[A[i] - 1]++;
            if(A[i] == X) {
                index = i;
            }

        }
        for(int i =0; i < arr.length; i++) {
            System.out.println("Check val: " + arr[i]);
        }
        
        for(int i = 0; i< X; i++) {
            if(arr[i] == 0) {
                index = -1;
            }
        }
                    
        System.out.println("Check index: " + index);
        
        return index;
    }
}
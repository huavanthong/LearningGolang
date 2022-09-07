public class Solution {

    static void decToBinary(int n) {

        int[] binaryNum = new int[1000];

        // counter for binary array 
        int i = 0;
        while (n > 0) {
            // storing remainder in binary array 
            binaryNum[i] = n % 2;
            n = n / 2;
            i++;
        }
        int ctr = 0, k = 0;
        ArrayList<Integer> al = new ArrayList<Integer>();

        // printing binary array in reverse order 
        for (int j = i - 1; j >= 0; j--) {
            if (binaryNum[j] == 0) {
                k = j;
                do {
                    ctr++;
                    k++;
                } while (binaryNum[k] == 0);
                al.add(ctr);
                ctr = 0;
            }
        }

        for (int ii = 0; ii < al.size(); ii++) {
            System.out.println(al.get(ii));
        }
    }

    // driver program 
    public static void main(String[] args) {
        int n = 1041;
        decToBinary(n);
    }
}

class Solution {
    public int solution(int N) {
        
        int[] binary = new int[1000];
        int i = 0;
        while(N > 0) {
            binary[i] = N % 2;
            N = N /2;
            i++;
        }

        int max = 0;
        int flag =0 ;

        for (int j = i - 1; j>=0; j--) {
            int count = 0;
            if(binary[j] == 1) {
                for(int z = j - 1; z>=0; z--) {
                    if(binary[z] == 0 ) {
                        count++;
                    }
                    
                    if(binary[z] == 1) {
                        flag = 1;
                        break;
                    }

                }
            }

            if (count >= max && flag == 1) {
                    max = count;
                }
            
        }
        return max;
    }
}

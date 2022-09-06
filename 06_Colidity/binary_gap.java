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
            System.out.print(binaryNum[j]);
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
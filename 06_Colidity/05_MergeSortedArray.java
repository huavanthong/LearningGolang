package main;

public class Main
{
	public static void main(String[] args) {
		System.out.println("Hello World");
		int[] nums1 = new int[]{1,2,3,0,0,0};
		int m = 3;
		
		int[] nums2 = new int[]{2,5,6};
		int n = 3;
		merge(nums1, m, nums2, n);

	}
	
	  public static void merge(int[] nums1, int m, int[] nums2, int n) {
        
        for(int i = m, j = 0; i < m+n; i++) {
            nums1[i] = nums2[j];
            j++;
        }
        
        for(int i =0; i <6; i++) {
            System.out.println("Val: " +nums1[i]);
        }
        
        for(int i = 0; i < m+n - 1; i++) {
            int  min_id = i;
            for(int j = i+1; j < m+n; j++) {
                if(nums1[j] < nums1[min_id]) {
                    min_id = j;
                }
                
                int temp = nums1[min_id];
                nums1[min_id] = nums1[i];
                nums1[i] = temp;
            }
        }
        
        for(int i =0; i <6; i++) {
            System.out.println("ValCheck: " +nums1[i]);
        }
	  }
}


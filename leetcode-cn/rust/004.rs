impl Solution {
    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let mut i1 = 0;
        let mut i2 = 0;
        let num_count = (nums1.len()+nums2.len()+1)/2;
        let mut count = 0;
        let mut nums:Vec<i32> = Vec::new();
        if (nums1.len()==0||nums2.len()==0){
            nums.extend_from_slice(nums1.as_slice());
            nums.extend_from_slice(nums2.as_slice());
        } else {
            loop {
                if nums1[i1]<nums2[i2] {
                    nums.push(nums1[i1]);
                    i1 = i1+1;
                } else {
                    nums.push(nums2[i2]);
                    i2 = i2+1;
                }
                if count>=num_count {break;}
                if i1>=nums1.len(){
                    nums.extend_from_slice(& nums2.split_at(i2).1);
                    break;
                }
                if i2>=nums2.len(){
                    nums.extend_from_slice(& nums1.split_at(i1).1);
                    break;
                }
                count+=1;
            }
        }
        
        return if (nums1.len()+nums2.len())%2>0 {
            nums[(nums1.len()+nums2.len())/2] as f64
        } else {
            ((nums[num_count-1] as f64)+(nums[num_count]as f64))/(2 as f64)            
        };
    }
}
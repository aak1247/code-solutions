impl Solution {
    pub fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
        let mut count:i32 = 0;
        let mut max:i32 = 0;
        for num in nums.iter() {
            if num > &0 {
                count+=1;
                if count > max {
                    max = count;
                }
            } else {
                count=0;
            }
        }
        return max;
    }
}
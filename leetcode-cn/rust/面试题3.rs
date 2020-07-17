impl Solution {
    pub fn find_repeat_number(nums: Vec<i32>) -> i32 {
        let mut a = Vec::new();
        a.extend_from_slice(nums.as_slice());
        a.sort();
        for i in 0..(a.len() - 1) as usize {
            if a[i]==a[i+1] {
                return a[i];
            }
        }
        return -1;
    }
}
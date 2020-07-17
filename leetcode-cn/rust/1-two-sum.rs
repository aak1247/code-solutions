use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map: HashMap<i32,i32> = HashMap::new();
        for i in 0..nums.len() {
            let cur = nums[i];
            let tar = target-cur;
            if map.contains_key(&tar) {
                let j = map.get(&tar).unwrap();
                if i!=(*j as usize) {
                    return vec![*j as i32,i as i32];
                }
            }
            map.insert(cur,i as i32);
        }
        return vec![0,1];
    }
}
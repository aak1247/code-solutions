impl Solution {
    pub fn max_coins(nums: Vec<i32>) -> i32 {
        let N = nums.len();
        let mut nums1 = Vec::new();
        nums1.extend_from_slice(nums.as_slice());
        nums1.insert(0,1);
        nums1.push(1);
        let len = nums1.len();
        let mut coins = vec![vec![0;len as usize];len as usize];
        for c in 1..=(N as usize) {
            for i in 1..=(N+1-c as usize) {
                let j:usize = i+c-1;
                for k in (i as usize)..=(j as usize) {
                    let t2 = nums1[i-1]*nums1[k]*nums1[j+1]+coins[i][k-1]+coins[k+1][j];
                    coins[i][j]=max(coins[i][j],t2);
                }
            }

        }
        coins[1][N]
    }
}

pub fn max(a:i32,b:i32)->i32{
    return if a>b {a}else{b};
}
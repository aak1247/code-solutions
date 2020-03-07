impl Solution {
    pub fn distribute_candies(candies: i32, num_people: i32) -> Vec<i32> {
        let mut res = vec![0;num_people as usize];
        let mut rmn = candies;
        let mut n = 1;
        'outer: loop {
            for i in 0..num_people as usize {
                if rmn <= n {
                    res[i]+=rmn;
                    break 'outer;
                } else {
                    res[i]+=n;
                    rmn-=n;
                    n+=1;
                }
            }
        }
        return res;
    }
}
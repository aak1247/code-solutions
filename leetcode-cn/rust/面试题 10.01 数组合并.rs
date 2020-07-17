impl Solution {
    pub fn merge(a: &mut Vec<i32>, m: i32, b: &mut Vec<i32>, n: i32) {
        for i in 0usize..(n as usize) {
            a[(m as usize)+i]=b[i];
        }
        a.sort();
    }
}
impl Solution {
    pub fn largest_perimeter(a: Vec<i32>) -> i32 {
        let mut b = a.to_vec();
        b.sort();
        b.reverse();
        for i in (0..b.len()-2) {
            let t = &b[i..=i+2];
            if t[0]<t[1]+t[2] {
                return t[0]+t[1]+t[2];
            }
        }
        return 0;
    }
}
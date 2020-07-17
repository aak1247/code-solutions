impl Solution {
    pub fn my_pow(x: f64, n: i32) -> f64 {
        if n == i32::min_value() {
            return 1.0/x * Solution::my_pow(x,n+1);
        }
        if x == 0.0000 {return 0.0}
        if x == 1.0 { return 1.0 }
        if n == 1 { return x }
        if n == 0 { return 1.0 }
        if n < 0 {return Solution::my_pow(1.0/x, -n)}
        if n % 2 == 0  {
            let r = Solution::my_pow(x, n/2);
            return r*r;
        } else {
            return Solution::my_pow(x, n-1) * x
        }
    }
}
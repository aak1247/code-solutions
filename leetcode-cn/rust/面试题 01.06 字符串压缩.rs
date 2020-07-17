impl Solution {
    pub fn compress_string(s: String) -> String {
        let mut s1 = String::new();
        let mut c = '';
        let mut count = 0;
        for i in c.as_mut_vec() {
            if c == (i as char) {
                count+=1;
            } else {
                s1.push(c);
                s1.push(count.to_string());
                count=0;
            }
            let c = (i as char);
        }
        return s1;
    }
}
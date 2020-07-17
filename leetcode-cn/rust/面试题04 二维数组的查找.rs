impl Solution {
    pub fn find_number_in2_d_array(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        if matrix.len()==0 || matrix[0].len()==0||matrix[0][0]>target {return false;}
        for i in matrix {
            for j in i {
                if  j== target {
                    return true;
                }
            }
        }
        return false;
    }
}
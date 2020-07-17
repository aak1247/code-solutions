impl Solution {
    pub fn exist(board: Vec<Vec<char>>, word: String) -> bool {
        if word=="" || board.len()==0||board[0].len()==0 || word.len() > board.len() * board[0].len() {
            return false;
        }
        let mut path = vec![vec![false;board[0].len()]; board.len()];
        for i in 0..board.len() as usize {
            for j in 0..board[i].len() as usize {
                if board[i][j] == word.chars().next().unwrap() {
                    let mut chars = word.chars();
                    if Solution::dfs(&board, i as i32, j as i32, path.clone(), chars.clone()) {return true;}
                } 
            }
        }
        return false;
    }
    pub fn dfs(board: &Vec<Vec<char>>, i:i32, j:i32, mut visited: Vec<Vec<bool>>, mut partial:std::str::Chars) -> bool {
        if partial.clone().next() == None {
            return true;
        }
        let cur = partial.next().unwrap();
        if i < 0 || j < 0 || i == board.len() as i32 || j == board[i as usize].len() as i32 || visited[i as usize][j as usize] || board[i as usize][j as usize] != cur {
            return false;
        }

        visited[i as usize][j as usize]=true;
        let r = Solution::dfs(board, i+1, j, visited.clone(), partial.clone())
                || Solution::dfs(board, i-1,j,visited.clone(), partial.clone())
                || Solution::dfs(board, i, j-1, visited.clone(), partial.clone())
                || Solution::dfs(board, i, j+1, visited.clone(), partial.clone());
        return r;
    }
}
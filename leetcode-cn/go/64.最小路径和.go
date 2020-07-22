func minPathSum(grid [][]int) int {
    var dp = make([][]int, len(grid))
    for i := range dp {
        dp[i] = make([]int, len(grid[i]))
    }
    dp[0][0]=grid[0][0]
    for i:=0 ; i < len(grid) ; i++ {
        
        for j := 0 ; j < len(grid[i]) ; j++ {
            if j == 0 {
                if i == 0 {
                    continue 
                }
                dp[i][j]=dp[i-1][j]+grid[i][j]
                continue
            }
            if i == 0 {
                dp[i][j]=dp[i][j-1]+grid[i][j]
                continue
            }
            dp[i][j]=min(dp[i-1][j], dp[i][j-1])+grid[i][j]
        }
    }
    return dp[len(dp)-1][len(dp[0])-1]
}

func min(a, b int) int {
    if a<b {
        return a
    }
    return b
}
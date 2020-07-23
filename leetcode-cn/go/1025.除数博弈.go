// n2 暴力
func divisorGame(N int) bool {
    if N == 2 {
        return true
    }
    if N == 1 {
        return false
    }
    for i := 2 ; i < N ; i++ {
        if N % i == 0 && divisorGame(N/i){
            return true
        }
    }
    return divisorGame(N-2)
}

// n2 dp 小常数
func divisorGame2(N int) bool {
	if N == 1 {
			return false
	}
	if N == 2 {
			return true
	}
	var dp = make([]bool, N+1)
	dp[1] = false
	dp[2] = true
	for i := 3 ; i <= N ; i++ {
		dp[i] = dp[i-2]||dp[i]
		t := N / i
		for j := 2 ; j <= t ; j++ {
			dp[j*i]=dp[(j-1)*i]||dp[i]
		}
	}
	return dp[N]
}

// 奇数开局下一轮总会回到奇数，最终输
// 偶数开局 总可以给1 对方拿奇数输
// o(1)
func divisorGame3(N int) bool {
	return N%2==0
}
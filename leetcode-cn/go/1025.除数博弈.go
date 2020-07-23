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

func integerBreak(n int) int {
    buffer := make([]int, n+1)
    for i:=0;i<=n;i++ {
        cur := 0
        for j:=1;j<=i/2;j++{
            t :=  max(buffer[j],j)*max(buffer[i-j],i-j)
            if cur < t {
                cur = t
            }
        }
        buffer[i]=cur
    }
    return buffer[n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
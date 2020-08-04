const m = 1000000000 + 7

func numSub(s string) int {
    return int(numByteSub([]byte(s))%m)
}

func numByteSub(s []byte) int64 {
    if len(s) == 0 {
        return 0
    }
    if len(s) == 1 {
        if s[0] == []byte("1")[0] {
            return 1
        } else {
            return 0
        }
    }
    //判断是否全1
    for i, b := range s {
        if b != []byte("1")[0] {
            if i == len(s)-1 {
                return numByteSub(s[:i])
            } else {
                return numByteSub(s[:i]) + numByteSub(s[i+1:])
            }
        }
    }
    
    //对于全1的

    return int64(len(s))*int64(len(s)+1)/2%m
}
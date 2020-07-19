func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1)+len(s2)!=len(s3) {
        return false
    }
    b1, b2, b3 := []byte(s1), []byte(s2), []byte(s3)
    return isBytesInterleave(b1, b2, b3)
}

func isBytesInterleave(b1, b2, b3 []byte) bool {
    if len(b3) == 0 {
        return len(b1)==0&&len(b2)==0
    }
    if len(b1)!=0 && b3[0] == b1[0] && isBytesInterleave(b1[1:], b2, b3[1:]) {
        return true
    }
    if len(b2)!=0 && b3[0] == b2[0] && isBytesInterleave(b1, b2[1:], b3[1:]) {
        return true
    }
    return false
}
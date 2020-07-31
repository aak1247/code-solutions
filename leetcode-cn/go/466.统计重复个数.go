// 记忆化
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
    res := make([][]int, 0)
    for i := range make([]int, len(s2)) {
        // 记录从i开始时的状态
        start, times := i, 0
        for j := range make([]int, len(s1)) {
            if s2[start%len(s2)]==s1[j] {
                // 循环匹配，如果s1长 则times大于1; 
                // 记录当前匹配在s2上的结束位置，下一轮最早即从此位置开始匹配
                start++
                if start%len(s2)==0 {
                    times++
                }
            }
        }
        res=append(res, []int{start%len(s2), times})
    }
    // 贪心求解每一步，下一次从最早可开始的位置开始进行匹配，则总匹配数一定是最多的
    // n1步后，剩余数据必然不够增加一次（如果够，则times就该+1了）
    start, times := 0, 0
    for range make([]int, n1) {
        times += res[start][1]
        start = res[start][0]
    }
    return times/n2
}

func getMaxRepetitions2(s1 string, n1 int, s2 string, n2 int) int {
	len1, len2 := len(s1), len(s2)
	index1, index2 := 0, 0 // 注意此处直接使用 Ra Rb 的下标，不取模

	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return 0
	}

	map1, map2 := make(map[int]int), make(map[int]int)
	ans := 0 // 注意，此处存储的是 Ra 中 Sb 的个数，而非 Ra 中 Rb 的个数

	for index1/len1 < n1 { // 遍历整个 Ra
		if index1%len1 == len1-1 { //在 Sa 末尾
			if val, ok := map1[index2%len2]; ok { // 出现了循环，进行快进
				cycleLen := index1/len1 - val/len1                 // 每个循环占多少个 Sa
				cycleNum := (n1 - 1 - index1/len1) / cycleLen      // 还有多少个循环
				cycleS2Num := index2/len2 - map2[index2%len2]/len2 // 每个循环含有多少个 Sb

				index1 += cycleNum * cycleLen * len1 // 将 Ra 快进到相应的位置
				ans += cycleNum * cycleS2Num         // 把快进部分的答案数量加上
			} else { // 第一次，注意存储的是未取模的
				map1[index2%len2] = index1
				map2[index2%len2] = index2
			}

		}

		if s1[index1%len1] == s2[index2%len2] {
			if index2%len2 == len2-1 {
				ans += 1
			}
			index2 += 1
		}
		index1 += 1
	}
	return ans / n2
}
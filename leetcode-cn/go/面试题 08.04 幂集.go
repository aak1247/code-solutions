func subsets(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{[]int{}}
    }
    var res = make([][]int, 0)
    res = append(res, []int{})

    for _, num:= range nums {
        for _, cur := range res {
            var tNums = make([]int, 1)
            tNums[0] = num
            tNums = append(tNums, cur...)
            res=append(res,tNums)
        }
        
    }
    return res
}


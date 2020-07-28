func numIdenticalPairs(nums []int) int {
    var num int
    if len(nums)<=1 {
        return 0
    }
    for i:=1;i<len(nums);i++ {
        if nums[i]==nums[0] {
            num++
        }
    }
    if len(nums)>1 {
        num += numIdenticalPairs(nums[1:])
    }
    return num
}
func findMagicIndex(nums []int) int {
    for i := range nums {
        if nums[i]==i {
            return i
        }
    }
    return -1
}
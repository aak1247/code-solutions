package main

func searchInsert(nums []int, target int) int {
    for i, num := range nums {
        if target <= num {
            return i
        }
    }
    return len(nums)
}

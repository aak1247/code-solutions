package main

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		index := indexOf(nums, left, right)
		quickSort(nums, 0, index-1)
		quickSort(nums, index+1, right)
	}
}
func indexOf(nums []int, left, right int) int {
	tmp := nums[left]
	for left < right {
		for left < right && nums[right] >= tmp {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= tmp {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = tmp
	return left
}

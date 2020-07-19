func twoSum(numbers []int, target int) []int {
    var i1, i2 = 0, len(numbers)-1
    for {
        if numbers[i1] + numbers[i2] > target {
            i2--
        }
        if numbers[i1] + numbers[i2] < target {
            i1++
        }
        if numbers[i1] + numbers[i2] == target {
            break
        }
    }
    return []int{i1+1,i2+1}
}
package bst

// 例1. 插入位置
// 给定一个排序数组nums（无重复元素）与目标值target，如果target在nums里出现，则返回target所在下标，如果target在nums里未出现，则返回
// target应该插入位置的数组下标，使得将target插入数组nums后，数组仍有序。
func SearchInsert(nums []int, target int) int {
	index := -1
	begin := 0
	end := len(nums) - 1

	for index == -1 {
		mid := (begin + end) / 2
		if target == nums[mid] {
			index = mid
		} else if target < nums[mid] {
			if mid == 0 || target > nums[mid-1] {
				index = mid
			}
			end = mid - 1
		} else {
			if mid == len(nums)-1 || target < nums[mid-1] {
				index = mid + 1
			}
			begin = mid + 1
		}
	}
	return index
}

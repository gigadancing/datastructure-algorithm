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

// 例2. 区间查找
// 给定一个排序数组nums（nums中有）与目标值target，如果target在nums中出现，则返回target所在区间的左右端点下标[左端点,右端点]，如果
// target在nums中未出现，则返回[-1,-1]
// 例如：
// nums=[5,7,7,8,8,8,8,10]
// target=8，返回：[3,6]
// target=6，返回：[-1,-1]
func SearchRange(nums []int, target int) []int {
	res := make([]int, 0)
	left := LeftBound(nums, target)
	right := RightBound(nums, target)
	res = append(res, left)
	res = append(res, right)
	return res
}

func LeftBound(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	for begin <= end {
		mid := (begin + end) / 2
		if target == nums[mid] {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			end = mid - 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return -1
}

func RightBound(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1

	for begin <= end {
		mid := (begin + end) / 2
		if target == nums[mid] {
			if mid == len(nums)-1 || target < nums[mid+1] {
				return mid
			}
			begin = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return -1
}

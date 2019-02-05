package bst

// 二分查找递归
// 返回值：
// true - 找到
// false - 未找到
func BinarySearchRecursion(sortedArray []int, begin, end, target int) bool {
	if len(sortedArray) == 0 || begin > end {
		return false
	}
	mid := (begin + end) / 2
	if target == sortedArray[mid] {
		return true
	} else if target < sortedArray[mid] {
		return BinarySearchRecursion(sortedArray, begin, mid-1, target)
	} else {
		return BinarySearchRecursion(sortedArray, mid+1, end, target)
	}
}

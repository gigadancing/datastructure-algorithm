package rtp

// 已知两个已排序数组，将这两个数组合并为一个排序数组。
// [2,5,8,20] [1,3,5,7,30,50]
// [1,2,3,5,5,7,8,20,30,50]
func MergeTwoSortedArray(nums1, nums2 []int) []int {
	result := make([]int, 0)
	var i, j int
	for i, j = 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	// 若i没有到最后一个元素，说明len(nums1)>len(nums2)，将nums1的剩余元素加到result
	if i != len(nums1)-1 {
		result = append(result, nums1[i:]...)
	}
	// 若j没有到最后一个元素，说明len(nums2)>len(nums1)，将nums2的剩余元素加到result
	if j != len(nums1)-1 {
		result = append(result, nums2[j:]...)
	}

	return result
}

// 归并排序
func MergeSort(nums []int) []int {
	if len(nums) < 2 { // 子问题足够小时，直接求解
		return nums
	}
	// 对原数组拆分了两个规模相同的数组，在对它们分别求解
	mid := len(nums) / 2
	// 对拆分后的两个子问题求解
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	// 将子问题的解进行合并
	return mergeSortTwoArray(left, right)
}

func mergeSortTwoArray(left, right []int) (result []int) {
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return
}

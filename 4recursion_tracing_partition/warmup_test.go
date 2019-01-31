package rtp

import (
	"fmt"
	"testing"
)

func TestMergeArray(t *testing.T) {
	nums1 := []int{2, 5, 8, 20}
	nums2 := []int{1, 3, 5, 7, 30, 50}
	result := MergeTwoSortedArray(nums1, nums2)
	fmt.Println(result) //[1,2,3,5,5,7,8,20,30,50]
	nums1 = []int{1, 1, 2, 3, 5, 10, 400}
	nums2 = []int{4, 4, 7, 8}
	result = MergeTwoSortedArray(nums1, nums2)
	fmt.Println(result) //[1,1,2,3,4,4,5,7,8,10,400]
}

func TestMergeSort(t *testing.T) {
	nums := []int{4, 9, 1, 19, 80, 34, 67, 22, 504, 402, 116}
	res := MergeSort(nums)
	fmt.Println(res)
}

func TestCountSmaller(t *testing.T) {
	nums := []int{5, 2, 6, 1}
	res := CountSmaller(nums)
	fmt.Println(res) //[2,1,1,0]
	nums = []int{6, 6, 6, 1, 1, 1}
	res = CountSmaller(nums)
	fmt.Println(res) //[3,3,3,0,0,0]
	nums = []int{5, -7, 9, 1, 3, 5, -2, 1}
	res = CountSmaller(nums)
	fmt.Println(res) //[5,0,5,1,2,2,0,0]
}

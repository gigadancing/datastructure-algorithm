package rtp

import (
	"fmt"
	"testing"
)

func TestMergeArray(t *testing.T) {
	nums1 := []int{2, 5, 8, 20}
	nums2 := []int{1, 3, 5, 7, 30, 50}
	result := MergeArray(nums1, nums2)
	fmt.Println(result) //[1,2,3,5,5,7,8,20,30,50]
	nums1 = []int{1, 1, 2, 3, 5, 10, 400}
	nums2 = []int{4, 4, 7, 8}
	result = MergeArray(nums1, nums2)
	fmt.Println(result) //[1,1,2,3,4,4,5,7,8,10,400]
}

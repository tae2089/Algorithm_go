package leetcode

import "sort"

func merge88(nums1 []int, m int, nums2 []int, n int) []int {
	if len(nums1) > m {
		nums1 = nums1[:m]
	}
	if len(nums2) > n {
		nums2 = nums2[:n]
	}
	for _, v := range nums2 {
		nums1 = append(nums1, v)
	}
	sort.Ints(nums1)
	return nums1
}

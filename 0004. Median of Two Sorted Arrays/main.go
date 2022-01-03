package main

import "fmt"

// 比较中规中矩的方法，每次排除最左右两边的数（另一种是归并之后取中位数，效率差不多）
//func median(nums []int) (int, float64) {
//	if len(nums)%2 == 1 {
//		return (len(nums) - 1) / 2, float64(nums[(len(nums)-1)/2])
//	} else {
//		return (len(nums) - 2) / 2, (float64(nums[(len(nums)-2)/2]) + float64(nums[len(nums)/2])) / 2
//	}
//}

func median(nums []int) float64 {
	if len(nums)%2 == 1 {
		return float64(nums[(len(nums)-1)/2])
	} else {
		return (float64(nums[(len(nums)-2)/2]) + float64(nums[len(nums)/2])) / 2
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return median(nums2)
	} else if len(nums2) == 0 {
		return median(nums1)
	} else if len(nums1) == 1 && len(nums2) == 1 {
		return float64(nums1[0] + nums2[0]) / 2
	} else {
		if nums1[0] < nums2[0]{
			nums1 = nums1[1:]
		} else {
			nums2 = nums2[1:]
		}
		if len(nums1) == 0{
			return median(nums2[:len(nums2)-1])
		}
		if len(nums2) == 0{
			return median(nums1[:len(nums1)-1])
		}
		if nums1[len(nums1)-1] < nums2[len(nums2)-1]{
			nums2 = nums2[:len(nums2)-1]
		} else {
			nums1 = nums1[:len(nums1)-1]
		}
		return findMedianSortedArrays(nums1, nums2)
	}
}

// 更好的方法其实就是找出两个有序数组的第k个最小数的特殊情况，利用二分法每次排除掉一半的元素

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,2,4,6,8,9,10}, []int{3,5,7}))
}

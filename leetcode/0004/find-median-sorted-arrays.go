package _004

import "fmt"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	n := l1 + l2

	hasEven := n%2 == 0
	var middle = n/2 + 1

	var nums1Index, nums2Index int
	var val, preVal int

	//12 3
	//12 34
	//1
	fmt.Println("nums1", nums1, "nums2", nums2)
	for i := 0; i < middle; i++ {

		fmt.Println("i", i, "nums1Index", nums1Index, "nums2Index", nums2Index)
		if nums1Index < l1 && nums2Index < l2 {
			if nums1[nums1Index] > nums2[nums2Index] {
				preVal, val = val, nums2[nums2Index]
				nums2Index++
			} else if nums1[nums1Index] < nums2[nums2Index] {
				preVal, val = val, nums1[nums1Index]
				nums1Index++
			} else {
				preVal, val = val, nums1[nums1Index]
				nums1Index++
			}
		} else if nums1Index < l1 {
			preVal, val = val, nums1[nums1Index]
			nums1Index++
		} else if nums2Index < l2 {
			preVal, val = val, nums2[nums2Index]
			nums2Index++
		}

	}

	fmt.Println(preVal, val, middle)

	if hasEven {
		return float64(val+preVal) / 2
	}
	return float64(val)

}

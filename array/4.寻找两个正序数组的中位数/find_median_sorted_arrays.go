package __寻找两个正序数组的中位数

// findMedianSortedArrays1 双指针比较两数组后移
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	var (
		l1              = len(nums1)
		l2              = len(nums2)
		hasEven         = (l1+l2)%2 == 0
		currNum, preNum int
		middle          = (l1+l2)/2 + 1
		n1idx, n2idx    int
	)

	for i := 0; i < middle; i++ {
		if n1idx < l1 && n2idx < l2 {
			if nums1[n1idx] <= nums2[n2idx] {
				preNum, currNum = currNum, nums1[n1idx]
				n1idx++
			} else {
				preNum, currNum = currNum, nums2[n2idx]
				n2idx++
			}
		} else if n1idx < l1 {
			preNum, currNum = currNum, nums1[n1idx]
			n1idx++
		} else if n2idx < l2 {
			preNum, currNum = currNum, nums2[n2idx]
			n2idx++
		}
	}

	if hasEven {
		return float64(preNum+currNum) / 2
	}

	return float64(currNum)

}

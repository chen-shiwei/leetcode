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

// findMedianSortedArrays2 二分法
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	hasEven := l%2 == 0
	mid := l / 2

	if hasEven {
		// 偶数的中间数为
		return float64(FindKthElement(nums1, nums2, mid)+FindKthElement(nums1, nums2, mid+1)) / 2.0
	} else {
		// 奇数的中间数为 mid
		return float64(FindKthElement(nums1, nums2, mid+1))
	}
}

// FindKthElement
// 两个有序数组
// 二分查找从小到大第k个数
// 比较 二分 开始比较nums1[0:(k/2)]和nums2[0:(k/2)]
func FindKthElement(nums1, nums2 []int, k int) int {
	var idx1, idx2 int
	for {
		// 越界后，直接取另一个kth 数
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}
		mid := k / 2
		newIdx1 := min(idx1+mid, len(nums1)) - 1
		newIdx2 := min(idx2+mid, len(nums2)) - 1
		n1, n2 := nums1[newIdx1], nums2[newIdx2]
		if n1 <= n2 {
			// 索引得+1
			k -= newIdx1 - idx1 + 1
			idx1 = newIdx1 + 1
		} else {
			k -= newIdx2 - idx2 + 1
			idx2 = newIdx2 + 1
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

package _088

// merge 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m + n - 1
	m = m - 1
	n = n - 1
	for m >= 0 || n >= 0 {
		// nums1 没值了 让 nums2
		if nums1[m] < 0 {
			nums1[l] = nums2[n]
			n--
			l--
			continue
		}

		// nums2 没值了 让 nums1
		if nums2[n] < 0 {
			nums1[l] = nums1[m]
			m--
			l--
			continue
		}

		if nums1[m] >= nums2[n] {
			nums1[l] = nums1[m]
			m--
			l--
		} else {
			nums1[l] = nums2[n]
			n--
			l--
		}
	}
	return
}

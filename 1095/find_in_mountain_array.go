package _095

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
type MountainArray struct {
	data []int
}

func NewMountainArray(data []int) *MountainArray {
	return &MountainArray{data: data}
}

func (this *MountainArray) get(index int) int {
	return this.data[index]
}
func (this *MountainArray) length() int {
	return len(this.data)

}

func findInMountainArray(target int, m *MountainArray) int {
	l := m.length()

	var left, right = 0, l - 1
	// 找出最高点 peak,一分为二
	// left->peak 有序
	// peak->right 有序
	for left+1 < right {
		mid := left + (right-left)/2
		midVal := m.get(mid)
		if m.get(mid-1) < midVal {
			left = mid
		} else {
			right = mid
		}
	}

	var peakIdx = left
	if m.get(right) > m.get(left) {
		peakIdx = right
	}

	idx := BinarySort(m, 0, peakIdx, target, true)
	if idx == -1 {
		idx = BinarySort(m, peakIdx+1, l-1, target, false)
	}

	return idx
}

func BinarySort(m *MountainArray, left, right int, target int, asc bool) int {

	for left <= right {
		mid := left + (right-left)/2
		midVal := m.get(mid)
		if midVal == target {
			return mid
		}
		if midVal < target {
			if asc {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if asc {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

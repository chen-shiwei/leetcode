package 判断小数组是否是大数组的连续子集

func isArraySub(bigNums []int, smallNums []int) int {
	bl := len(bigNums)
	sl := len(smallNums)

	var i, j int

	var l int
	var startIdx int
	for i < bl && j < sl {
		if bigNums[i] == smallNums[j] {
			if l == 0 {
				startIdx = i
			}
			l++
			// 同步 small、big 递增idx
			i++
			j++
		} else {
			i++
			l = 0
			// 重置 small idx
			j = 0
		}
	}

	if l == sl {
		return startIdx
	}

	return -1
}

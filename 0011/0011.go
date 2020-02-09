package _011

func MaxArea(height []int) int {
	l := len(height)
	i, j, result := 0, l-1, 0

	for i < j {
		if height[i] <= height[j] {
			result = max(result, (height[i])*(j-i))
			i++
		} else {
			result = max(result, (height[j])*(j-i))
			j--
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

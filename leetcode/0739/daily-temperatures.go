package _739

// DailyTemperatures
// https://leetcode-cn.com/problems/daily-temperatures
func DailyTemperatures(T []int) []int {
	l := len(T)

	var result []int

	var count int
	for i := 0; i < l; i++ {
		var ok bool
		var now = T[i]
		for j := i + 1; j < l; j++ {
			count++
			if T[j] > now {
				ok = true
				break
			}
		}
		if ok {
			result = append(result, count)
		} else {
			result = append(result, 0)

		}
		count = 0
	}

	return result

}

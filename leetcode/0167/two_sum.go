package _167

func twoSumWithDict(numbers []int, target int) []int {

	l := len(numbers)

	var dict = make(map[int]int)

	var result []int
	for i := 0; i < l; i++ {
		n := target - numbers[i]

		v, ok := dict[n]
		if ok {
			result = append(result, v+1, i+1)
			return result
		}
		dict[numbers[i]] = i
	}

	return nil

}

func twoSumWithSortArray(numbers []int, target int) []int {

	l := len(numbers)

	var start, end = 0, l - 1
	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		}
		if sum < target {
			start = start + 1
		} else {
			end = end - 1
		}
	}

	return nil

}

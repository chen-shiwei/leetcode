package offer_11

func minArray(numbers []int) int {
	l := len(numbers)
	var start, end = 0, l - 1

	for start < end {
		if numbers[start] > numbers[start+1] {
			return numbers[start+1]
		}

		if numbers[end] < numbers[end-1] {
			return numbers[end]
		}
		start++
		end--
	}

	return numbers[0]

}

func minArray2(numbers []int) int {
	l := len(numbers)
	var start, end = 0, l - 1

	for start < end {
		mid := start + (end-start)/2

		if numbers[mid] > numbers[end] {
			start = mid + 1
		} else if numbers[mid] < numbers[end] {
			end = mid
		} else {
			end--
		}
	}

	return numbers[start]

}

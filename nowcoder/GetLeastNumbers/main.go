package GetLeastNumbers

func GetLeastNumbers(input []int, k int) []int {
	// write code here
	var l = len(input)
	for i := 0; i < k+1; i++ {
		for j := l - 1; j >= i; j-- {
			if input[j] < input[j-1] {
				input[j], input[j-1] = input[j-1], input[j]
			}
		}
	}

	return input[:k]
}

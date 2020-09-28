package _248

func numOfSubArrays(nums []int, k int) int {
	oddNums := make([]int, len(nums), cap(nums))

	var n int
	for i, v := range nums {
		n = n + (v & 1)
		oddNums[i] = n
	}
	if n == 0 {
		return 0
	}

	dict := make(map[int]int, 0)
	dict[0] = 1
	var result int
	for _, num := range oddNums {
		if _, ok := dict[num-k]; ok {
			result += dict[num-k]
		}
		dict[num]++
	}

	return result

}

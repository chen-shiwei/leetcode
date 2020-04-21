package _560

func subarraySum(nums []int, k int) int {
	var sumNums = make([]int, len(nums), cap(nums))
	var n int
	for i, v := range nums {
		n += v
		sumNums[i] = n
	}

	var dict = make(map[int]int, 0)

	var count int

	dict[0] = 1
	for _, v := range sumNums {
		count += dict[v-k]
		dict[v]++
	}

	return count
}

// subarraySum2 前缀和+穷举
func subarraySum2(nums []int, k int) int {
	l := len(nums)
	sumNums := make([]int, l+1)

	for i := 0; i < l; i++ {
		sumNums[i+1] = sumNums[i] + nums[i]
	}

	var count int
	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if sumNums[i]-sumNums[j] == k {
				count++
			}
		}
	}
	return count
}

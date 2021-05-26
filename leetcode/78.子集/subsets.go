package _8_子集

func subsets(nums []int) [][]int {
	var (
		path  []int
		paths [][]int
		fn    func(nums []int, startIdx int)
	)

	fn = func(nums []int, startIdx int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		paths = append(paths, tmp)
		if startIdx >= len(nums) {
			//tmp := make([]int, len(path))
			//copy(tmp, path)
			//paths = append(paths, tmp)
			return
		}

		for i := startIdx; i < len(nums); i++ {
			path = append(path, nums[i])
			fn(nums, i+1)
			path = path[:len(path)-1]
		}
	}

	fn(nums, 0)

	return paths
}

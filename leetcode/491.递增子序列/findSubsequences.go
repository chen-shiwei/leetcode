package _91_递增子序列

// 不能排序去重，用map 去重，
func findSubsequences(nums []int) [][]int {
	var (
		paths [][]int
		path  []int
		fn    func(nums []int, startIdx int)
	)

	fn = func(nums []int, startIdx int) {

		if len(path) >= 2 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
			if startIdx >= len(nums) {
				return
			}
		}

		var used = make(map[int]bool)
		for i := startIdx; i < len(nums); i++ {
			// 不能排序去重，用map 去重，
			if len(path) != 0 && path[len(path)-1] > nums[i] /*判断递增*/ || used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			path = append(path, nums[i])
			fn(nums, i+1)
			used[i] = false
			path = path[:len(path)-1]
		}

	}

	fn(nums, 0)

	return paths
}

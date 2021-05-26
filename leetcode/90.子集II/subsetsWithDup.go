package _0_子集II

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var (
		used  = make([]bool, len(nums))
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
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			fn(nums, i+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	sort.Ints(nums)

	fn(nums, 0)

	return paths
}

package _7_全排列II

import "sort"

func permuteUnique(nums []int) [][]int {
	var (
		paths [][]int
		path  []int
		used  = make([]bool, len(nums))
		fn    func(nums []int)
	)

	fn = func(nums []int) {
		if len(nums) == len(path) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// used[i-1] == true 同一树枝使用过
			// used[i-1] == false 同一树层使用过
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			fn(nums)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	sort.Ints(nums)
	fn(nums)
	return paths
}

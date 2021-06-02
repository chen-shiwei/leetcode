package _6_全排列

func permute(nums []int) [][]int {
	var (
		paths [][]int
		path  []int
		used  = make([]bool, len(nums))
		fn    func(nums []int)
	)

	fn = func(nums []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, path)
			paths = append(paths, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			fn(nums)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	fn(nums)

	return paths
}

package _9_组合总和

func combinationSum(candidates []int, target int) [][]int {
	var (
		paths [][]int
		path  []int
		sum   int
		fn    func(candidates []int, target int, startIdx int)
	)

	fn = func(candidates []int, target int, startIdx int) {
		if sum > target {
			return
		}

		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
			return
		}

		for i := startIdx; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			fn(candidates, target, i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}

	fn(candidates, target, 0)

	return paths
}

package _0_组合总和II

import "sort"

func combinationSum2(candidates []int, target int) [][]int {

	var (
		use   = make([]bool, len(candidates))
		paths [][]int
		path  []int
		fn    func(candidates []int, target int, startIdx int)
		sum   int
	)
	sort.Ints(candidates)

	fn = func(candidates []int, target int, startIdx int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
			return
		}

		if sum > target || startIdx >= len(candidates) {
			return
		}

		for i := startIdx; i < len(candidates); i++ {
			// 判断树层是够被使用 false
			if i > 0 && candidates[i] == candidates[i-1] && !use[i-1] {
				continue
			}
			/*树支开始*/
			use[i] = true
			path = append(path, candidates[i])
			sum += candidates[i]
			fn(candidates, target, i+1)
			/*树支结束*/
			/*树层开始*/
			use[i] = false
			path = path[:len(path)-1]
			sum -= candidates[i]
			/*树层结束*/
		}

	}

	fn(candidates, target, 0)

	return paths
}

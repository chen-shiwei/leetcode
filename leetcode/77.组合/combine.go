package _7_组合

func combine(n int, k int) [][]int {

	var (
		paths [][]int
		path  []int
	)
	var backtracking func(n, k, startIdx int)

	backtracking = func(n, k, startIdx int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
			return
		}

		for i := startIdx; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(n, k, 1)
	return paths
}

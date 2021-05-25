package _16_组合总和III

func combinationSum3(k int, n int) [][]int {

	var (
		fn    func(k, n, startIdx int)
		paths [][]int
		path  []int
		// 计算sum
		sum int
	)
	fn = func(k, n, startIdx int) {
		if len(path) == k { // 符合3个数
			if sum == n { // 判断总和
				tmp := make([]int, len(path), cap(path))
				copy(tmp, path)
				paths = append(paths, tmp)
			}
			return
		}
		for i := startIdx; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			sum += i
			fn(k, n, i+1)
			path = path[:len(path)-1]
			sum -= i
		}
	}

	fn(k, n, 1)

	return paths
}

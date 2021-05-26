package _31_分割回文串

func partition(s string) [][]string {
	var (
		paths [][]string
		path  []string
		fn    func(s string, startIdx int)
	)
	if s == "" {
		return nil
	}

	fn = func(s string, startIdx int) {

		if startIdx >= len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
			return
		}
		for i := startIdx; i < len(s); i++ {
			if !isPalindrome(s, startIdx, i) {
				continue
			}
			path = append(path, s[startIdx:i+1])
			fn(s, i+1)
			path = path[:len(path)-1]
		}
	}
	fn(s, 0)

	return paths
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

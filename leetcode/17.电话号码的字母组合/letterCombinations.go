package _7_电话号码的字母组合

var letterMap = []string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

func letterCombinations(digits string) []string {
	var (
		paths []string
		path  []byte
		fn    func(digits string, startIdx int)
	)
	if len(digits) == 0 {
		return nil
	}

	fn = func(digits string, startIdx int) {
		if len(path) == len(digits) {
			tmp := make([]byte, len(digits), len(digits))
			copy(tmp, path)
			paths = append(paths, string(tmp))
			return
		}

		world := letterMap[digits[startIdx]-'0']
		// for 横向遍历
		for _, v := range world {
			path = append(path, byte(v))
			// 递归纵向遍历
			fn(digits, startIdx+1)
			// 回溯不断调整结果集
			path = path[:len(path)-1]
		}

	}

	fn(digits, 0)

	return paths
}

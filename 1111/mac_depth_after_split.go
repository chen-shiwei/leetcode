package _111

// MaxDepthAfterSplit 把有效括号分配到A（0），B（1）
func MaxDepthAfterSplit(seq string) []int {

	var res = make([]int, 0, len(seq))
	var d = -1
	var last int32
	for _, word := range seq {
		if d == -1 {
			d += 1
			res = append(res, d%2)
			last = word
			continue
		}

		if last != word {
			res = append(res, d%2)
			last = word
			continue
		}

		if last == word {
			d += 1
			res = append(res, d%2)
			last = word
			continue
		}

	}
	return res
}

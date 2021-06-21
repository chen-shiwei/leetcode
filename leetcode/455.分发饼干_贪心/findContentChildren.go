package _55_分发饼干_贪心

import "sort"

func findContentChildren(g []int, s []int) int {

	sort.Slice(g, func(i, j int) bool {
		return g[i] > g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})

	var (
		idx   = 0
		count int
	)

	for i := 0; i < len(g); i++ {
		if idx < len(s) && s[idx] >= g[i] {
			count++
			idx++
		}
	}

	return count

}

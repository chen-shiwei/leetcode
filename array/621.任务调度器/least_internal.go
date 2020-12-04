package _21_任务调度器

import (
	"sort"
)

func leastInterval(tasks []byte, n int) int {

	l := len(tasks)
	var dict = make([]int, 26)
	for i := 0; i < l; i++ {
		dict[tasks[i]-'A']++
	}
	sort.Ints(dict)

	var c int

	for dict[25] > 0 {
		var i int
		for i <= n {
			if dict[25] == 0 {
				break
			}
			if i < 26 && dict[25-i] > 0 {
				dict[25-i]--
			}
			c++
			i++
		}
		sort.Ints(dict)
	}

	return c

}

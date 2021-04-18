package _21_任务调度器

import (
	"fmt"
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
	fmt.Println(dict)

	for dict[25] > 0 {
		var i int
		for i <= n {
			if dict[25] == 0 {
				break
			}
			if i < 26 && dict[25-i] > 0 {
				dict[25-i]--
				fmt.Println("1")
			}
			c++
			i++
		}
		sort.Ints(dict)
		fmt.Println("1 sort", dict)

	}

	return c

}

func leastInterval1(tasks []byte, n int) int {

	l := len(tasks)
	var dict = make([]int, 26)
	for i := 0; i < l; i++ {
		dict[tasks[i]-'A']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(dict)))
	var c int
	for dict[0] > 0 {
		var i int
		for i <= n {
			if dict[0] == 0 {
				break
			}
			if i < 26 && dict[i] > 0 {
				dict[i]--
				fmt.Println("2")
			}
			c++
			i++
		}
		sort.Sort(sort.Reverse(sort.IntSlice(dict)))
	}

	return c

}

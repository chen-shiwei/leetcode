package cpus

import (
	"fmt"
	"sort"
)

func Schedule(m []int, task []int) {
	sort.Ints(m)

	fmt.Println(m)
	m1 := make([]int, len(m))

	var tasks = make([][3]int, len(task))
	for i := 0; i < len(task); i++ {

		min := m1[0]
		minI := 0
		for j := 1; j < len(m); j++ {
			if min > m1[j] {
				min = m1[j]
				minI = j
			}
		}

		m1[minI] += task[i]
		tasks[i] = [3]int{i + 1, m1[minI], minI + 1}

	}

	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i])
	}
}

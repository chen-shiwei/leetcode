package _406_根据身高重建队列

import (
	"container/list"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	length := len(people)
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]

	})

	l := list.New().Init()

	for i := 0; i < length; i++ {
		if people[i][1] == i {
			l.PushBack(people[i])
			continue
		}
		c := l.Front()
		for j := 0; j < people[i][1]; j++ {
			if c != nil {
				c = c.Next()
			} else {
				break
			}
		}
		l.InsertBefore(people[i], c)
	}

	cur := l.Front()
	var s [][]int
	for i := 0; i < length; i++ {
		s = append(s, cur.Value.([]int))
		cur = cur.Next()
	}
	return s
}

package _5_三数之和

import (
	"fmt"
	"sort"
)

func threeSum1(nums []int) [][]int {
	l := len(nums)

	var dict = make(map[int][]int)
	var existRecord = make(map[string]bool)
	var result [][]int

	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			v, ok := dict[nums[j]]
			fmt.Println("mark ok ", dict[nums[j]], j)
			if ok {
				fmt.Println("mark ok", v)
				subNums := append(v, nums[j])
				sort.Ints(subNums)
				key := fmt.Sprintf("%d,%d,%d", subNums[0], subNums[1], subNums[2])
				if _, ok := existRecord[key]; ok {
					continue
				}
				existRecord[key] = true
				result = append(result, subNums)
				continue
			}
			mark := -(nums[i] + nums[j])
			dict[mark] = []int{nums[i], nums[j]}
			fmt.Println("mark", dict, j < l-1)

		}
	}

	return result
}

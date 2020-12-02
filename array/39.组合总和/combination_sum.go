package _9_组合总和

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	// 排序数组，才能使用剪枝功能
	sort.Ints(candidates)
	var result [][]int

	var dfs func(target int, nums []int, subResult []int)
	dfs = func(target int, nums []int, subResult []int) {
		if target == 0 {
			result = append(result, append([]int{}, subResult...))
		}

		for i, num := range nums {
			newTarget := target - num
			if newTarget < 0 {
				return
			}
			dfs(newTarget, nums[i:], append(subResult, num))

		}
	}

	dfs(target, candidates, []int(nil))

	return result
}

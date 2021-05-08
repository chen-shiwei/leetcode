package _47_前K个高频元素

import (
	"container/heap"
	"github.com/chen-shiwei/leetcode/leetcode/datastruct"
)

func topKFrequent(nums []int, k int) []int {
	var dict = make(map[int]int)

	for _, v := range nums {
		dict[v]++
	}

	var freqNums [][]int

	for key, value := range dict {
		freqNums = append(freqNums, []int{key, value})
	}

	QSort(freqNums)
	var items []int
	for _, item := range freqNums[:k] {
		items = append(items, item[0])
	}
	return items
}

func QSort(nums [][]int) {
	if len(nums) == 0 {
		return
	}
	selected := nums[0][1]
	var left, right = 0, len(nums) - 1

	for i := 1; i <= right; {
		if nums[i][1] >= selected {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		} else {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		}
	}

	QSort(nums[:left])
	QSort(nums[left+1:])

}

func topKFrequentWithIHeap(nums []int, k int) []int {
	dict := map[int]int{}
	for _, v := range nums {
		dict[v]++
	}

	h := &datastruct.IHeap{}
	heap.Init(h)
	for i, v := range dict {
		heap.Push(h, [2]int{i, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[k-i-1] = heap.Pop(h).([2]int)[0]
	}

	return result
}

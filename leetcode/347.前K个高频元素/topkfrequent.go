package _47_前K个高频元素

import (
	"fmt"
	"math/rand"
	"time"
)

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	values := [][]int{}
	for key, value := range occurrences {
		values = append(values, []int{key, value})
	}
	ret := make([]int, k)
	QSort(values, 0, len(values)-1, ret, 0, k)
	return ret
}

func QSort(nums [][]int, start, end int, result []int, resultIdx, k int) {
	//rand.Seed(time.Now().UnixNano())
	//selectedIdx := rand.Int()%(end-start) + start
	selectedIdx := start

	nums[start], nums[selectedIdx] = nums[selectedIdx], nums[start]

	selected := nums[start][1]
	var index = start
	for i := start + 1; i <= end; i++ {
		if nums[i][1] >= selected {
			nums[index+1], nums[i] = nums[i], nums[index+1]
			index++
		}
	}

	nums[start], nums[index] = nums[index], nums[start]

	if k <= index-start {
		QSort(nums, start, index-1, result, resultIdx, k)
	} else {
		for i := start; i <= index; i++ {
			result[resultIdx] = nums[i][0]
			resultIdx++
		}
		if k > index-start+1 {
			QSort(nums, index+1, end, result, resultIdx, k-(index-start+1))
		}
	}

}

func qsort(values [][]int, start, end int, ret []int, retIndex, k int) {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(end-start+1) + start

	values[picked], values[start] = values[start], values[picked]

	fmt.Printf("values: %+v \n", values)

	pivot := values[start][1]
	index := start

	for i := start + 1; i <= end; i++ {
		if values[i][1] >= pivot {
			values[index+1], values[i] = values[i], values[index+1]
			index++
		}
	}

	values[start], values[index] = values[index], values[start]
	if k <= index-start {
		qsort(values, start, index-1, ret, retIndex, k)
	} else {
		for i := start; i <= index; i++ {
			ret[retIndex] = values[i][0]
			retIndex++
		}
		if k > index-start+1 {
			qsort(values, index+1, end, ret, retIndex, k-(index-start+1))
		}
	}
}

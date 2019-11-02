package _015

import "fmt"

//[
//[-1, 0, 1],
//[-1, -1, 2]
//]
func ThreeSum(nums []int) [][]int {
	var (
		result [][]int
		l      = len(nums)
	)

	fmt.Println(nums)

	var dict = make(map[int][]int)
	for i := 0; i < l-2; i++ {
		for j := i; j < l-1; j++ {
			v, ok := dict[nums[j]]
			if ok {

				result = append(result, []int{nums[j], v[0], v[1]})
				delete(dict, nums[j])
			} else {
				dict[0-nums[j]-nums[i]] = []int{nums[i], nums[j]}
			}
		}
	}

	return result

}

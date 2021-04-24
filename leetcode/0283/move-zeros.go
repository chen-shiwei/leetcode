package _283

import "fmt"

func MoveZeroes(nums []int) {
	var length = len(nums)
	var i, j int
	for j < length {
		if nums[i] == 0 {
			if nums[j] == 0 {
				j++
			} else {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			}
		} else {
			i++
			j++
		}
	}

}

func moveZeroes(nums []int) {
	l := len(nums)
	var slow = 0
	for quick := 0; quick < l; quick++ {
		if nums[quick] != 0 {
			nums[slow] = nums[quick]
			slow++
		}
	}

	for ; slow < l; slow++ {
		nums[slow] = 0
	}
	fmt.Println(nums)
}

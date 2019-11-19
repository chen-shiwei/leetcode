package sort

// https://zh.wikipedia.org/wiki/%E6%8F%92%E5%85%A5%E6%8E%92%E5%BA%8F#Go

// InsertSort 插入排序
func InsertSort(nums []int) []int {

	l := len(nums)

	for i := 0; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				break
			}
		}
	}

	return nums

}

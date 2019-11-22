package _078_todo

func subsets(nums []int) [][]int {

	l := len(nums)

	var result = [][]int{{}, nums[:]}
	if l < 2 {
		return result
	}

	for j := 0; j < l; j++ {

		result = append(result, []int{nums[j]})

		var tmp []int
		tmp = append(tmp, nums[0:j]...)
		tmp = append(tmp, nums[j+1:]...)

		if len(tmp) > 1 {
			result = append(result, tmp)

		}

	}
	return result

}

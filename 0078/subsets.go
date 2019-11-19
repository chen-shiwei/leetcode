package _078

func Subsets(nums []int) [][]int {

	l := len(nums)

	//var dict = make(map[string]bool)

	var result = [][]int{{}, nums[:]}

	if l < 2 {
		return result
	}

	for j := 0; j < l; j++ {

		result = append(result, []int{nums[j]})

		result = append(result, nums[:j])
		result = append(result, nums[j+1:])

		var tmp []int
		if j != 0 {
			tmp = append(tmp, nums[:j]...)
		}
		tmp = append(tmp, nums[j+1:]...)

		if len(tmp) > 1 {
			result = append(result, tmp)
		}

	}
	return result

}

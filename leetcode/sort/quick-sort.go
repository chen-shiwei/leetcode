package sort

func QuickSort(a []int) {
	l := len(a)
	if l < 2 {
		return
	}

	var left, right = 0, l - 1
	var flag = a[0]
	for i := 1; i <= right; {
		if a[i] <= flag {
			a[i], a[left] = a[left], a[i]
			left++
			i++
		} else {
			a[i], a[right] = a[right], a[i]
			right--
		}
	}

	QuickSort(a[:left])
	QuickSort(a[left+1:])

}

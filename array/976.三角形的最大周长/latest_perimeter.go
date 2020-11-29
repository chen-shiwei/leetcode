package _76_三角形的最大周长

func largestPerimeter(A []int) int {

	var l = len(A)

	if l < 3 {
		return 0
	}
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
		if i >= 2 {
			if A[l-1-i]+A[l-1-i+1] > A[l-1-i+2] {
				return A[l-1-i] + A[l-1-i+1] + A[l-1-i+2]
			}
		}
	}

	if A[0]+A[1] > A[2] {
		return A[0] + A[1] + A[2]
	}

	return 0

}

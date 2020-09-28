package _461

//func HammingDistance(x, y int) int {
//	return bits.OnesCount(uint(x) ^ uint(y))
//}

func HammingDistance(x, y int) int {
	var result int
	n := x ^ y

	for n != 0 {
		result++
		n = n & (n - 1)
	}
	return result
}

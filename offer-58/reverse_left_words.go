package offer_58

func reverseLeftWords(s string, n int) string {
	return (s + s)[n : len([]rune(s))+n]
}

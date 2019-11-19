package _524

import "sort"

func FindLongestWord(s string, d []string) string {
	sort.Sort(words(d))

	for _, v := range d {
		if isSub(s, v) {
			return v
		}
	}

	return ""

}

func isSub(s, sub string) bool {
	i, j := 0, 0
	il := len(s)
	jl := len(sub)
	if il < jl {
		return false
	}
	for i < il && j < jl {
		if s[i] == sub[j] {
			j++
		}
		i++
	}
	return j == jl
}

type words []string

// Len is the number of elements in the collection.
func (w words) Len() int {
	return len(w)
}

// Less reports whether the element with
// index i should search before the element with index j.
func (w words) Less(i, j int) bool {
	return len(w[i]) > len(w[j])
}

// Swap swaps the elements with indexes i and j.
func (w words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

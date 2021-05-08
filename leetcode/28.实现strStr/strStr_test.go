package _8_实现strStr

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入：haystack = "hello", needle = "ll"
输出：2`, args: args{
			haystack: "hello",
			needle:   "ll",
		}, want: 2},
		{name: `输入：haystack = "aabaabaafa", needle = "aabaaf"
输出：2`, args: args{
			haystack: "aabaabaafa",
			needle:   "aabaaf",
		}, want: 3},
	}
	for _, tt := range tests {
		//t.Run(tt.name, func(t *testing.T) {
		//	if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
		//		t.Errorf("strStr() = %v, want %v", got, tt.want)
		//	}
		//})
		t.Run(tt.name, func(t *testing.T) {
			if got := strStrWithKMP(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStrWithKMP() = %v, want %v", got, tt.want)
			}
		})
	}

}

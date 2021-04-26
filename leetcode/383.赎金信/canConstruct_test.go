package _83_赎金信

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: `输入：ransomNote = "a", magazine = "b"
输出：false`, args: args{
			ransomNote: "a",
			magazine:   "b",
		}},
		{name: `输入：ransomNote = "aa", magazine = "aab"
输出：true`, args: args{
			ransomNote: "aa",
			magazine:   "aab",
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

package 面试题08_01_三步问题

import "testing"

func Test_waysToStep(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `3`, args: args{n: 3}, want: 4},
		{name: `1`, args: args{n: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToStep1(tt.args.n); got != tt.want {
				t.Errorf("waysToStep1() = %v, want %v", got, tt.want)
			}
		})
	}
}

package 面试题_08_11

import "testing"

func Test_waysToChange(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `示例1:
				 输入: n = 5
				 输出：2
				 解释: 有两种方式可以凑成总金额:
				5=5
				5=1+1+1+1+1`, args: args{n: 5}, want: 2},
		{name: `示例2:
				 输入: n = 10
				 输出：4
				 解释: 有四种方式可以凑成总金额:
				10=10
				10=5+5
				10=5+1+1+1+1+1
				10=1+1+1+1+1+1+1+1+1+1`, args: args{n: 10}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToChange(tt.args.n); got != tt.want {
				t.Errorf("waysToChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

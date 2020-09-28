package _111

import (
	"reflect"
	"testing"
)

func TestMaxDepthAfterSplit(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: `输入：seq = "(()())"
输出：[0,1,1,1,1,0]`, args: args{seq: "(()())"}, want: []int{0, 1, 1, 1, 1, 0}},
		{name: `输入：seq = "()(())()"
输出：[0,0,0,1,1,0,1,1]`, args: args{seq: "()(())()"}, want: []int{0, 0, 0, 1, 1, 0, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDepthAfterSplit(tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxDepthAfterSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

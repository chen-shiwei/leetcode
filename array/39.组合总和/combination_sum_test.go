package _9_组合总和

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: `输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]`, args: args{
			candidates: []int{2, 3, 6, 7},
			target:     7,
		}, want: [][]int{
			{7},
			{2, 2, 3},
		}},
		{name: `[2,7,6,3,5,1]
9`, args: args{
			candidates: []int{2, 7, 6, 3, 5, 1},
			target:     9,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

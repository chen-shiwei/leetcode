package _739

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	type args struct {
		T []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: `temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]`, args: args{T: []int{73, 74, 75, 71, 69, 72, 76, 73}}, want: []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DailyTemperatures(tt.args.T); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

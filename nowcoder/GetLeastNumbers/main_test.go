package GetLeastNumbers

import (
	"reflect"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	type args struct {
		input []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "22, 34, 3, 40, 18, 4", args: args{
			input: []int{22, 34, 3, 40, 18, 4},
			k:     3,
		}, want: []int{3, 4, 18}},
		{name: "", args: args{
			input: []int{1, 2},
			k:     2,
		}, want: []int{1, 2}},
		{name: "4,5,1,6,2,7,3,8=>1,2,3,4", args: args{
			input: []int{4, 5, 1, 6, 2, 7, 3, 8},
			k:     4,
		}, want: []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLeastNumbers(tt.args.input, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Log(tt.args.input)
				t.Errorf("GetLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

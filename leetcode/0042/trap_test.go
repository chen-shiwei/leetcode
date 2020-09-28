package _042

import "testing"

func TestTrap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `输入: [0,1,0,2,1,0,1,3,2,1,2,1]
				输出: 6`,
			args: args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap(tt.args.height); got != tt.want {
				t.Errorf("Trap() = %v, want %v", got, tt.want)
			}
		})
	}
}

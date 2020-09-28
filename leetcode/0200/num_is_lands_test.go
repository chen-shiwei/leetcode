package _200

import "testing"

func Test_numIsLands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入:
				11110
				11010
				11000
				00000
				输出: 1
				`, args: args{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
		}, want: 1},
		{name: `输入:
				11000
				11000
				00100
				00011
				输出: 3`, args: args{grid: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}}, want: 3},
		{name: `
				{1,1,1,0},
				{1,1,0,0},
				{1,1,0,0},
				{0,0,0,0}`,
			args: args{[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIsLands(tt.args.grid)
			t.Logf("got:%d", got)
			if got != tt.want {
				t.Errorf("numIsLands() = %v, want %v", got, tt.want)
			}
			t.Logf("%+v", tt.args.grid)

		})
	}
}

func TestSlice(t *testing.T) {
	var s1 = []int{1, 2, 3, 4}

	var s2 = make([]int, len(s1), cap(s1))

	copy(s2, s1)

	s1[0] = 9

	t.Logf("s1:%+v s2:%+v", s1, s2)

}

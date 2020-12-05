package _21_任务调度器

import "testing"

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `输入：tasks = ["A","A","A","B","B","B"], n = 2
输出：8`, args: args{
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     2,
		}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval1(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval1() = %v, want %v", got, tt.want)
			}
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

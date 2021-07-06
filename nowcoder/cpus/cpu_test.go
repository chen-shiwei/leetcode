package cpus

import "testing"

// 小米面试 2021年07月06日23:15:45
// // m[] cpu, m[i] =
//// task[]  ->
//
//func Schedule(m []int, task[] int) {
//}
//// [1, 2]
//// [1, 1, 1, 1]
//-> 1 -> 1 -> 1
//-> 2 -> 1 -> 2
//-> 3 -> 2 -> 1
//-> 4 -> 2 -> 2
//
//// [1,2]
////[2, 1, 1, 1]
//-> task 1 -> 1s -> 1 cpu
//-> task 2 -> 1s -> 2 cpu
//-> task 3 -> 2 s -> 2 cpu
//-> task 4 -> 3s -> 1 cpu
func TestSchedule(t *testing.T) {
	type args struct {
		m    []int
		task []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: `// [1, 2] // [1, 1, 1, 1]`, args: args{
			m:    []int{1, 2},
			task: []int{1, 1, 1, 1},
		}},
		{name: `// [1,2] //[2, 1, 1, 1]`, args: args{
			m:    []int{1, 2},
			task: []int{2, 1, 1, 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Schedule(tt.args.m, tt.args.task)

		})
	}
}

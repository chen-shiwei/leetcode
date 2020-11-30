package 反转链表

import (
	"reflect"
	"testing"
)

func Test_reverseLinklist(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: `1->2->3->4->5 => 54321`, args: args{head: &Node{
			Val: 1,
			Next: &Node{
				Val: 2,
				Next: &Node{
					Val: 3,
					Next: &Node{
						Val: 4,
						Next: &Node{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLinklist(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseLinklist() = %v, want %v", got, tt.want)
			}
		})
	}
}

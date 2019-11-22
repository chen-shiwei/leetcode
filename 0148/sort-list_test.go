package _148

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

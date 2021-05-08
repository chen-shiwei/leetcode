package datastruct

import (
	"reflect"
	"testing"
)

func TestDeque(t *testing.T) {
	dq := NewDeque()
	dq.PushBack(1)
	if !reflect.DeepEqual(dq.data, []int{1}) {
		t.Error()
	}
	dq.PushBack(2)
	if !reflect.DeepEqual(dq.data, []int{1, 2}) {
		t.Error()
	}
	dq.PushFront(3)
	if !reflect.DeepEqual(dq.data, []int{3, 1, 2}) {
		t.Error()
	}
	if !reflect.DeepEqual(dq.PopFront(), 3) {
		t.Error()
	}
	if !reflect.DeepEqual(dq.PopBack(), 2) {
		t.Error()
	}
	if !reflect.DeepEqual(dq.PopFront(), 1) {
		t.Error()
	}
	if !reflect.DeepEqual(dq.PopFront(), -1) {
		t.Error()
	}
}

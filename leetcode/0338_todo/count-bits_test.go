package _338_todo

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {

	var n int64 = 2
	fmt.Println(DecBin(n))
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBits(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func DecBin(n int64) string {
	if n < 0 {
		log.Println("Decimal to binary error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 2 {
		m := q % 2
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

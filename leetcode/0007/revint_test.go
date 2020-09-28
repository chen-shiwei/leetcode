package _007

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "321", args: args{x: 321}, want: 123},
		{name: "123", args: args{x: 123}, want: 321},
		{name: "-123", args: args{x: -123}, want: -321},
		{name: "120", args: args{x: 120}, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.x); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

package _022

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

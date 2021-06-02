package _1_N皇后_回溯

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	type args struct {
		chessboard [][]byte
		row        int
		col        int
		n          int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{
			chessboard: [][]byte{
				{'.', '.', '.', '.'},
				{'.', 'Q', '.', '.'},
				{'Q', '.', '.', '.'},
				{'Q', '.', '.', '.'},
			},
			row: 2,
			col: 2,
			n:   4,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.chessboard, tt.args.row, tt.args.col, tt.args.n); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: `输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]`, args: args{n: 4}, want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

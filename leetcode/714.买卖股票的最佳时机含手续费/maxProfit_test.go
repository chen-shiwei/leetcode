package _714_买卖股票的最佳时机含手续费

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `prices = [1, 3, 2, 8, 4, 9], fee = 2 输出: 8`, args: args{
			prices: []int{1, 3, 2, 8, 4, 9},
			fee:    2,
		}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

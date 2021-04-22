package 爬楼梯

// 一个人爬楼梯，每次只能爬 1 个或 2 个台阶，假设有 n 个台阶，那么这个人有多少种不同的爬楼梯方法？
//思路：
//由于第 n 级台阶一定是从 n - 1 级台阶或者 n - 2 级台阶来的，因此到第 n 级台阶的数目就是 到第 n - 1 级台阶的数目加上到第 n - 1 级台阶的数目。

// 普通递归
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 记忆化递归
func climbStairs1(n int) int {
	var record = make(map[int]int)
	var fn func(n int) int
	fn = func(n int) int {
		if v, ok := record[n]; ok {
			return v
		}
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}
		v := fn(n-1) + fn(n-2)
		record[n] = v
		return v
	}
	result := fn(n)
	return result
}

func climbStairs2(n int) int {

	pre := 1
	cur := 2
	for i := 3; i <= n; i++ {
		pre, cur = cur, pre+cur
	}

	return cur

}

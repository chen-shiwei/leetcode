package 面试题08_01_三步问题

func waysToStep(n int) int {
	var memory = make(map[int]int)
	var fn func(n int) int
	fn = func(n int) int {
		if v, ok := memory[n]; ok {
			return v
		}
		if n < 3 {
			return n
		}
		if n == 3 {
			return 4
		}
		v := (fn(n-1) + fn(n-2) + fn(n-3)) % 1000000007
		memory[n] = v
		return v
	}
	result := fn(n)
	return result
}

func waysToStep1(n int) int {
	if n < 3 {
		return n
	}
	var one, two, three = 1, 2, 4
	for i := 4; i <= n; i++ {
		one, two, three = two, three, (((one+two)%1000000007)+three)%1000000007
	}
	return three
}

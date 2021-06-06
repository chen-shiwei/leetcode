package _134_加油站

func canCompleteCircuit(gas []int, cost []int) int {
	l1 := len(gas)

	for i := 0; i < l1; i++ {
		var ans = gas[i] - cost[i]
		idx := (i + 1) % l1
		for idx != i && ans > 0 {
			ans += gas[idx] - cost[idx]
			idx = (idx + 1) % l1
		}
		if ans >= 0 && idx == i {
			return i
		}
	}
	return -1
}

func canCompleteCircuitWithGreedy(gas []int, cost []int) int {
	l := len(gas)
	var curSum, sum, idx int
	for i := 0; i < l; i++ {
		sum += gas[i] - cost[i]
		curSum += gas[i] - cost[i]
		if curSum < 0 {
			idx = i + 1
			curSum = 0
		}
	}
	if sum < 0 {
		return -1
	}

	return idx
}

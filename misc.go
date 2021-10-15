package main

func Solution(A []int) int {
	ret := 0
	memo := make(map[int]int)
	for _, v := range A {
		memo[v]++
	}
	for k := range memo {
		if _, ok := memo[-k]; ok {
			return abs(k)
		}
	}
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Solution2(S string, C []int) int {
	mincost := 0
	runes := []rune(S)
	// We have to consider each adjacent position once -> O(n)
	for i := len(runes) - 1; i > 0; i-- {
		if runes[i] == runes[i-1] {
			mincost += min(C[i], C[i-1])
		}
	}
	// We stored the running total, just return it
	return mincost
}

func Solution3(a1, a2 []int) int {
	return 0
}

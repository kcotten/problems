package main

func knapsack(profits, weights []int, capacity, idx int) int {
	memo := make([][]int, len(profits))
	for i := range memo {
		memo[i] = make([]int, capacity+1)
	}
	var helper func([]int, []int, int, int) int
	helper = func(profits, weights []int, capacity, idx int) int {
		if capacity <= 0 || idx >= len(profits) {
			return 0
		}
		if memo[idx][capacity] != 0 {
			return memo[idx][capacity]
		}
		p1 := 0
		if weights[idx] <= capacity {
			p1 = profits[idx] + helper(profits, weights, capacity-weights[idx], idx+1)
		}
		p2 := helper(profits, weights, capacity, idx+1)
		memo[idx][capacity] = max(p1, p2)
		return memo[idx][capacity]
	}
	ret := helper(profits, weights, capacity, idx)
	return ret
}

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := wordSet[s[j:i]]; ok {
				if dp[j] {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)]
}

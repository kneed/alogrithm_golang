package main

func coinChange(coins []int, amount int) int {

	var dp = make([]int, amount+1, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 0; i < amount+1; i++ {
		for _, j := range coins {
			if i-j < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-j])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

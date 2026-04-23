func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := slices.Repeat([]int{amount + 1}, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
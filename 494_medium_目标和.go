package main

//给你一个整数数组 nums 和一个整数 target 。
//
//向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
//例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
//返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
//
//
//
//示例 1：
//
//输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3
//示例 2：
//
//输入：nums = [1], target = 1
//输出：1

func findTargetSumWays1(nums []int, target int) int {
	result := 0

	var backTrack func(nums []int, i int, rest int)

	backTrack = func(nums []int, i int, rest int) {
		if i == len(nums) {
			if rest == 0 {
				result += 1
			}
			return
		}

		rest += nums[i]
		backTrack(nums, i+1, rest)
		rest -= nums[i]

		rest -= nums[i]
		backTrack(nums, i+1, rest)
		rest += nums[i]
	}
	backTrack(nums, 0, target)
	return result

}

// dp的解法
func findTargetSumWays(nums []int, target int) int {
	var dp = make([]map[int]int, len(nums)+1)
	for i := range dp {
		dp[i] = map[int]int{}
	}
	dp[0][0] = 1
	for i, num := range nums {
		for k, v := range dp[i] {
			dp[i+1][k+num] += v
			dp[i+1][k-num] += v
		}
	}
	return dp[len(nums)][target]
}

package main

import "fmt"

//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//子数组 是数组中的一个连续部分。
//
//
//
//示例 1：
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//示例 2：
//
//输入：nums = [1]
//输出：1
//示例 3：
//
//输入：nums = [5,4,-1,7,8]
//输出：23

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var dp = make([]int, 2)
	dp[0] = nums[0]
	var maxResult = dp[0]
	for i := 1; i < len(nums); i++ {
		dp[1] = max(nums[i], dp[0]+nums[i])
		dp[0] = dp[1]
		maxResult = max(maxResult, dp[1])
	}
	return maxResult
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

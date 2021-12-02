package main

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//示例 2：
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//示例 3：
//
//输入：nums = [1]
//输出：[[1]]
var result [][]int

func permute(nums []int) [][]int {
	var track []int
	backtrace(nums, track)
	r := result
	result = [][]int{}
	return r
}

func backtrace(nums, track []int) {
	// 返回条件
	if len(nums) == len(track) {
		var temp = make([]int, len(track))
		copy(temp, track)
		result = append(result, temp)
		return
	}

	// 循环遍历
	for i := 0; i < len(nums); i++ {
		// 已经遍历过的数字过滤
		var isContain bool
		for _, x := range track {
			if nums[i] == x {
				isContain = true
			}
		}
		if isContain {
			continue
		}
		track = append(track, nums[i])
		backtrace(nums, track)
		track = track[:len(track)-1]

	}
}

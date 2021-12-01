package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxInIntList(arr []int) int {

	//假设第一个元素是最大值，下标为0
	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		//从第二个 元素开始循环比较，如果发现有更大的，则交换
		if maxVal < arr[i] {
			maxVal = arr[i]

		}
	}
	return maxVal
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

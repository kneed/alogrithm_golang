package main

import "math"

//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//
//路径和 是路径中各节点值的总和。
//
//给你一个二叉树的根节点 root ，返回其 最大路径和 。
//

var maxSum = math.MinInt32

func maxPathSum(root *TreeNode) int {
	var maxSum = math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 左节点的最大值.
		left := max(0, maxGain(node.Left))
		// 右节点的最大值.
		right := max(0, maxGain(node.Right))
		//本节点的最大路径和
		newPathValue := node.Val + left + right

		// 更新最大路径和
		maxSum = max(maxSum, newPathValue)

		// 返回当前节点的最大值
		return node.Val + max(left, right)
	}
	maxGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetCode显示的最优解
//type ResultType struct {
//	SinglePath int
//	MaxPath    int
//}
//
//func maxPathSum(root *TreeNode) int {
//	result := helper(root)
//	return result.MaxPath
//}
//
//func helper(root *TreeNode) ResultType {
//	if root == nil {
//		return ResultType{
//			SinglePath: 0,
//			MaxPath:    -(1 << 31),
//		}
//	}
//
//	left := helper(root.Left)
//	right := helper(root.Right)
//
//	result := ResultType{}
//	if left.SinglePath > right.SinglePath {
//		result.SinglePath = max(left.SinglePath+root.Val, 0)
//	} else {
//		result.SinglePath = max(right.SinglePath+root.Val, 0)
//	}
//
//	maxPath := max(right.MaxPath, left.MaxPath)
//	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
//	return result
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}

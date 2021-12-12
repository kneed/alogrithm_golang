package main

//给定一个二叉树，找出其最小深度。
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//说明：叶子节点是指没有子节点的节点。

// BFS解法

func minDepth(root *TreeNode) int {
	// 特殊边界判断
	if root == nil {
		return 0
	}
	// BFS需要一个队列来存储下次遍历的节点
	var queue []*TreeNode
	// 根节点入队
	queue = append(queue, root)
	depth := 1

	// 遍历queue
	for len(queue) != 0 {
		qSize := len(queue)

		for i := 0; i < qSize; i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				return depth
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[qSize:]
		// 每一次等于遍历一层, 如果一层都没有返回则深度+1
		depth += 1
	}
	return depth
}

// 这个写法里的queue需要多次截取分配不太划算, 实际上可以通过一个移动下标去减少queue的重新赋值导致的新分配内存

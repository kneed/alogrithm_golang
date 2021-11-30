package main

//给定一棵树的前序遍历 preorder 与中序遍历  inorder。请构造二叉树并返回其根节点。

//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 前序遍历的第一个节点即树的根节点
	root := &TreeNode{preorder[0], nil, nil}

	// 找到前序遍历根节点在中序遍历的位置
	i := 0
	for _, val := range inorder {
		if val == preorder[0] {
			break
		}
		i++
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root

}

//LeetCode的最优答案.没看懂跟我的区别在哪里

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder) == 0{
//		return nil
//	}
//	root := &TreeNode{preorder[0], nil, nil}
//	index := 0
//	for _, val := range inorder{
//		if val == preorder[0]{
//			break
//		}
//		index += 1
//	}
//	root.Left = buildTree(preorder[1: len(inorder[:index])+1], inorder[:index])
//	root.Right = buildTree(preorder[(len(inorder[:index])+1):], inorder[index+1:])
//
//	return root
//}

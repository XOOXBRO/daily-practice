package main

/*
给你一棵 完整二叉树 的根，这棵树有以下特征：

叶子节点 要么值为 0 要么值为 1 ，其中 0 表示 False ，1 表示 True 。
非叶子节点 要么值为 2 要么值为 3 ，其中 2 表示逻辑或 OR ，3 表示逻辑与 AND 。
计算 一个节点的值方式如下：

如果节点是个叶子节点，那么节点的 值 为它本身，即 True 或者 False 。
否则，计算 两个孩子的节点值，然后将该节点的运算符对两个孩子值进行 运算 。
返回根节点 root 的布尔运算值。

完整二叉树 是每个节点有 0 个或者 2 个孩子的二叉树。

叶子节点 是没有孩子的节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// See https://leetcode.cn/problems/evaluate-boolean-binary-tree/description/ for more details
func evaluateTree(root *TreeNode) bool {
	if root.Left == nil {
		if root.Val == 0 {
			return false
		}
		return true
	}
	l := evaluateTree(root.Left)
	r := evaluateTree(root.Right)
	if !l && !r {
		return false
	}
	if root.Val == 3 && (!l || !r) {
		return false
	}
	return true
}

// authority's method
func evaluateTree2(root *TreeNode) bool {
	switch root.Val {
	case 0:
		return false
	case 1:
		return true
	case 2:
		return evaluateTree2(root.Left) || evaluateTree2(root.Right)
	case 3:
		return evaluateTree2(root.Left) && evaluateTree2(root.Right)
	default:
		panic("unreachable")
	}
}

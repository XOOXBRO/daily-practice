package main

/*
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。

See https://leetcode.cn/problems/binary-tree-postorder-traversal/description/ for more details
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
/*
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	l := postorderTraversal(root.Left)
	r := postorderTraversal(root.Right)

	res = append(res, l...)
	res = append(res, r...)
	res = append(res, root.Val)

	return res
}


*/

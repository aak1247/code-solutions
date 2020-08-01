/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	flattenWithRight(root, nil)
}

func flattenWithRight(root, right *TreeNode) {
	if root == nil {
		return 
	}
	if root.Right != nil {
		flattenWithRight(root.Right, right)
	} else {
		root.Right = right
	}
	if root.Left != nil {
		flattenWithRight(root.Left, root.Right)
		root.Right = root.Left
		root.Left = nil
	}
}
// @lc code=end


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    return maxOf(root, 0)
}

func maxOf(root *TreeNode, l int) int {
    if root == nil {
        return l
    } else {
        return max(maxOf(root.Left, l+1), maxOf(root.Right, l+1))
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
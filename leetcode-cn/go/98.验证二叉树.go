/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
func isValidBST(root *TreeNode) bool {
    return root == nil || isValidBSTWithMaxOrMin(root, INT_MIN, INT_MAX)
}

func maxOrMin(num1, num2 int, isMaxOrMin bool) int {
    if isMaxOrMin && num1 > num2 || !isMaxOrMin && num2 > num1 {
        return num1
    }
    return num2
} 

func isValidBSTWithMaxOrMin(root *TreeNode, min, max int) bool {
    val := root.Val
    if root.Left != nil {
        if !isValidBSTWithMaxOrMin(root.Left, min, maxOrMin(val, max, false)) {
            return false
        }
        if root.Left.Val <= min || root.Left.Val >= maxOrMin(val, max, false) {
            return false
        }
    }

    if root.Right != nil {
        if !isValidBSTWithMaxOrMin(root.Right, maxOrMin(val, min, true), max) {
            return false
        }
        if root.Right.Val <= maxOrMin(val, min, true) || root.Right.Val >= max {
            return false
        }
    }
    return true
}
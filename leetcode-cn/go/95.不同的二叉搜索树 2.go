/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }
    return generateTreesBetween(0,n)
}

//l开
func generateTreesBetween(l, r int) []*TreeNode {
    if l == r {
        return []*TreeNode{nil}
    }
    if r - l == 1 {
        cur := new(TreeNode)
        cur.Val = r
        return []*TreeNode{cur}
    }
    if r - l == 2 {
        l1, l2, r1, r2  := new(TreeNode), new(TreeNode), new(TreeNode), new(TreeNode)
        l1.Val, l2.Val, r1.Val, r2.Val = l+1,l+1,r,r
        l1.Right = r1
        r2.Left = l2
        return []*TreeNode{l1,r2}
    }
    var res = make([]*TreeNode, 0)
    for i := l+1; i<=r;i++ {
        leftTrees := generateTreesBetween(l,i-1)
        rightTrees := generateTreesBetween(i,r)
        for _, leftTree := range leftTrees {
            for _, rightTree := range rightTrees {
                root := new(TreeNode)
                root.Val = i
                root.Left = leftTree
                root.Right = rightTree
                res = append(res, root)
            }
        }
    }
    return res
}
package container

// binary tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewBinaryTree returns a new TreeNode
// lever order
// if src[i] == -999 -> node == nil
func NewBinaryTree(src []int) *TreeNode {
	if len(src) == 0 {
		return &TreeNode{}
	}

	root := &TreeNode{Val: src[0]}
	var pre, next []*TreeNode
	var node *TreeNode
	pre = append(pre, root)
	var isLeft = true

	i, j := 0, 1
	for len(pre) != 0 {
		for ; i < len(pre) && j < len(src); j++ {
			node = pre[i]
			if isLeft {
				node.Left = &TreeNode{Val: src[j]}
				next = append(next, node.Left)
				isLeft = false
			} else {
				node.Right = &TreeNode{Val: src[j]}
				next = append(next, node.Right)
				isLeft = true
				i++
			}
		}

		i = 0
		pre = next
		next = make([]*TreeNode, 0)
	}

	return root
}

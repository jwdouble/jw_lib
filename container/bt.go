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

	var node *TreeNode
	var pre, next []*TreeNode
	root := &TreeNode{Val: src[0]}
	pre = append(pre, root)

	i, j := 0, 1
	var isLeft = true
	for len(pre) != 0 {
		for ; i < len(pre) && j < len(src); j++ {
			node = pre[i]
			if isLeft {
				if src[j] != -999 {
					node.Left = &TreeNode{Val: src[j]}
					next = append(next, node.Left)
				}

				isLeft = false
			} else {
				if src[j] != -999 {
					node.Right = &TreeNode{Val: src[j]}
					next = append(next, node.Right)
				}
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

package container

import (
	"log"
)

// link list

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkList(src []int) *ListNode {
	if len(src) == 0 {
		return nil
	}
	root := &ListNode{}
	node := root
	for _, v := range src {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}

	return root.Next
}

func (root *ListNode) List() {
	node := root
	for node != nil {
		log.Println(node.Val)
		node = node.Next
	}
	log.Println("finish")
}

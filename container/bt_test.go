package container

import (
	"fmt"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := NewBinaryTree(src)
	fmt.Println(root)
}

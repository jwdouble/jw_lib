package container

import (
	"testing"
)

func Test_linkList(t *testing.T) {
	l := NewLinkList([]int{1, 2, 3, 4, 5})
	l.List()
}

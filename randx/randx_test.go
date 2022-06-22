package randx

import (
	"fmt"
	"testing"
)

func TestNewString(t *testing.T) {
	fmt.Printf("-> %s \r\n", NewString(32))
}

func TestNewInt(t *testing.T) {
	fmt.Printf("-> %d \r\n", NewInt(32))
}

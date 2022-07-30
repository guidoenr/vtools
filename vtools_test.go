package vtools

import (
	"testing"
)

func TestPrint(t *testing.T) {
	list := []string{"Hello", "World"}

	Print("list", list)
}

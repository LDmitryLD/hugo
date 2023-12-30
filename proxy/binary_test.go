package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMer = "10 --> 0\n10 --> 30\n30 --> 20\n30 --> 40\n"

func Test_Binary(t *testing.T) {
	tree := GenerateTree(5)
	mer := tree.ToMermaid()
	assert.True(t, isBalanced(tree.Root))
	assert.Equal(t, testMer, mer)
}

func isBalanced(node *NodeB) bool {
	if node == nil {
		return true
	}

	balanceF := getBalance(node)
	if balanceF < -1 || balanceF > 1 {
		return false
	}

	return isBalanced(node.Left) && isBalanced(node.Right)
}

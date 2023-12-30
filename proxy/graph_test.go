package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testGraph = []*Node{
		{ID: 0, Name: "A", Form: "Circle"},
		{ID: 1, Name: "B", Form: "Circle", Links: []*Node{}},
	}

	expectMer = "A[Circle] --> B[Circle]\n"
)

func Test_Graph(t *testing.T) {
	testGraph[0].Links = append(testGraph[0].Links, testGraph[1])
	mer := grapToMermaid(testGraph)
	assert.Equal(t, expectMer, mer)
	assert.True(t, exist(testGraph, testGraph[0]))
	assert.False(t, exist(testGraph, &Node{ID: 4, Name: "D"}))
}

package mermaid

import (
	"fmt"
	"math/rand"
	"time"
)

type NodeB struct {
	Key    int
	Height int
	Left   *NodeB
	Right  *NodeB
}

type AVLTree struct {
	Root *NodeB
}

func NewNode(key int) *NodeB {
	return &NodeB{Key: key, Height: 1}
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func (t *AVLTree) ToMermaid() string {
	return recursiveFunc(t.Root)
}

func recursiveFunc(node *NodeB) string {
	if node == nil {
		return ""
	}

	var left, right string

	if node.Left != nil {
		left = fmt.Sprintf("%d --> %d\n", node.Key, node.Left.Key)
		left += recursiveFunc(node.Left)
	}

	if node.Right != nil {
		right = fmt.Sprintf("%d --> %d\n", node.Key, node.Right.Key)
		right += recursiveFunc(node.Right)
	}

	return left + right
}

func height(node *NodeB) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func updateHeight(node *NodeB) {
	node.Height = max(height(node.Left), height(node.Right)) + 1
}

func getBalance(node *NodeB) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func leftRotate(x *NodeB) *NodeB {
	right := x.Right
	x.Right = right.Left
	right.Left = x
	updateHeight(x)
	updateHeight(right)
	return right
}

func rightRotate(y *NodeB) *NodeB {
	left := y.Left
	y.Left = left.Right
	left.Right = y
	updateHeight(y)
	updateHeight(left)
	return left
}

func insert(node *NodeB, key int) *NodeB {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	}

	return balance(node)
}

func balance(node *NodeB) *NodeB {
	updateHeight(node)
	balanceF := getBalance(node)
	if balanceF > 1 {
		if getBalance(node.Left) < 0 {
			node.Left = leftRotate(node.Left)
		}
		return rightRotate(node)
	}

	if balanceF < -1 {
		if getBalance(node.Right) > 0 {
			node.Right = rightRotate(node.Right)
		}
		return leftRotate(node)
	}

	return node
}

func GenerateTree(count int) *AVLTree {
	AVL := AVLTree{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {

		AVL.Insert(i * 10)
	}

	return &AVL
}

const BinaryContent = `---
menu:
    after:
        name: binary_tree
        weight: 2
title: Построение сбалансированного бинарного дерева
---

# Задача построить сбалансированное бинарное дерево
Используя AVL дерево, постройте сбалансированное бинарное дерево, на текущей странице.

Нужно написать воркер, который стартует дерево с 5 элементов, и каждые 5 секунд добавляет новый элемент в дерево.

Каждые 5 секунд на странице появляется актуальная версия, сбалансированного дерева.

При вставке нового элемента, в дерево, нужно перестраивать дерево, чтобы оно оставалось сбалансированным.

Как только дерево достигнет 100 элементов, генерируется новое дерево с 5 элементами.

` + "```" + `go
package binary

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func NewNode(key int) *Node {
	return &Node{Key: key, Height: 1}
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func (t *AVLTree) ToMermaid() string {

}

func height(node *Node) int {

}

func max(a, b int) int {

}

func updateHeight(node *Node) {

}

func getBalance(node *Node) int {

}

func leftRotate(x *Node) *Node {

}

func rightRotate(y *Node) *Node {

}

func insert(node *Node, key int) *Node {

}

func GenerateTree(count int) *AVLTree {

}
` + "```" + `

Не обязательно использовать выше описанный код, можно использовать любую реализацию, выдающую сбалансированное бинарное дерево.

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Пример вывода

{{< columns >}}
` + "```" + `tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph TD
%s
{{</*/* /mermaid */*/>}}
` + "```" + `

{{< /columns >}}

{{< mermaid >}}
graph TD
%s

{{< /mermaid >}}`

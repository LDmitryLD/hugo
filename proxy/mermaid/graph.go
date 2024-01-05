package mermaid

import (
	"fmt"
	"math/rand"
	"time"
)

const GraphConten = `---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

Нужно написать воркер, который будет строить граф на текущей странице, каждые 5 секунд
От 5 до 30 элементов, случайным образом. Все ноды графа должны быть связаны.

` + "```" + `go
type Node struct {
    ID int
    Name string
	Form string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
    Links []*Node
}
` + "```" + `


## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Пример

{{< columns >}}
` + "```" + `tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph LR
%s
{{</*/* /mermaid */*/>}}
` + "```" + `

<--->

{{< mermaid >}}
graph LR
%s
{{< /mermaid >}}

{{< /columns >}}
`

type Node struct {
	ID    int
	Name  string
	Form  string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
	Links []*Node
}

var forms = []string{"circle", "rect", "square", "ellipse", "round-rect", "rhombus"}

func MakeGraph() string {

	rand.Seed(time.Now().UnixNano())

	nodesNum := rand.Intn((30-5)+1) + 5
	nodes := make([]*Node, 0, nodesNum)

	for i := 0; i < nodesNum; i++ {
		nodes = append(nodes, &Node{
			ID:   i,
			Name: string(rune('A' + i)),
			Form: forms[rand.Intn(6)],
		})
	}

	for i, node := range nodes {
		numLinks := rand.Intn(nodesNum - i)

		if numLinks > 3 {
			numLinks = 3
		}

		linkedNodes := make([]*Node, 0, numLinks)

		for len(linkedNodes) < numLinks {
			idx := rand.Intn(nodesNum)

			if idx != i && !exist(linkedNodes, nodes[idx]) {
				linkedNodes = append(linkedNodes, nodes[idx])
			}
		}

		node.Links = linkedNodes
	}

	return grapToMermaid(nodes)
}

func grapToMermaid(nodes []*Node) string {
	var res string
	for _, node := range nodes {
		for _, link := range node.Links {
			res += fmt.Sprintf("%s[%s] --> %s[%s]\n", node.Name, node.Form, link.Name, link.Form)
		}
	}
	return res
}

func exist(nodes []*Node, N *Node) bool {
	for _, node := range nodes {
		if node == N {
			return true
		}
	}
	return false
}

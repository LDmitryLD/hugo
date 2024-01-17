package mermaid

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func WorkerTest() {
	t := time.NewTicker(5 * time.Second)
	rand.Seed(time.Now().UnixNano())
	var tree *AVLTree
	var treeCountre int
	var b byte = 0
	for {
		select {
		case <-t.C:
			graph := MakeGraph()

			if treeCountre == 0 || treeCountre == 100 {
				tree = GenerateTree(5)
				treeCountre = 5
			} else {
				tree.Insert(rand.Intn(150))
				treeCountre++
			}

			mer := tree.ToMermaid()
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, time.Now().Format("2006-01-02 15:04:05"), b)), 0644)
			if err != nil {
				log.Println("ошибка при записи файла:", err)
			}
			b++

			err = os.WriteFile("/app/static/tasks/graph.md", []byte(fmt.Sprintf(GraphConten, graph, graph)), 0644)
			if err != nil {
				log.Println("ошибка при записи файла:", err)
			}

			err = os.WriteFile("/app/static/tasks/binary.md", []byte(fmt.Sprintf(BinaryContent, mer, mer)), 0644)
			if err != nil {
				log.Println("ошибка при записи файла:", err)
			}
		default:
			_ = 4
		}
	}
}

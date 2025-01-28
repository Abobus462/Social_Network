package main

import (
	"fmt"
	"wintersc/graph"
)

func main() {
	graph1 := graph.NewGraph()

	graph1.AddEdge(1, 2, 2)
	graph1.AddEdge(2, 3, 4)
	graph1.AddEdge(3, 4, 4)
	graph1.AddEdge(10, 11, 4)

	fmt.Println(graph.ConnectedComponents(graph1))
}

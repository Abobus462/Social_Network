package main

import (
	"fmt"
	"wintersc/graph"
)

func main() {
	graph1 := graph.NewGraph()

	graph1.AddEdge(0, 1, 10)
	graph1.AddEdge(0, 2, 4)
	graph1.AddEdge(2, 3, 8)
	graph1.AddEdge(3, 4, 7)
	graph1.AddEdge(4, 5, 6)
	graph1.AddEdge(0, 5, 12)
	graph1.AddEdge(4, 1, 20)

	fmt.Println(graph.BoruvkaMST(len(graph1.Adj), graph1.GetAllEdges()))

}

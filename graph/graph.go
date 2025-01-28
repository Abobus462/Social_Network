package graph

type Graph struct {
	Adj  map[int][]int
	Edge []Edge
}

type Edge struct {
	U, V, W int
}

func NewGraph() *Graph {
	return &Graph{
		Adj:  make(map[int][]int),
		Edge: []Edge{},
	}
}

func (g *Graph) AddEdge(u, v, w int) {
	if _, exists := g.Adj[u]; !exists {
		g.Adj[u] = []int{}
	}
	if _, exists := g.Adj[v]; !exists {
		g.Adj[v] = []int{}
	}
	g.Adj[u] = append(g.Adj[u], v)
	g.Adj[v] = append(g.Adj[v], u)

	g.Edge = append(g.Edge, Edge{U: u, V: v, W: w})
}

func HasEdge(g *Graph, u, v int) bool {
	if neighbors, exists := g.Adj[u]; exists {
		for _, neighbor := range neighbors {
			if neighbor == v {
				return true
			}
		}
	}
	return false
}

func ConnectedComponents(g *Graph) (count int, comp map[int]int) {
	visited := make(map[int]bool)
	comp = make(map[int]int)
	count = 0

	for key := range g.Adj {
		if !visited[key] {
			count++
			component := DFS(g, key)
			for _, value := range component {
				visited[value] = true
				comp[value] = count
			}
		}
	}

	return count, comp
}

// ?????????????????????????????????????????????????????????
func (g *Graph) GetAllEdges() []Edge {
	var edges []Edge
	for u, neighbors := range g.Adj {
		for _, v := range neighbors {
			if u < v {
				var weight int
				for _, edge := range g.Edge {
					if (edge.U == u && edge.V == v) || (edge.U == v && edge.V == u) {
						weight = edge.W
						break
					}
				}
				edges = append(edges, Edge{U: u, V: v, W: weight})
			}
		}
	}
	return edges
}

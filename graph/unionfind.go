package graph

type DisjointSet struct {
	Parent []int
	Rank   []int
}

// NewDisjointSet создает структуру данных DisjointSet с n элементами
func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &DisjointSet{
		Parent: parent,
		Rank:   rank,
	}
}

// Find находит корень компоненты с сжатием пути
func (ds *DisjointSet) Find(x int) int {
	if ds.Parent[x] != x {
		ds.Parent[x] = ds.Find(ds.Parent[x]) // Сжатие пути
	}
	return ds.Parent[x]
}

// Union объединяет две компоненты с использованием объединения по рангу
func (ds *DisjointSet) Union(x, y int) bool {
	rootX := ds.Find(x)
	rootY := ds.Find(y)

	if rootX == rootY {
		return false
	}

	// Объединяем по рангу
	if ds.Rank[rootX] > ds.Rank[rootY] {
		ds.Parent[rootY] = rootX
	} else if ds.Rank[rootX] < ds.Rank[rootY] {
		ds.Parent[rootX] = rootY
	} else {
		ds.Parent[rootY] = rootX // rootX становится родителем
		ds.Rank[rootX]++         // Увеличиваем ранг
	}

	return true
}

package graph

import "math"

func BoruvkaMST(n int, edges []Edge) (mst []Edge, totalWeight int) {
	// Инициализация множества компонент
	ds := NewDisjointSet(n)
	mst = []Edge{}
	totalWeight = 0

	numComponents := n

	// Пока в графе больше одной компоненты связности
	for numComponents > 1 {
		// Массив для хранения минимальных рёбер для каждой компоненты
		minEdges := make([]Edge, n)

		// Инициализируем минимальные рёбра значением с очень большим весом
		for i := range minEdges {
			minEdges[i] = Edge{-1, -1, math.MaxInt32}
		}

		// Ищем минимальные рёбра для каждой компоненты
		for _, edge := range edges {
			u, v, w := edge.U, edge.V, edge.W

			// Находим компоненты для обеих вершин
			compU := ds.Find(u)
			compV := ds.Find(v)

			// Если вершины в разных компонентах, пытаемся обновить минимальное ребро
			if compU != compV {
				// Для компоненты u
				if w < minEdges[compU].W {
					minEdges[compU] = edge
				}
				// Для компоненты v
				if w < minEdges[compV].W {
					minEdges[compV] = edge
				}
			}
		}

		// Добавляем минимальные рёбра в результат
		for _, edge := range minEdges {
			// Пропускаем рёбра с значением -1, которые не были обновлены
			if edge.U == -1 && edge.V == -1 {
				continue
			}

			u, v, w := edge.U, edge.V, edge.W

			// Если вершины в разных компонентах, объединяем их
			if ds.Find(u) != ds.Find(v) {
				ds.Union(u, v)
				mst = append(mst, edge)
				totalWeight += w
				numComponents--
			}
		}
	}

	return mst, totalWeight
}

func MergeSort(edges []Edge) []Edge {
	if len(edges) <= 1 {
		return edges
	}
	mid := len(edges) / 2
	left := MergeSort(edges[:mid])
	right := MergeSort(edges[mid:])

	return Merge(left, right)
}

func Merge(left, right []Edge) []Edge {
	var result []Edge
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].W <= right[j].W {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}

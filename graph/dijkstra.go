package graph

import (
	"math"
)

func Dijkstra(graph Graph, start int) map[int]int {
	// Создаём таблицу расстояний, изначально все - "бесконечность"
	distances := make(map[int]int)
	for i := range graph.Adj {
		distances[i] = math.MaxInt32
	}
	distances[start] = 0 // Начальная вершина

	// Очередь с приоритетом
	pq := NewPriorityQueue()
	pq.Push(start, 0)

	// Основной цикл алгоритма
	for len(pq.Data) > 0 {
		// Берём вершину с минимальным расстоянием
		current := pq.Pop()
		currentNode := current.Vertex
		currentDistance := current.Dist

		// Если найденный путь длиннее, чем уже известный, пропускаем
		if currentDistance > distances[currentNode] {
			continue
		}

		// Обновляем расстояния до соседей
		for _, edge := range graph.GetNeighbors() {
			newDistance := currentDistance + edge.W
			if newDistance < distances[edge.v] {
				distances[edge.to] = newDistance
				pq.Push(edge.v, newDistance)
			}
		}
	}

	return distances
}

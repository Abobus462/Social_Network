package graph

import (
	"math"
)

func BellmanFord(g *Graph, start int) ([]int, []int, bool) {
	// Количество вершин
	V := len(g.Adj)

	// Инициализация массивов
	dist := make([]int, V)
	prev := make([]int, V)
	// Устанавливаем все расстояния как бесконечность
	for i := 0; i < V; i++ {
		dist[i] = math.MaxInt
		prev[i] = -1
	}
	// Расстояние до начальной вершины равно 0
	dist[start] = 0

	// Основной цикл алгоритма (проходим V-1 раз)
	for i := 1; i < V; i++ {
		for _, edge := range g.Edge {
			// Если найден более короткий путь через ребро
			if dist[edge.U] != math.MaxInt && dist[edge.U]+edge.W < dist[edge.V] {
				dist[edge.V] = dist[edge.U] + edge.W
				prev[edge.V] = edge.U
			}
		}
	}

	// Проверка на наличие отрицательных циклов
	negativeCycle := false
	for _, edge := range g.Edge {
		if dist[edge.U] != math.MaxInt && dist[edge.U]+edge.W < dist[edge.V] {
			negativeCycle = true
			break
		}
	}

	return dist, prev, negativeCycle
}

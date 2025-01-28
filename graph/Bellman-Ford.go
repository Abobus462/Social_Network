package graph

import (
	"math"
)

func BellmanFord(g *Graph, start int) ([]int, []int, bool) {
	// Количество вершин
	V := len(g.Adj)
	// Инициализация массивов
	dist := make([]int, V) // Расстояния до вершин
	prev := make([]int, V) // Предыдущие вершины (для восстановления пути)

	// Инициализация: все расстояния - бесконечность, кроме стартовой вершины
	for i := range dist {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[start] = 0

	// Основной цикл (V-1 итераций)
	for i := 0; i < V-1; i++ {
		updated := false // Флаг, если не было изменений - можно остановиться
		for _, edge := range g.GetAllEdges() {
			u, v, w := edge.U, edge.V, edge.W
			if dist[u] != math.MaxInt32 && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				prev[v] = u
				updated = true
			}
		}
		if !updated {
			break // Если на итерации не было изменений, то можно остановиться
		}
	}

	// Проверка наличия отрицательного цикла
	for _, edge := range g.GetAllEdges() {
		u, v, w := edge.U, edge.V, edge.W
		if dist[u] != math.MaxInt32 && dist[u]+w < dist[v] {
			return nil, nil, true // Обнаружен отрицательный цикл
		}
	}

	return dist, prev, false // Возвращаем расстояния, пути и флаг наличия цикла
}

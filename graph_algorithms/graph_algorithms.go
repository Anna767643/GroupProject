package graphalgorithms

import (
	"fmt"
	graph "simplenavigator/graph"
)

// Очень большое число 1*10 в 9ой степени (1 000 000 000)
var inf int = 1e9

/*
Остовное дерево в связном неориентированном графе — это такой подграф,
который включает в себя все вершины исходного графа,
является связным, не содержит циклов и состоит ровно из n - 1 ребра
(где n — количество вершин)
*/

/*
Для взвешенных графов часто ищут минимальное остовное дерево
(MST — Minimum Spanning Tree) — остов,
сумма весов ребер которого минимальна
*/

/*
Алгоритм Прима:
• Начинает построение с любой выбранной стартовой вершины.
• На каждом шаге выбирает ребро минимального веса,
которое соединяет уже построенную часть дерева с еще непосещенной вершиной.
• Повторяет шаг, пока в дерево не войдут все вершины.
*/

func GetLeastSpanningTree(graph *graph.Graph) ([][]int, error) {
	// Проверяет, что б граф был неориентированный
	if graph.IsDirected() {
		return [][]int{}, fmt.Errorf("graph must not be directed")
	}

	size := graph.Size()

	inTree := make([]bool, size+1)
	minEdge := make([]int, size+1)
	parent := make([]int, size+1)

	for i := range minEdge {
		// Устанавливаем большое число
		minEdge[i] = inf
		// Маркер, что родителя нет
		parent[i] = -1
	}

	minEdge[1] = 0

	resultMatrix := make([][]int, size)
	for i := range resultMatrix {
		resultMatrix[i] = make([]int, size)
	}

	for step := 0; step < size; step++ {
		value := -1

		// Ищем вершину с минимальным весом ребра, которая еще не в дереве
		for u := 1; u <= size; u++ {
			if !inTree[u] && (value == -1 || minEdge[u] < minEdge[value]) {
				value = u
			}
		}

		// Помечаем найденную вершину как добавленную в дерево
		inTree[value] = true

		// В первом цикли ничего не устанавливается
		if p := parent[value]; p != -1 {
			w := graph.Weight(p, value)
			resultMatrix[p-1][value-1], resultMatrix[value-1][p-1] = w, w
		}

		neighbors, _ := graph.Neighbors(value)
		for _, neighbor := range neighbors {
			if weight := graph.Weight(value, neighbor); !inTree[neighbor] && weight < minEdge[neighbor] {
				minEdge[neighbor] = weight
				parent[neighbor] = value
			}
		}
	}

	return resultMatrix, nil
}

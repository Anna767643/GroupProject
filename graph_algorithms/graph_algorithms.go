package s21_graph_algorithms

import (
	"fmt"
	graph "simplenavigator/graph"
	"simplenavigator/minHeap"
	queue "simplenavigator/queue"
	stack "simplenavigator/stack"
)

// Очень большое число 1*10 в 9ой степени (1 000 000 000)
var inf int = 1e9

//          1
//         / \
//        2   3
//         \ /
//          4
//graph := [][]int{
//{0, 1, 1, 0},
//{1, 0, 0, 1},
//{1, 0, 0, 1},
//{0, 1, 1, 0},
//}

func DepthFirstSearch(graph *graph.Graph, start_vertex int) []int {

	row := stack.New()
	row.Push(start_vertex)
	arr := make(map[int]struct{})
	arr[start_vertex] = struct{}{}
	res := make([]int, 0, graph.Size())

	for !row.IsEmpty() {
		sl, flag := row.Pop()

		if flag {
			if _, ok := arr[sl]; ok {
				continue
			}
			arr[sl] = struct{}{}
			res = append(res, sl)

			neighbors, err := graph.Neighbors(sl)
			if err != nil {
				return nil
			}

			for _, neighbor := range neighbors {
				if _, ok := arr[neighbor]; !ok {
					row.Push(neighbor)
				}
			}
		}
	}

	return res
}

func BreadthFirstSearch(graph *graph.Graph, start_vertex int) []int {

	row := queue.New()
	row.Push(start_vertex)

	arr := make(map[int]struct{})
	res := make([]int, 0, graph.Size())
	arr[start_vertex] = struct{}{}

	for !row.IsEmpty() {
		sl, flag := row.Pop()

		if flag {
			res = append(res, sl)
			neighbors, err := graph.Neighbors(sl)
			if err != nil {
				return nil
			}

			for _, neighbor := range neighbors {
				if _, ok := arr[neighbor]; !ok {
					row.Push(neighbor)
					arr[neighbor] = struct{}{}
				}
			}
		}
	}

	return res
}

func GetShortestPathBetweenVertices(graph *graph.Graph, vertex1 int, vertex2 int) (int, bool) {

	if vertex1 < 1 || vertex1 > graph.Size() ||
		vertex2 < 1 || vertex2 > graph.Size() {
		return 0, false
	}
	distances := make([]int, graph.Size())

	for i, _ := range distances {
		distances[i] = inf
	}

	distances[vertex1-1] = 0

	heap := minHeap.New()

	heap.Push(minHeap.Item{
		Vertex:   vertex1,
		Distance: 0,
	})

	for !heap.IsEmpty() {
		el, _ := heap.Pop()

		if el.Distance > distances[el.Vertex-1] {
			continue
		}
		neighbors, err := graph.Neighbors(el.Vertex)
		if err != nil {
			return 0, false
		}
		for _, neighbor := range neighbors {
			weight := graph.Weight(el.Vertex, neighbor)
			newDistance := el.Distance + weight

			if newDistance < distances[neighbor-1] {
				distances[neighbor-1] = newDistance

				heap.Push(minHeap.Item{
					Vertex:   neighbor,
					Distance: newDistance,
				})
			}
		}
	}

	if distances[vertex2-1] == inf {
		return 0, false
	}

	return distances[vertex2-1], true
}

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

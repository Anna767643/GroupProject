package graph_algorithms

import (
	"s21_graph_algorithms/minHeap"
	"s21_graph_algorithms/queue"
	"s21_graph_algorithms/stack"
)

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

func DepthFirstSearch(graph *Graph, start_vertex int) []int {

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

func BreadthFirstSearch(graph *Graph, start_vertex int) []int {

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

func GetShortestPathBetweenVertices(graph *Graph, vertex1 int, vertex2 int) (int, bool) {

	if vertex1 < 1 || vertex1 > graph.Size() ||
		vertex2 < 1 || vertex2 > graph.Size() {
		return 0, false
	}
	distances := make([]int, graph.Size())

	for i, _ := range distances {
		distances[i] = INF
	}

	distances[vertex1-1] = 0

	heap := minHeap.New()

	heap.Push(Item{
		Vertex:   vertex1,
		Distance: 0,
	})

	for !heap.IsEmpty() {
		el, _ := heap.Pop()

		if el.Distance > distances[el.Vertex-1] {
			continue
		}
		for _, neighbor := range graph.Neighbors(el.Vertex) {
			weight := graph.Weight(el.Vertex, neighbor)
			newDistance := el.Distance + weight

			if newDistance < distances[neighbor-1] {
				distances[neighbor-1] = newDistance

				heap.Push(Item{
					Vertex:   neighbor,
					Distance: newDistance,
				})
			}
		}
	}

	if distances[vertex2-1] == INF {
		return 0, false
	}

	return distances[vertex2-1], true
}

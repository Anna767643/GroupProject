package graph_algorithms

import (
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
			for _, neighbor := range graph.Neighbours(sl) {
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

			for _, neighbor := range graph.Neighbours(sl) {
				if _, ok := arr[neighbor]; !ok {
					row.Push(neighbor)
					arr[neighbor] = struct{}{}
				}
			}
		}
	}

	return res
}

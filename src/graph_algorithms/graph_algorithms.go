package main

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
	//тут изменим пока нет метода который возвращает количество вершин графа
	//res := make([]int, 0, len(matrix))

	res := make([]int, 0)

	for !row.IsEmpty() {
		sl, flag := row.Pop()

		if flag {
			if _, ok := arr[sl]; ok {
				continue
			}
			arr[sl] = struct{}{}
			res = append(res, sl)
			for _, neighbor := range graph.GetAdjacentVertices(sl) {
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
	//res := make([]int, 0, len(matrix))
	res := make([]int, 0)
	arr[start_vertex] = struct{}{}

	for !row.IsEmpty() {
		sl, flag := row.Pop()

		if flag {
			res = append(res, sl)
			//for j, _ := range matrix {
			//	if _, ok := arr[j+1]; !ok {
			//		if matrix[sl-1][j] == 1 {
			//			row.Push(j + 1)
			//			arr[j+1] = struct{}{}
			//		}
			//	}
			//}

			for _, neighbor := range graph.GetAdjacentVertices(sl) {
				if _, ok := arr[neighbor]; !ok {
					row.Push(neighbor)
					arr[neighbor] = struct{}{}
				}
			}
		}
	}

	return res
}

//func GetShortestPathBetweenVertices(graph *Graph, vertex1 int, vertex2 int){}

//func main() {
//
//	graph := [][]int{
//		{0, 5, 2, 0},
//		{5, 0, 0, 1},
//		{2, 0, 0, 7},
//		{0, 1, 7, 0},
//	}
//	arr := make(map[int]int)
//	//arr := make([]int, 0, len(graph))
//	vertex1 := 1
//	//vertex2 := 4
//	arr[vertex1] = 0
//	for i, _ := range graph {
//		for j, _ := range graph {
//			if _, ok := arr[i+1]; !ok {
//				arr[i+1] += graph[i][j]
//				//fmt.Println(graph[i][j])
//			}
//		}
//	}
//	fmt.Println(arr)
//
//}

package s21_graph_algorithms

import (
	"fmt"
	"math"
	"math/rand"
	graph "simplenavigator/graph"
	"simplenavigator/minHeap"
	queue "simplenavigator/queue"
	stack "simplenavigator/stack"
)

const inf = math.MaxInt / 2

func DepthFirstSearch(graph *graph.Graph, start_vertex int) []int {

	row := stack.New()
	row.Push(start_vertex)
	arr := make(map[int]struct{})
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

	for i := range distances {
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

func GetShortestPathsBetweenAllVertices(graph *graph.Graph) [][]int {
	n := graph.Size()
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
	}

	for i := range dist {
		for j := range dist {
			w := graph.Weight(i+1, j+1)
			if w == 0 {
				w = inf
			}
			if i == j {
				w = 0
			}
			dist[i][j] = w
		}
	}

	for k := range dist {
		for i := range dist {
			for j := range dist {
				if dist[i][k] == inf || dist[k][j] == inf {
					continue
				}
				sum := dist[i][k] + dist[k][j]
				if dist[i][j] > sum {
					dist[i][j] = sum
				}
			}
		}
	}

	return dist
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

// TsmResult - описывает найденный маршрут коммивояжера
type TsmResult struct {
	Vertices []int // порядок вершин с 1, первая вершина повторяется в конце
	Distance int   // длина маршрута
}

// candidate - вершина и феромон
type candidate struct {
	node int     // индекс вершины
	prob float64 // вычисленная привлекательность
}

// SolveTravelingSalesmanProblem - решение задачи коммивояжера с помощью муравьиного алгоритма
func SolveTravelingSalesmanProblem(graph *graph.Graph) (*TsmResult, error) {
	size := graph.Size()
	if size == 0 {
		return nil, fmt.Errorf("size is zero")
	}

	if size == 1 {
		return &TsmResult{
			Vertices: []int{1, 1},
			Distance: 0,
		}, nil
	}

	// Инициализация матрицы феромонов
	tau := make([][]float64, size)
	for i := 0; i < size; i++ {
		tau[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			if i != j {
				tau[i][j] = 1.0
			}
		}
	}

	bestResult := &TsmResult{
		Vertices: nil,
		Distance: math.MaxInt,
	}
	found := false
	antsCount := size

	localPheromone := make([][]float64, size)
	for i := 0; i < size; i++ {
		localPheromone[i] = make([]float64, size)
	}

	iterations := 300

	for iter := 0; iter < iterations; iter++ {
		// Очистка локального обновления феромона для новой итерации
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				localPheromone[i][j] = 0
			}
		}

		// Поиск решения муравьями
		for ant := 0; ant < antsCount; ant++ {
			// 1. Получаем результат конкретного муравья
			foundPath, antResult := createPathWithOneAnt(graph, tau, ant, localPheromone)

			// 2. Обновляем глобальный лучший результат ТОЛЬКО если он действительно лучше
			if foundPath && antResult.Distance < bestResult.Distance {
				bestResult = antResult
				found = true
			}
		}

		updateTau(tau, localPheromone, size)
	}

	if !found {
		return nil, fmt.Errorf("the route does not exist")
	}

	return bestResult, nil
}

func createPathWithOneAnt(g *graph.Graph, tau [][]float64, ant int, localPheromone [][]float64) (bool, *TsmResult) {
	size := g.Size()
	start := ant % size
	path := []int{start}
	current := start
	distance := 0
	valid := true

	visited := make([]bool, size)
	visited[start] = true

	for len(path) < size {
		cands, sumProb := findNeighbors(g, current, tau, visited)

		if len(cands) == 0 {
			valid = false
			break
		}

		var nextNode int
		if len(cands) == 1 {
			nextNode = cands[0].node
		} else {
			r := rand.Float64() * sumProb
			nextNode = cands[len(cands)-1].node // fallback
			for _, c := range cands {
				r -= c.prob
				if r <= 0 {
					nextNode = c.node
					break
				}
			}
		}

		path = append(path, nextNode)
		visited[nextNode] = true // Отмечаем как посещенную для следующих шагов

		// +1, так как graph.Weight ожидает 1-индексацию
		distance += g.Weight(current+1, nextNode+1)
		current = nextNode
	}

	if valid && len(path) == size {
		wBack := g.Weight(current+1, start+1)
		if wBack > 0 {
			distance += wBack
			path = append(path, start) // замыкаем цикл

			resVertices := make([]int, len(path))
			for i, v := range path {
				resVertices[i] = v + 1 // переводим в 1-индексацию для ответа
			}

			// Локальное обновление феромона
			Q := 100.0
			delta := Q / float64(distance)
			for i := 0; i < len(path)-1; i++ {
				u := path[i]
				v := path[i+1]
				localPheromone[u][v] += delta
			}

			return true, &TsmResult{
				Vertices: resVertices,
				Distance: distance,
			}
		}
	}

	return false, &TsmResult{}
}

func findNeighbors(g *graph.Graph, current int, tau [][]float64, visited []bool) ([]candidate, float64) {
	alpha := 1.0
	beta := 2.0

	size := g.Size()
	var cands []candidate
	sumProb := 0.0

	for nextNode := 0; nextNode < size; nextNode++ {
		if !visited[nextNode] {
			weight := g.Weight(current+1, nextNode+1)
			if weight > 0 {
				heuristic := 1.0 / float64(weight)
				prob := math.Pow(tau[current][nextNode], alpha) * math.Pow(heuristic, beta)

				cands = append(cands, candidate{node: nextNode, prob: prob})
				sumProb += prob
			}
		}
	}

	return cands, sumProb
}

// updateTau - глобальное обновление феромона (испарение + локальное обновление)
func updateTau(tau [][]float64, localPheromone [][]float64, size int) {
	rho := 0.5 // коэффициент испарения феромона

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i != j {
				tau[i][j] = tau[i][j]*(1-rho) + localPheromone[i][j]
				if tau[i][j] < 0.01 {
					tau[i][j] = 0.01
				}
			}
		}
	}
}

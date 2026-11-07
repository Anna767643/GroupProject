package graph_algorithms

import (
	"slices"
	"testing"
)

// Порядок обхода зафиксирован контрактом Graph.Neighbors:
// соседи идут по возрастанию номера и кладутся в стек/очередь слева направо.
// У стека следующим берётся последний добавленный сосед.

func TestDepthFirstSearch(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		start int
		want  []int
	}{
		{
			name:  "ромб от 1",
			graph: diamond,
			start: 1,
			want:  []int{1, 3, 4, 2},
		},
		{
			name:  "ромб от 2",
			graph: diamond,
			start: 2,
			want:  []int{2, 4, 3, 1},
		},
		{
			name:  "ромб от 3",
			graph: diamond,
			start: 3,
			want:  []int{3, 4, 2, 1},
		},
		{
			name:  "ромб от 4",
			graph: diamond,
			start: 4,
			want:  []int{4, 3, 1, 2},
		},
		{
			name:  "цепь от 1",
			graph: chain,
			start: 1,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "цепь от 2",
			graph: chain,
			start: 2,
			want:  []int{2, 3, 4, 1},
		},
		{
			name:  "звезда от центра",
			graph: star,
			start: 1,
			want:  []int{1, 4, 3, 2},
		},
		{
			name:  "дерево от корня",
			graph: tree,
			start: 1,
			want:  []int{1, 5, 2, 4, 3},
		},
		{
			name:  "только своя компонента",
			graph: disconnected,
			start: 1,
			want:  []int{1, 2},
		},
		{
			name:  "вторая компонента",
			graph: disconnected,
			start: 3,
			want:  []int{3, 4},
		},
		{
			name:  "ориентированный граф",
			graph: directed,
			start: 1,
			want:  []int{1, 3, 4, 2},
		},
		{
			name:  "тупик ориентированного графа",
			graph: directed,
			start: 4,
			want:  []int{4},
		},
		{
			name:  "повторное обнаружение вершины в стеке",
			graph: stackOrder,
			start: 1,
			want:  []int{1, 4, 2, 5, 3},
		},
		{
			name:  "одна вершина",
			graph: single,
			start: 1,
			want:  []int{1},
		},
		{
			name:  "изолированная вершина",
			graph: noEdge,
			start: 1,
			want:  []int{1},
		},
		{
			name:  "петля не дублирует вершину",
			graph: selfLoop,
			start: 1,
			want:  []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DepthFirstSearch(newTestGraph(t, tt.graph), tt.start)
			if !slices.Equal(got, tt.want) {
				t.Errorf("DepthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDepthFirstSearchInvalidStart(t *testing.T) {
	graph := newTestGraph(t, diamond)
	for _, start := range []int{0, -1, 5} {
		got := DepthFirstSearch(graph, start)
		if got != nil {
			t.Errorf("DepthFirstSearch(%d) = %v, want nil", start, got)
		}
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		start int
		want  []int
	}{
		{
			name:  "ромб от 1",
			graph: diamond,
			start: 1,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "ромб от 2",
			graph: diamond,
			start: 2,
			want:  []int{2, 1, 4, 3},
		},
		{
			name:  "ромб от 3",
			graph: diamond,
			start: 3,
			want:  []int{3, 1, 4, 2},
		},
		{
			name:  "ромб от 4",
			graph: diamond,
			start: 4,
			want:  []int{4, 2, 3, 1},
		},
		{
			name:  "цепь от 1",
			graph: chain,
			start: 1,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "цепь от 2",
			graph: chain,
			start: 2,
			want:  []int{2, 1, 3, 4},
		},
		{
			name:  "звезда от центра",
			graph: star,
			start: 1,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "дерево от корня",
			graph: tree,
			start: 1,
			want:  []int{1, 2, 5, 3, 4},
		},
		{
			name:  "только своя компонента",
			graph: disconnected,
			start: 1,
			want:  []int{1, 2},
		},
		{
			name:  "вторая компонента с другого конца",
			graph: disconnected,
			start: 4,
			want:  []int{4, 3},
		},
		{
			name:  "ориентированный граф",
			graph: directed,
			start: 1,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "тупик ориентированного графа",
			graph: directed,
			start: 4,
			want:  []int{4},
		},
		{
			name:  "вершина на стеке не меняет порядок BFS",
			graph: stackOrder,
			start: 1,
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "одна вершина",
			graph: single,
			start: 1,
			want:  []int{1},
		},
		{
			name:  "изолированная вершина",
			graph: noEdge,
			start: 1,
			want:  []int{1},
		},
		{
			name:  "петля не дублирует вершину",
			graph: selfLoop,
			start: 1,
			want:  []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BreadthFirstSearch(newTestGraph(t, tt.graph), tt.start)
			if !slices.Equal(got, tt.want) {
				t.Errorf("BreadthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBreadthFirstSearchInvalidStart(t *testing.T) {
	graph := newTestGraph(t, diamond)
	for _, start := range []int{0, -1, 5} {
		got := BreadthFirstSearch(graph, start)
		if got != nil {
			t.Errorf("BreadthFirstSearch(%d) = %v, want nil", start, got)
		}
	}
}

func TestDepthAndBreadthDifferOnBranch(t *testing.T) {
	graph := newTestGraph(t, diamond)
	dfs := DepthFirstSearch(graph, 1)
	bfs := BreadthFirstSearch(graph, 1)
	if slices.Equal(dfs, bfs) {
		t.Errorf("DFS and BFS returned the same order %v", dfs)
	}
}

func TestGetShortestPathBetweenVertices(t *testing.T) {
	tests := []struct {
		name   string
		graph  [][]int
		from   int
		to     int
		want   int
		wantOK bool
	}{
		{name: "та же вершина", graph: diamond, from: 1, to: 1, want: 0, wantOK: true},
		{name: "сосед в ромбе", graph: diamond, from: 1, to: 2, want: 1, wantOK: true},
		{name: "диагональ ромба", graph: diamond, from: 1, to: 4, want: 2, wantOK: true},
		{name: "путь обратно", graph: diamond, from: 4, to: 1, want: 2, wantOK: true},
		{name: "ребро своей компоненты", graph: disconnected, from: 3, to: 4, want: 1, wantOK: true},
		{name: "чужая компонента", graph: disconnected, from: 1, to: 4, want: 0, wantOK: false},
		{name: "нет ребра", graph: noEdge, from: 1, to: 2, want: 0, wantOK: false},
		{name: "единственная вершина", graph: single, from: 1, to: 1, want: 0, wantOK: true},
		{name: "дешёвый обходной путь", graph: cheapLongPath, from: 1, to: 4, want: 3, wantOK: true},
		{name: "прямой путь дороже", graph: directIsWorse, from: 1, to: 3, want: 3, wantOK: true},
		{name: "ориентированный 1 к 4", graph: weightedDirected, from: 1, to: 4, want: 6, wantOK: true},
		{name: "ориентированный 1 к 3", graph: weightedDirected, from: 1, to: 3, want: 3, wantOK: true},
		{name: "ориентированный 2 к 4", graph: weightedDirected, from: 2, to: 4, want: 5, wantOK: true},
		{name: "обратного ребра нет", graph: weightedDirected, from: 4, to: 1, want: 0, wantOK: false},
		{name: "из середины назад нельзя", graph: weightedDirected, from: 3, to: 1, want: 0, wantOK: false},
		{name: "два равных пути", graph: equalPaths, from: 1, to: 4, want: 4, wantOK: true},
		{name: "обход дешевле жадного", graph: weightedUndirected, from: 1, to: 4, want: 8, wantOK: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := GetShortestPathBetweenVertices(newTestGraph(t, tt.graph), tt.from, tt.to)
			if ok != tt.wantOK || got != tt.want {
				t.Errorf("GetShortestPathBetweenVertices(%d, %d) = (%d, %t), want (%d, %t)",
					tt.from, tt.to, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

func TestGetShortestPathBetweenVerticesInvalidVertex(t *testing.T) {
	graph := newTestGraph(t, diamond)
	tests := []struct {
		from int
		to   int
	}{
		{from: 0, to: 1},
		{from: 1, to: 0},
		{from: -1, to: 2},
		{from: 1, to: 5},
		{from: 5, to: 5},
	}

	for _, tt := range tests {
		got, ok := GetShortestPathBetweenVertices(graph, tt.from, tt.to)
		if ok || got != 0 {
			t.Errorf("GetShortestPathBetweenVertices(%d, %d) = (%d, %t), want (0, false)",
				tt.from, tt.to, got, ok)
		}
	}
}

// diamond: 1—2, 1—3, 2—4, 3—4.
var diamond = [][]int{
	{0, 1, 1, 0},
	{1, 0, 0, 1},
	{1, 0, 0, 1},
	{0, 1, 1, 0},
}

// chain: 1—2—3—4.
var chain = [][]int{
	{0, 1, 0, 0},
	{1, 0, 1, 0},
	{0, 1, 0, 1},
	{0, 0, 1, 0},
}

// star: центр 1 соединён с 2, 3 и 4.
var star = [][]int{
	{0, 1, 1, 1},
	{1, 0, 0, 0},
	{1, 0, 0, 0},
	{1, 0, 0, 0},
}

// tree: 1—2, 1—5, 2—3, 2—4.
var tree = [][]int{
	{0, 1, 0, 0, 1},
	{1, 0, 1, 1, 0},
	{0, 1, 0, 0, 0},
	{0, 1, 0, 0, 0},
	{1, 0, 0, 0, 0},
}

// disconnected: компоненты {1, 2} и {3, 4}.
var disconnected = [][]int{
	{0, 1, 0, 0},
	{1, 0, 0, 0},
	{0, 0, 0, 1},
	{0, 0, 1, 0},
}

// directed: 1→2→4, 1→3→4.
var directed = [][]int{
	{0, 1, 1, 0},
	{0, 0, 0, 1},
	{0, 0, 0, 1},
	{0, 0, 0, 0},
}

// stackOrder отличает пометку при извлечении из стека от пометки при добавлении.
// Рёбра: 1→2, 1→3, 1→4, 4→2, 2→5.
var stackOrder = [][]int{
	{0, 1, 1, 1, 0},
	{0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0},
	{0, 0, 0, 0, 0},
}

var single = [][]int{
	{0},
}

var noEdge = [][]int{
	{0, 0},
	{0, 0},
}

// selfLoop: петля веса 2 у вершины 1 и ребро 1—2.
var selfLoop = [][]int{
	{2, 1},
	{1, 0},
}

// cheapLongPath: 1→2→3→4 весит 3, прямое ребро 1→4 весит 100.
var cheapLongPath = [][]int{
	{0, 1, 0, 100},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
	{0, 0, 0, 0},
}

// directIsWorse: прямое 1—3 весит 10, путь 1—2—3 весит 3.
var directIsWorse = [][]int{
	{0, 2, 10},
	{2, 0, 1},
	{10, 1, 0},
}

// weightedDirected:
// 1→2: 1, 1→3: 4, 2→3: 2, 2→4: 6, 3→4: 3.
var weightedDirected = [][]int{
	{0, 1, 4, 0},
	{0, 0, 2, 6},
	{0, 0, 0, 3},
	{0, 0, 0, 0},
}

// equalPaths: 1—2—4 и 1—3—4, оба веса 4.
var equalPaths = [][]int{
	{0, 2, 2, 0},
	{2, 0, 0, 2},
	{2, 0, 0, 2},
	{0, 2, 2, 0},
}

// weightedUndirected: кратчайший 1—4 равен 8 (1—3—2—4), а не 9 и не 10.
var weightedUndirected = [][]int{
	{0, 4, 2, 0},
	{4, 0, 1, 5},
	{2, 1, 0, 8},
	{0, 5, 8, 0},
}

func newTestGraph(t *testing.T, matrix [][]int) *Graph {
	t.Helper()

	n := len(matrix)
	copied := make([][]int, n)
	for i, row := range matrix {
		if len(row) != n {
			t.Fatalf("row %d has %d columns, want %d", i, len(row), n)
		}
		copied[i] = append([]int(nil), row...)
	}
	return &Graph{size: n, matrix: copied}
}

package s21_graph_algorithms

import (
	"fmt"
	"os"
	"path/filepath"
	graph "simplenavigator/graph"
	"slices"
	"strings"
	"testing"
)

func TestGetLeastSpanningTree(t *testing.T) {
	testCases := []struct {
		testName       string
		filename       string
		expectedError  string
		expectedSize   int
		expectedMatrix [][]int
	}{
		{
			testName:      "graph must not be directed",
			filename:      "../test_data/weighted_directed.txt",
			expectedError: "graph must not be directed",
		},
		{
			testName:      "OK",
			filename:      "../test_data/matrix.txt",
			expectedError: "",
			expectedSize:  4,
			expectedMatrix: [][]int{
				{0, 10, 0, 5},
				{10, 0, 0, 0},
				{0, 0, 0, 4},
				{5, 0, 4, 0},
			},
		},
		{
			testName:      "OK",
			filename:      "../test_data/weighted.txt",
			expectedError: "",
			expectedSize:  11,
			expectedMatrix: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 12, 4, 0, 0},
				{0, 0, 0, 0, 0, 0, 72, 0, 0, 0, 12},
				{0, 0, 0, 0, 0, 0, 0, 9, 0, 0, 0},
				{0, 0, 0, 0, 4, 12, 0, 0, 0, 0, 0},
				{0, 0, 0, 4, 0, 0, 0, 9, 0, 0, 0},
				{0, 0, 0, 12, 0, 0, 0, 0, 0, 3, 0},
				{0, 72, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{12, 0, 9, 0, 9, 0, 0, 0, 0, 0, 13},
				{4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0},
				{0, 12, 0, 0, 0, 0, 0, 13, 0, 0, 0},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var graph graph.Graph
			err := graph.LoadGraphFromFile(tc.filename)
			if err != nil {
				t.Errorf("LoadGraphFromFile got error %q, but want nil", err)

			}

			matrix, err := GetLeastSpanningTree(&graph)
			size := graph.Size()

			if tc.expectedError == "" {
				if size != len(matrix) {
					t.Errorf("Expected size %d, but got matrix with len: %d", tc.expectedSize, len(matrix))
					return
				}

				for i := 0; i < size; i++ {
					for j := 0; j < size; j++ {
						if matrix[i][j] != tc.expectedMatrix[i][j] {
							t.Errorf("Expected number %d from matrix[%d][%d], but got: %d", tc.expectedMatrix[i][j], i, j, matrix[i][j])
						}
					}
				}

			} else {
				if !strings.Contains(strings.ToLower(err.Error()), tc.expectedError) {
					t.Errorf("Expected error %q, but got: %v", tc.expectedError, err)
				}
			}
		})
	}
}

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
	g := newTestGraph(t, diamond)
	dfs := DepthFirstSearch(g, 1)
	bfs := BreadthFirstSearch(g, 1)
	if len(dfs) == 0 || len(bfs) == 0 {
		t.Fatalf("traversal must be non-empty, DFS = %v, BFS = %v", dfs, bfs)
	}
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

// Нет пути — ячейка равна inf. Диагональ равна 0: петля не заменяет расстояние до самой себя.
func TestGetShortestPathsBetweenAllVertices(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  [][]int
	}{
		{
			name:  "ромб",
			graph: diamond,
			want: [][]int{
				{0, 1, 1, 2},
				{1, 0, 2, 1},
				{1, 2, 0, 1},
				{2, 1, 1, 0},
			},
		},
		{
			name:  "цепь",
			graph: chain,
			want: [][]int{
				{0, 1, 2, 3},
				{1, 0, 1, 2},
				{2, 1, 0, 1},
				{3, 2, 1, 0},
			},
		},
		{
			name:  "звезда",
			graph: star,
			want: [][]int{
				{0, 1, 1, 1},
				{1, 0, 2, 2},
				{1, 2, 0, 2},
				{1, 2, 2, 0},
			},
		},
		{
			name:  "дерево",
			graph: tree,
			want: [][]int{
				{0, 1, 2, 2, 1},
				{1, 0, 1, 1, 2},
				{2, 1, 0, 2, 3},
				{2, 1, 2, 0, 3},
				{1, 2, 3, 3, 0},
			},
		},
		{
			name:  "две компоненты",
			graph: disconnected,
			want: [][]int{
				{0, 1, inf, inf},
				{1, 0, inf, inf},
				{inf, inf, 0, 1},
				{inf, inf, 1, 0},
			},
		},
		{
			name:  "ориентированный",
			graph: directed,
			want: [][]int{
				{0, 1, 1, 2},
				{inf, 0, inf, 1},
				{inf, inf, 0, 1},
				{inf, inf, inf, 0},
			},
		},
		{
			name:  "одна вершина",
			graph: single,
			want:  [][]int{{0}},
		},
		{
			name:  "нет рёбер",
			graph: noEdge,
			want: [][]int{
				{0, inf},
				{inf, 0},
			},
		},
		{
			name:  "петля не меняет диагональ",
			graph: selfLoop,
			want: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name:  "обход дешевле прямого ребра",
			graph: cheapLongPath,
			want: [][]int{
				{0, 1, 2, 3},
				{inf, 0, 1, 2},
				{inf, inf, 0, 1},
				{inf, inf, inf, 0},
			},
		},
		{
			name:  "прямой путь дороже",
			graph: directIsWorse,
			want: [][]int{
				{0, 2, 3},
				{2, 0, 1},
				{3, 1, 0},
			},
		},
		{
			name:  "ориентированный взвешенный",
			graph: weightedDirected,
			want: [][]int{
				{0, 1, 3, 6},
				{inf, 0, 2, 5},
				{inf, inf, 0, 3},
				{inf, inf, inf, 0},
			},
		},
		{
			name:  "два равных пути",
			graph: equalPaths,
			want: [][]int{
				{0, 2, 2, 4},
				{2, 0, 4, 2},
				{2, 4, 0, 2},
				{4, 2, 2, 0},
			},
		},
		{
			name:  "промежуточная вершина короче жадного выбора",
			graph: weightedUndirected,
			want: [][]int{
				{0, 3, 2, 8},
				{3, 0, 1, 5},
				{2, 1, 0, 6},
				{8, 5, 6, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetShortestPathsBetweenAllVertices(newTestGraph(t, tt.graph))
			if !equalMatrix(got, tt.want) {
				t.Errorf("GetShortestPathsBetweenAllVertices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetShortestPathsBetweenAllVerticesMatchesPairwise(t *testing.T) {
	graphs := [][][]int{
		diamond, chain, star, tree, disconnected, directed,
		single, noEdge, selfLoop, cheapLongPath, directIsWorse,
		weightedDirected, equalPaths, weightedUndirected,
	}

	for _, matrix := range graphs {
		g := newTestGraph(t, matrix)
		got := GetShortestPathsBetweenAllVertices(g)
		if len(got) != len(matrix) {
			t.Fatalf("rows = %d, want %d", len(got), len(matrix))
		}
		for from := 1; from <= len(matrix); from++ {
			if len(got[from-1]) != len(matrix) {
				t.Fatalf("row %d has %d columns, want %d", from, len(got[from-1]), len(matrix))
			}
			for to := 1; to <= len(matrix); to++ {
				dist, ok := GetShortestPathBetweenVertices(g, from, to)
				want := inf
				if ok {
					want = dist
				}
				if got[from-1][to-1] != want {
					t.Errorf("path %d → %d = %d, want %d", from, to, got[from-1][to-1], want)
				}
			}
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

func newTestGraph(t *testing.T, matrix [][]int) *graph.Graph {
	t.Helper()

	n := len(matrix)
	var body strings.Builder
	fmt.Fprintf(&body, "%d\n", n)
	for i, row := range matrix {
		if len(row) != n {
			t.Fatalf("row %d has %d columns, want %d", i, len(row), n)
		}
		for j, value := range row {
			if j > 0 {
				body.WriteByte(' ')
			}
			fmt.Fprintf(&body, "%d", value)
		}
		body.WriteByte('\n')
	}

	path := filepath.Join(t.TempDir(), "graph.txt")
	if err := os.WriteFile(path, []byte(body.String()), 0o644); err != nil {
		t.Fatalf("write graph file: %v", err)
	}

	var g graph.Graph
	if err := g.LoadGraphFromFile(path); err != nil {
		t.Fatalf("LoadGraphFromFile: %v", err)
	}
	return &g
}

func equalMatrix(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range want {
		if !slices.Equal(got[i], want[i]) {
			return false
		}
	}
	return true
}

func TestSolveTravelingSalesmanProblem(t *testing.T) {
	testCases := []struct {
		testName         string
		size             int
		matrix           [][]int
		expectedError    string
		expectedDistance int
	}{
		{
			testName:      "graph size is zero",
			size:          0,
			matrix:        [][]int{}, // Пустой срез, чтобы пройти валидацию NewGraph
			expectedError: "size is zero",
		},
		{
			testName:         "graph size is one",
			size:             1,
			matrix:           [][]int{{0}},
			expectedError:    "",
			expectedDistance: 0,
		},
		{
			testName: "simple triangle graph (complete)",
			size:     3,
			matrix: [][]int{
				{0, 10, 15},
				{10, 0, 20},
				{15, 20, 0},
			},
			expectedError:    "",
			expectedDistance: 45, // 10 + 20 + 15 (единственный возможный цикл)
		},
		{
			testName: "disconnected graph (no valid route)",
			size:     3,
			matrix: [][]int{
				{0, 10, 0},
				{10, 0, 0},
				{0, 0, 0}, // Вершина 3 изолирована
			},
			expectedError: "the route does not exist",
		},
		{
			testName: "square graph with one expensive edge",
			size:     4,
			matrix: [][]int{
				{0, 1, 5, 5},
				{1, 0, 1, 5},
				{5, 1, 0, 1},
				{5, 5, 1, 0},
			},
			expectedError:    "",
			expectedDistance: 8, // Оптимальный путь: 1->2->3->4->1 (1 + 1 + 1 + 5)
		},
		{
			testName: "graph with missing edges but valid route exists",
			size:     4,
			matrix: [][]int{
				{0, 5, 0, 5}, // ИСПРАВЛЕНО: добавлено ребро 1-4 (вес 5)
				{5, 0, 5, 0},
				{0, 5, 0, 5},
				{5, 0, 5, 0}, // ИСПРАВЛЕНО: добавлено ребро 4-1 (вес 5)
			},
			expectedError:    "",
			expectedDistance: 20, // 5 + 5 + 5 + 5
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			g, err := graph.NewGraph(tc.size, tc.matrix)

			if tc.size > 0 && err != nil {
				t.Fatalf("Failed to create graph: %v", err)
			}

			result, err := SolveTravelingSalesmanProblem(g)

			if tc.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error containing %q, but got nil", tc.expectedError)
					return
				}
				if !strings.Contains(strings.ToLower(err.Error()), tc.expectedError) {
					t.Errorf("Expected error containing %q, but got: %v", tc.expectedError, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Expected no error, but got: %v", err)
				return
			}

			if result == nil {
				t.Errorf("Expected result, but got nil")
				return
			}

			if result.Distance != tc.expectedDistance {
				t.Errorf("Expected distance %d, but got: %d", tc.expectedDistance, result.Distance)
			}

			if !isValidTSPRoute(result, tc.size) {
				t.Errorf("Invalid route structure returned: %v", result.Vertices)
			}
		})
	}
}

func isValidTSPRoute(result *TsmResult, size int) bool {
	if size == 1 {
		return len(result.Vertices) == 2 && result.Vertices[0] == 1 && result.Vertices[1] == 1
	}

	if len(result.Vertices) != size+1 {
		return false
	}

	if result.Vertices[0] != result.Vertices[size] {
		return false
	}

	visited := make(map[int]bool)
	for i := 0; i < size; i++ {
		v := result.Vertices[i]
		if v < 1 || v > size {
			return false
		}
		if visited[v] {
			return false
		}
		visited[v] = true
	} // Вершина посещена дважды до замыкания цикла

	return len(visited) == size
}

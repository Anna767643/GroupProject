package graph

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	size   int
	matrix [][]int
}

/*
Неориентированный, невзвешенный, связный граф

	1 2 3 4 5
	___________

1 | 0 1 1 0 0
2 | 1 0 0 1 1
3 | 1 0 0 1 0
4 | 0 1 1 0 1
5 | 0 1 0 1 0
*/

var (
	ErrEmptyFile    = errors.New("graph: file is empty")
	ErrBadSize      = errors.New("graph: size must be a positive integer")
	ErrBadRow       = errors.New("graph: row has wrong number of columns")
	ErrNegativeEdge = errors.New("graph: weights must be non-negative")
	ErrVertice      = errors.New("graph: vertice must be less graph size or more zero")
)

func NewGraph(size int, matrix [][]int) (*Graph, error) {
	if matrix == nil {
		return nil, fmt.Errorf("matrix cannot be nil")
	}

	if len(matrix) != size {
		return nil, fmt.Errorf("matrix rows count mismatch: expected %d, got %d", size, len(matrix))
	}

	for i, row := range matrix {
		if len(row) != size {
			return nil, fmt.Errorf("invalid row length at index %d: expected %d, got %d", i, size, len(row))
		}
	}

	matrixCopy := make([][]int, size)
	for i := 0; i < size; i++ {
		matrixCopy[i] = make([]int, size)
		copy(matrixCopy[i], matrix[i])
	}

	return &Graph{
		size:   size,
		matrix: matrixCopy,
	}, nil
}

// Size - возвращает количество вершин
func (g *Graph) Size() int { return g.size }

// Neighbors - возвращает массив вершин соседей
// Ошибки:
// ErrVertice
func (g *Graph) Neighbors(vertice int) ([]int, error) {
	if vertice > g.size || vertice <= 0 {
		return []int{}, ErrVertice
	}

	neighbors := make([]int, 0)
	for column, value := range g.matrix[vertice-1] {
		if value != 0 {
			neighbors = append(neighbors, column+1)
		}
	}
	return neighbors, nil
}

// Weight - возвращает вес ребра verticeFirst -> verticeSecond
// 0 означает отсутствие ребра
func (g *Graph) Weight(verticeFirst, verticeSecond int) int {
	return g.matrix[verticeFirst-1][verticeSecond-1]
}

// IsDirected - сообщает ориентированный ли граф (матрица несимметрична)
// true - ориентированный
// false - неориентированный
func (g *Graph) IsDirected() bool {
	for i := 0; i < g.size; i++ {
		for j := i + 1; j < g.size; j++ {
			if g.matrix[i][j] != g.matrix[j][i] {
				return true
			}
		}
	}
	return false
}

/*
LoadGraphFromFile - загрузка графа из файла в формате матрицы смежности
Ошибки:
"graph: error to open file: "
"graph: error to read file: "
"graph: error to close file: "
ErrEmptyFile
ErrBadSize
ErrBadRow
"graph: bad number in line %d col %d: %w"
ErrNegativeEdge
*/
func (g *Graph) LoadGraphFromFile(filename string) error {
	bytes, err := g.readFile(filename)
	if err != nil {
		return err
	}

	if len(bytes) == 0 {
		return ErrEmptyFile
	}

	text := string(bytes)

	lines := strings.Split(text, "\n")
	lines[0] = strings.TrimSpace(lines[0])
	size, err := strconv.Atoi(lines[0])

	if err != nil || size <= 0 {
		return ErrBadSize
	}

	if len(lines) < size+1 {
		return ErrBadRow
	}

	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		numbers := strings.Fields(lines[i+1])
		if len(numbers) != size {
			return ErrBadRow
		}
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			matrix[i][j], err = strconv.Atoi(numbers[j])
			if err != nil {
				return fmt.Errorf("graph: bad number in line %d col %d: %w", i+2, j+1, err)
			}
			if matrix[i][j] < 0 {
				return ErrNegativeEdge
			}
		}
	}

	g.size = size
	g.matrix = matrix

	return nil
}

// ExportGraphToDot - выгрузка графа в файл в формате Dot
// Вершины нумеруются с 1
// Ошибки:
// При открытии или закрытия файла, отдадутся os.WriteFile
func (g *Graph) ExportGraphToDot(filename string) error {
	var stringB strings.Builder
	directed := g.IsDirected()
	description, dash := "graph graphname", "--"
	if directed {
		description, dash = "digraph graphname", "->"
	}
	fmt.Fprintf(&stringB, "%s {\n", description)
	for i := 1; i <= g.size; i++ {
		fmt.Fprintf(&stringB, "     %d;\n", i)
	}
	for i := 0; i < g.size; i++ {
		start := 0
		if !directed {
			start = i
		}
		for j := start; j < g.size; j++ {
			if weight := g.matrix[i][j]; weight != 0 && weight != 1 {
				fmt.Fprintf(&stringB, "     %d %s %d [weight=%d];\n", i+1, dash, j+1, weight)
			} else if weight != 0 {
				fmt.Fprintf(&stringB, "     %d %s %d;\n", i+1, dash, j+1)
			}
		}
	}
	stringB.WriteString("}")
	// Задаем дополнительно права доступа к файлу
	// - чтение и запись для владельца и все остальные ток чтение
	return os.WriteFile(filename, []byte(stringB.String()), 0o644)
}

/*
readFile - считывает файл в формате матрицы смежности
Возвращает массив байтов
Ошибки:
"graph: error to open file: "
"graph: error to read file: "
"graph: error to close file: "
*/
func (g *Graph) readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("graph: error to open file: %w", err)
	}

	data := make([]byte, 64)
	result := make([]byte, 0)

	for {
		n, err := file.Read(data)
		if n > 0 {
			result = append(result, data[:n]...)
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("graph: error to read file: %w", err)
		}
	}

	if err := file.Close(); err != nil {
		return []byte{}, fmt.Errorf("graph: error to close file: %w", err)
	}

	return result, nil
}

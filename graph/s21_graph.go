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

var (
	ErrReadFile     = errors.New("graph: failed to read file")
	ErrEmptyFile    = errors.New("graph: file is empty")
	ErrBadSize      = errors.New("graph: size must be a positive integer")
	ErrBadRow       = errors.New("graph: row has wrong number of columns")
	ErrNegativeEdge = errors.New("graph: weights must be non-negative")
)

// Size - возвращает количество вершин
func (g *Graph) Size() int { return g.size }

// Надо подумать какую еще нужно возвращать информацию

/*
LoadGraphFromFile - загрузка графа из файла в формате матрицы смежности
Ошибки:
ErrReadFile
ErrEmptyFile
ErrBadSize
ErrBadRow
"graph: bad number in line %d col %d: %w"
ErrNegativeEdge
*/
func (g *Graph) LoadGraphFromFile(filename string) error {
	bytes, err := g.readFile(filename)
	if err != nil {
		return ErrReadFile
	}

	if len(bytes) == 0 {
		return ErrEmptyFile
	}

	text := string(bytes)

	lines := strings.Split(text, "\n")
	lines[0] = strings.Trim(lines[0], " ")
	size, err := strconv.Atoi(lines[0])

	if err != nil || size <= 0 {
		return ErrBadSize
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

// ExportGraphToDot - выгрузка графа в файл в формате dot
func (g *Graph) ExportGraphToDot(filename string) {

}

/*
readFile - считывает файл в формате матрицы смежности
Возвращает массив байтов
Ошибки:
"error to open file: "
"error to read file: "
"error to close file: "
*/
func (g *Graph) readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error to open file: %w", err)
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
			return nil, fmt.Errorf("error to read file: %w", err)
		}
	}

	if err := file.Close(); err != nil {
		return []byte{}, fmt.Errorf("error to close file: %w", err)
	}

	return result, nil
}

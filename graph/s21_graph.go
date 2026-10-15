package graph

import (
	"fmt"
	"io"
	"os"
)

type Graph struct {
	size             int
	adjacency_matrix [][]int
}

// загрузка графа из файла в формате матрицы смежности
func (g *Graph) LoadGraphFromFile(filename string) {

}

// выгрузка графа в файл в формате dot (см. материалы)
func (g *Graph) ExportGraphToDot(filename string) {

}

// файл с расширением .gv или .dot

func readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error to read file: ", err)
		os.Exit(1)
	}
	defer file.Close()

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
			return nil, err
		}
	}

	return result, nil
}

package s21_graph_algorithms

import (
	graph "simplenavigator/graph"
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

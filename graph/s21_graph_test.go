package graph

import (
	"fmt"
	"strings"
	"testing"
)

func TestTest(t *testing.T) {
	x := 5
	if x != 5 {
		t.Errorf("Error!")
	}
	fmt.Println("Start testing")
}

func TestReadFile(t *testing.T) {
	var graph Graph
	result, err := graph.readFile("../test_data/file.txt")

	expectedResult := `11
0   29  20  21  16  31  100 12  4   31  18
29  0   15  29  28  40  72  21  29  41  12
20  15  0   15  14  25  81  9   23  27  13
21  29  15  0   4   12  92  12  25  13  25
16  28  14  4   0   16  94  9   20  16  22
31  40  25  12  16  0   95  24  36  3   37
100 72  81  92  94  95  0   90  101 99  84
12  21  9   12  9   24  90  0   15  25  13
4   29  23  25  20  36  101 15  0   35  18
31  41  27  13  16  3   99  25  35  0   38
18  12  13  25  22  37  84  13  18  38  0`

	if expectedResult != string(result) {
		t.Errorf("Get %s, want %s", result, expectedResult)
	}
	if err != nil {
		t.Errorf("Get error: %s", err)
	}
}

func TestReadFileErrorOpen(t *testing.T) {
	var graph Graph
	result, err := graph.readFile("../test_data/wrongFile.txt")

	expectedResult := ``
	expectedError := "error to open file: "

	if expectedResult != string(result) {
		t.Errorf("Get %s, want %s", result, expectedResult)
	}

	if !strings.Contains(strings.ToLower(err.Error()), expectedError) {
		t.Errorf("Expected error %q, but got: %v", expectedError, err)
	} else {
		if expectedResult != string(result) {
			t.Errorf("Get %s, want %s", result, expectedResult)
		}
	}
}

func TestLoadGraphFromFile(t *testing.T) {
	testCases := []struct {
		testName       string
		filename       string
		expectedError  string
		expectedSize   int
		expectedMatrix [][]int
	}{
		{
			testName:      "ErrReadFile",
			filename:      "../test_data/wrongFile.txt",
			expectedError: ErrReadFile.Error(),
		},
		{
			testName:      "ErrEmptyFile",
			filename:      "../test_data/emptyFile.txt",
			expectedError: ErrEmptyFile.Error(),
		},
		{
			testName:      "ErrBadRow",
			filename:      "../test_data/badrow.txt",
			expectedError: ErrBadRow.Error(),
		},
		{
			testName:      "ErrBadNumber",
			filename:      "../test_data/badnumber.txt",
			expectedError: "graph: bad number in line",
		},
		{
			testName:      "ErrNegativeEdge",
			filename:      "../test_data/negativeedge.txt",
			expectedError: ErrNegativeEdge.Error(),
		},
		{
			testName:      "OK",
			filename:      "../test_data/file.txt",
			expectedError: "",
			expectedSize:  11,
			expectedMatrix: [][]int{
				{0, 29, 20, 21, 16, 31, 100, 12, 4, 31, 18},
				{29, 0, 15, 29, 28, 40, 72, 21, 29, 41, 12},
				{20, 15, 0, 15, 14, 25, 81, 9, 23, 27, 13},
				{21, 29, 15, 0, 4, 12, 92, 12, 25, 13, 25},
				{16, 28, 14, 4, 0, 16, 94, 9, 20, 16, 22},
				{31, 40, 25, 12, 16, 0, 95, 24, 36, 3, 37},
				{100, 72, 81, 92, 94, 95, 0, 90, 101, 99, 84},
				{12, 21, 9, 12, 9, 24, 90, 0, 15, 25, 13},
				{4, 29, 23, 25, 20, 36, 101, 15, 0, 35, 18},
				{31, 41, 27, 13, 16, 3, 99, 25, 35, 0, 38},
				{18, 12, 13, 25, 22, 37, 84, 13, 18, 38, 0},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var graph Graph
			err := graph.LoadGraphFromFile(tc.filename)

			if tc.expectedError == "" {
				if graph.size != tc.expectedSize {
					t.Errorf("Expected size %d, but got: %d", tc.expectedSize, graph.size)
					return
				}

				for i := 0; i < graph.size; i++ {
					for j := 0; j < graph.size; j++ {
						if graph.matrix[i][j] != tc.expectedMatrix[i][j] {
							t.Errorf("Expected number %d from matrix[%d][%d], but got: %d", tc.expectedMatrix[i][j], i, j, graph.matrix[i][j])
						}
					}
				}

				if graph.size != tc.expectedSize {
					t.Errorf("Expected size %d, but got: %d", tc.expectedSize, graph.size)
				}

			} else {
				if !strings.Contains(strings.ToLower(err.Error()), tc.expectedError) {
					t.Errorf("Expected error %q, but got: %v", tc.expectedError, err)
				}
			}
		})
	}
}

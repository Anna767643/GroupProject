package graph

import (
	"errors"
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
	result, err := graph.readFile("../test_data/weighted.txt")

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
	expectedError := "graph: error to open file: "

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
			testName:      "Error to read file",
			filename:      "../test_data/wrongFile.txt",
			expectedError: "graph: error to open file: ",
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
			filename:      "../test_data/weighted.txt",
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

func TestNeighbors(t *testing.T) {
	testCases := []struct {
		testName          string
		fileName          string
		veritace          int
		expectedNeighbors []int
		expectedError     error
	}{
		{
			testName:          "OK, unweighted",
			fileName:          "../test_data/unweighted.txt",
			veritace:          2,
			expectedNeighbors: []int{1, 4, 5},
			expectedError:     nil,
		},
		{
			testName:          "OK, weighted",
			fileName:          "../test_data/weighted.txt",
			veritace:          3,
			expectedNeighbors: []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11},
			expectedError:     nil,
		},
		{
			testName:          "graph: vertice must be more zero",
			fileName:          "../test_data/weighted.txt",
			veritace:          0,
			expectedNeighbors: []int{},
			expectedError:     ErrVertice,
		},
		{
			testName:          "graph: vertice must be less graph sizeo",
			fileName:          "../test_data/weighted.txt",
			veritace:          12,
			expectedNeighbors: []int{},
			expectedError:     ErrVertice,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var graph Graph
			err := graph.LoadGraphFromFile(tc.fileName)
			if err != nil {
				t.Errorf("graph.LoadGraphFromFile() get error %v, but want nil", err)
				return
			}
			fmt.Println(graph.Size())
			neighbors, err := graph.Neighbors(tc.veritace)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("graph.Neighbors() get error %v, but want %v", err, tc.expectedError)
				return
			}

			if len(neighbors) != len(tc.expectedNeighbors) {
				t.Errorf("len neighbors %d, but want %d", len(neighbors), len(tc.expectedNeighbors))
				return
			}

			for i := range tc.expectedNeighbors {
				if neighbors[i] != tc.expectedNeighbors[i] {
					t.Errorf("graph.Neighbors() get neighbors %d, but want %d", neighbors[i], tc.expectedNeighbors[i])
					return
				}
			}
		})
	}
}

func TestWeight(t *testing.T) {
	testCases := []struct {
		testName       string
		fileName       string
		verticeFirst   int
		verticeSecond  int
		expectedWeight int
	}{
		{
			testName:       "unweighted 2 - 4",
			fileName:       "../test_data/unweighted.txt",
			verticeFirst:   2,
			verticeSecond:  4,
			expectedWeight: 1,
		},
		{
			testName:       "unweighted 2 - 3",
			fileName:       "../test_data/unweighted.txt",
			verticeFirst:   2,
			verticeSecond:  3,
			expectedWeight: 0,
		},
		{
			testName:       "weighted 2 - 3",
			fileName:       "../test_data/weighted.txt",
			verticeFirst:   2,
			verticeSecond:  3,
			expectedWeight: 15,
		},
		{
			testName:       "weighted 3 - 3",
			fileName:       "../test_data/weighted.txt",
			verticeFirst:   3,
			verticeSecond:  3,
			expectedWeight: 0,
		},
		{
			testName:       "weighted_directed 1 - 11",
			fileName:       "../test_data/weighted_directed.txt",
			verticeFirst:   1,
			verticeSecond:  11,
			expectedWeight: 55,
		},
		{
			testName:       "weighted_directed 11 - 1",
			fileName:       "../test_data/weighted_directed.txt",
			verticeFirst:   11,
			verticeSecond:  1,
			expectedWeight: 18,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var graph Graph
			err := graph.LoadGraphFromFile(tc.fileName)
			if err != nil {
				t.Errorf("graph.LoadGraphFromFile() get error %v, but want nil", err)
				return
			}
			weight := graph.Weight(tc.verticeFirst, tc.verticeSecond)

			if weight != tc.expectedWeight {
				t.Errorf("graph.Weight() get %d, but want %d", weight, tc.expectedWeight)
				return
			}

		})
	}
}

func TestIsDirected(t *testing.T) {
	testCases := []struct {
		testName       string
		fileName       string
		expectedResult bool
	}{
		{
			testName:       "unweighted",
			fileName:       "../test_data/unweighted.txt",
			expectedResult: false,
		},
		{
			testName:       "weighted",
			fileName:       "../test_data/weighted.txt",
			expectedResult: false,
		},
		{
			testName:       "weighted_directed",
			fileName:       "../test_data/weighted_directed.txt",
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var graph Graph
			err := graph.LoadGraphFromFile(tc.fileName)
			if err != nil {
				t.Errorf("graph.LoadGraphFromFile() get error %v, but want nil", err)
				return
			}
			result := graph.IsDirected()

			if result != tc.expectedResult {
				t.Errorf("graph.IsDirected() get %t, but want %t", result, tc.expectedResult)
				return
			}

		})
	}
}

func TestExportGraphToDot(t *testing.T) {
	testCases := []struct {
		testName       string
		fileNameLoad   string
		fileNameExport string
		expectedResult string
	}{
		{
			testName:       "unweighted",
			fileNameLoad:   "../test_data/unweighted.txt",
			fileNameExport: "../test_data/files_dot/unweighted.dot",
			expectedResult: `graph graphname {
     1;
     2;
     3;
     4;
     5;
     1 -- 2;
     1 -- 3;
     2 -- 4;
     2 -- 5;
     3 -- 4;
     4 -- 5;
}`,
		},
		{
			testName:       "weighted",
			fileNameLoad:   "../test_data/weighted.txt",
			fileNameExport: "../test_data/files_dot/weighted.dot",
			expectedResult: `graph graphname {
     1;
     2;
     3;
     4;
     5;
     6;
     7;
     8;
     9;
     10;
     11;
     1 -- 2 [weight=29];
     1 -- 3 [weight=20];
     1 -- 4 [weight=21];
     1 -- 5 [weight=16];
     1 -- 6 [weight=31];
     1 -- 7 [weight=100];
     1 -- 8 [weight=12];
     1 -- 9 [weight=4];
     1 -- 10 [weight=31];
     1 -- 11 [weight=18];
     2 -- 3 [weight=15];
     2 -- 4 [weight=29];
     2 -- 5 [weight=28];
     2 -- 6 [weight=40];
     2 -- 7 [weight=72];
     2 -- 8 [weight=21];
     2 -- 9 [weight=29];
     2 -- 10 [weight=41];
     2 -- 11 [weight=12];
     3 -- 4 [weight=15];
     3 -- 5 [weight=14];
     3 -- 6 [weight=25];
     3 -- 7 [weight=81];
     3 -- 8 [weight=9];
     3 -- 9 [weight=23];
     3 -- 10 [weight=27];
     3 -- 11 [weight=13];
     4 -- 5 [weight=4];
     4 -- 6 [weight=12];
     4 -- 7 [weight=92];
     4 -- 8 [weight=12];
     4 -- 9 [weight=25];
     4 -- 10 [weight=13];
     4 -- 11 [weight=25];
     5 -- 6 [weight=16];
     5 -- 7 [weight=94];
     5 -- 8 [weight=9];
     5 -- 9 [weight=20];
     5 -- 10 [weight=16];
     5 -- 11 [weight=22];
     6 -- 7 [weight=95];
     6 -- 8 [weight=24];
     6 -- 9 [weight=36];
     6 -- 10 [weight=3];
     6 -- 11 [weight=37];
     7 -- 8 [weight=90];
     7 -- 9 [weight=101];
     7 -- 10 [weight=99];
     7 -- 11 [weight=84];
     8 -- 9 [weight=15];
     8 -- 10 [weight=25];
     8 -- 11 [weight=13];
     9 -- 10 [weight=35];
     9 -- 11 [weight=18];
     10 -- 11 [weight=38];
}`,
		},
		{
			testName:       "weighted_directed",
			fileNameLoad:   "../test_data/weighted_directed.txt",
			fileNameExport: "../test_data/files_dot/weighted_directed.dot",
			expectedResult: `digraph graphname {
     1;
     2;
     3;
     4;
     5;
     6;
     7;
     8;
     9;
     10;
     11;
     1 -> 2 [weight=29];
     1 -> 3 [weight=20];
     1 -> 4 [weight=21];
     1 -> 5 [weight=16];
     1 -> 6 [weight=31];
     1 -> 7 [weight=100];
     1 -> 8 [weight=12];
     1 -> 9 [weight=4];
     1 -> 10 [weight=31];
     1 -> 11 [weight=55];
     2 -> 1 [weight=29];
     2 -> 3 [weight=15];
     2 -> 4 [weight=29];
     2 -> 5 [weight=28];
     2 -> 6 [weight=40];
     2 -> 7 [weight=72];
     2 -> 8 [weight=21];
     2 -> 9 [weight=29];
     2 -> 10 [weight=41];
     2 -> 11 [weight=12];
     3 -> 1 [weight=20];
     3 -> 2 [weight=15];
     3 -> 4 [weight=15];
     3 -> 5 [weight=14];
     3 -> 6 [weight=25];
     3 -> 7 [weight=81];
     3 -> 8 [weight=9];
     3 -> 9 [weight=23];
     3 -> 10 [weight=27];
     3 -> 11 [weight=13];
     4 -> 1 [weight=21];
     4 -> 2 [weight=29];
     4 -> 3 [weight=15];
     4 -> 5 [weight=4];
     4 -> 6 [weight=12];
     4 -> 7 [weight=11];
     4 -> 8 [weight=12];
     4 -> 9 [weight=25];
     4 -> 10 [weight=13];
     4 -> 11 [weight=25];
     5 -> 1 [weight=16];
     5 -> 2 [weight=28];
     5 -> 3 [weight=14];
     5 -> 4 [weight=4];
     5 -> 6 [weight=16];
     5 -> 7 [weight=94];
     5 -> 8 [weight=9];
     5 -> 9 [weight=20];
     5 -> 10 [weight=16];
     5 -> 11 [weight=22];
     6 -> 1 [weight=31];
     6 -> 2 [weight=40];
     6 -> 3 [weight=25];
     6 -> 4 [weight=12];
     6 -> 5 [weight=16];
     6 -> 7 [weight=95];
     6 -> 8 [weight=24];
     6 -> 9 [weight=36];
     6 -> 10 [weight=3];
     6 -> 11 [weight=37];
     7 -> 1 [weight=100];
     7 -> 2 [weight=72];
     7 -> 3 [weight=81];
     7 -> 4 [weight=92];
     7 -> 5 [weight=94];
     7 -> 6 [weight=95];
     7 -> 8 [weight=90];
     7 -> 9 [weight=101];
     7 -> 10 [weight=99];
     7 -> 11 [weight=84];
     8 -> 1 [weight=12];
     8 -> 2 [weight=21];
     8 -> 3 [weight=9];
     8 -> 4 [weight=12];
     8 -> 5 [weight=8];
     8 -> 6 [weight=24];
     8 -> 7 [weight=90];
     8 -> 9 [weight=15];
     8 -> 10 [weight=25];
     8 -> 11 [weight=13];
     9 -> 1 [weight=4];
     9 -> 2 [weight=12];
     9 -> 3 [weight=23];
     9 -> 4 [weight=25];
     9 -> 5 [weight=20];
     9 -> 6 [weight=36];
     9 -> 7 [weight=101];
     9 -> 8 [weight=15];
     9 -> 10 [weight=35];
     9 -> 11 [weight=18];
     10 -> 1 [weight=31];
     10 -> 2 [weight=41];
     10 -> 3 [weight=27];
     10 -> 4 [weight=13];
     10 -> 5 [weight=16];
     10 -> 6 [weight=3];
     10 -> 7 [weight=99];
     10 -> 8 [weight=25];
     10 -> 9 [weight=35];
     10 -> 11 [weight=38];
     11 -> 1 [weight=18];
     11 -> 2 [weight=12];
     11 -> 3 [weight=13];
     11 -> 4 [weight=25];
     11 -> 5 [weight=22];
     11 -> 6 [weight=37];
     11 -> 7 [weight=84];
     11 -> 8 [weight=13];
     11 -> 9 [weight=18];
     11 -> 10 [weight=38];
}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var graph Graph
			err := graph.LoadGraphFromFile(tc.fileNameLoad)
			if err != nil {
				t.Errorf("graph.LoadGraphFromFile() get error %v, but want nil", err)
				return
			}

			err = graph.ExportGraphToDot(tc.fileNameExport)

			if err != nil {
				t.Errorf("graph.ExportGraphToDot() get error %v, but want nil", err)
				return
			}

			result, err := graph.readFile(tc.fileNameExport)
			if err != nil {
				t.Errorf("graph.readFile() get error %v, but want nil", err)
				return
			}

			if string(result) != tc.expectedResult {
				t.Errorf("graph.ExportGraphToDot() get %s, but want %s", result, tc.expectedResult)
				return
			}

		})
	}
}

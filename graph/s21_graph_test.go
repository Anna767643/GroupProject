package graph

import (
	"fmt"
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
	result, err := readFile("file.dot")

	expextResult := `graph graphname {
    a;
    b;
    c;
    d;
    a -- b;
    b -- c;
    b -- d;
}`

	if expextResult != string(result) {
		t.Errorf("Get %s, want %s", result, expextResult)
	}
	if err != nil {
		t.Errorf("Get error: %s", err)
	}
	fmt.Println("Start testing")
}

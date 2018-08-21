package graph

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestWriteto(t *testing.T) {
	adjList := [][]int{
		{3, 4},
		{0, 2},
		{3},
		{2, 4},
		{0},
	}
	w := bytes.NewBuffer(nil)
	err := Writeto(w, adjList)
	if err != nil {
		t.Error(err)
	}

	expected := "5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n"
	if expected != w.String() {
		t.Logf("expected: %s\n", expected)
		t.Errorf("found: %s\n", w.String())
	}
}

func TestReadFrom(t *testing.T) {
	expected := [][]int{
		{3, 4},
		{0, 2},
		{3},
		{2, 4},
		{0},
	}
	r := strings.NewReader("5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n")
	var adjList [][]int
	if err := Readfrom(r, &adjList); err != nil {
		fmt.Println(err)
	}
	if reflect.DeepEqual(expected, adjList) == false {
		t.Log("expected:", expected, "\n")
		t.Error("found:", adjList, "\n")
	}
}

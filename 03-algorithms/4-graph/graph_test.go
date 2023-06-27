package graph

import (
	"fmt"
	"testing"
)

func TestList_MinSpanTree(t *testing.T) {
	l := NewList()
	l.AddNodes(1, 2, 3, 4, 5, 6)
	l.AddEdge(1, 2, 10)
	l.AddEdge(1, 3, 15)
	l.AddEdge(1, 5, 20)
	l.AddEdge(2, 3, 20)
	l.AddEdge(2, 5, 35)
	l.AddEdge(2, 4, 10)
	l.AddEdge(3, 4, 20)
	l.AddEdge(3, 6, 50)
	l.AddEdge(4, 6, 30)

	mst := l.MinSpanTree()
	fmt.Println(mst)
	fmt.Println(mst.Weight())

	want := 85
	got := mst.Weight()
	if got != want {
		t.Fatalf("длина дерева %d, а ожидалось %d", got, want)
	}
}

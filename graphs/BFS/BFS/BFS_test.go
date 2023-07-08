package BFS

import (
	"algorithms/graphs/Graph"
	"testing"
)

func TestDetourBFS(t *testing.T) {
	g := Graph.CreateGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	expected := "1 2 3 4 "
	actual := BFS(g, 1)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	g.AddEdge(3, 5)
	g.AddEdge(1, 10)
	expected = "1 2 3 10 4 5 "
	actual = BFS(g, 1)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

package Graph

import (
	"testing"
)

func TestCreateGraph(t *testing.T) {
	g := CreateGraph()
	if g == nil {
		t.Error("Graph is nil")
	}
	if g.Edges == nil {
		t.Error("Edges is nil")
	}
}

func TestAddEdge(t *testing.T) {
	g := CreateGraph()
	g.AddEdge(1, 2)
	if g.Edges[1] == nil {
		t.Error("Edges[1] is nil")
	}
	if g.Edges[1][0] != 2 {
		t.Error("Edges[1][0] != 2")
	}
}

package Graph

type Graph struct {
	Edges map[int][]int
}

func CreateGraph() *Graph {
	return &Graph{Edges: make(map[int][]int)}
}

func (g *Graph) AddEdge(key, value int) {
	if _, have := g.Edges[key]; have {
		g.Edges[key] = append(g.Edges[key], value)
	} else {
		g.Edges[key] = []int{value}
	}
}

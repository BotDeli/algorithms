package BFS

import (
	"algorithms/graphs/Graph"
	"algorithms/queue"
	"fmt"
)

// Алгоритм нахождения минимального количества шагов или расстояния
// между начальной и конечной точками в графах
// алгоритм поиска в ширину

func BFS(graph *Graph.Graph, start int) (allEdges string) {
	queueEdges := queue.CreateQueue()
	visited := make(map[int]bool)
	queueEdges.Enqueue(start)
	for queueEdges.NotEmpty() {
		edge := queueEdges.Dequeue().(int)
		allEdges += fmt.Sprintf("%d ", edge)
		visited[edge] = true
		for _, nextEdge := range graph.Edges[edge] {
			if !visited[nextEdge] {
				queueEdges.Enqueue(nextEdge)
			}
		}
	}
	return allEdges
}

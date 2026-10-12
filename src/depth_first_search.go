package src

import "log/slog"


func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) AddAdjList(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
}

func (g *Graph) DepthFirstSearch(v int, visited map[int]bool) {
	slog.Info("Visiting node", "node", v)
	visited[v] = true
	for _, neighbor := range g.adjList[v] {
		if !visited[neighbor] {
			g.DepthFirstSearch(neighbor, visited)
		}
	}
}
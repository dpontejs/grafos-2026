package algoritmos

import "github.com/PauloFH/grafos-2026/internal/grafo"

func SaoAdjacentes(g *grafo.Grafo, a, b string) bool {
	return g.SaoAdjacentes(a, b)
}

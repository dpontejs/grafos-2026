package conversoes

import "github.com/PauloFH/grafos-2026/internal/grafo"

func ListaParaMatriz(g *grafo.Grafo) {
	n := len(g.Vertices)
	idx := make(map[string]int, n)
	for i, v := range g.Vertices {
		idx[v] = i
	}

	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}

	for _, v := range g.Vertices {
		i := idx[v]
		for _, viz := range g.ListaAdj[v] {
			j := idx[viz]
			m[i][j] = 1
		}
	}

	g.MatrizAdj = m
}

func MatrizParaLista(g *grafo.Grafo) {
	for _, v := range g.Vertices {
		g.ListaAdj[v] = make([]string, 0)
	}

	for i, v := range g.Vertices {
		for j, viz := range g.Vertices {
			if g.MatrizAdj[i][j] == 1 {
				g.ListaAdj[v] = append(g.ListaAdj[v], viz)
			}
		}
	}
}

package search

import (
	"strconv"

	"github.com/daniel-orlov/go-data-structures/graph"
	ll "github.com/daniel-orlov/go-data-structures/linked"
	"github.com/daniel-orlov/go-data-structures/stack"
)

// DepthFirstSearch is an algorithm for performing search in a graph.
// This is an iterative implementation of it.
func (g *graph.Graph) DepthFirstSearch(vertexID int) {
	zeroth := ll.NewElemWithID("0", g.Vertices[0])
	lifo := stack.NewStack("dfs", zeroth)

	discovered := make(map[int]bool) // TODO: Add size of Graph here once g.Size() implemented

	for !lifo.IsEmpty() {
		el := lifo.Pop()

		if !discovered[el.ID] {
			discovered[el.ID] = true
			vtx := el.Data
			neighbours := vtx.Vertices
			for k, v := range neighbours {
				keyName := strconv.Itoa(k)
				vertex := ll.NewElem(keyName, v)
				lifo.Push(vertex)
			}
		}
	}
}

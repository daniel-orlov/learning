package search

import (
	"github.com/daniel-orlov/go-data-structures/graph"
	ll "github.com/daniel-orlov/go-data-structures/linked"
	"github.com/daniel-orlov/go-data-structures/stack"
)

/*procedure DFS_iterative(G, v) is
  let S be a stack
  S.push(v)
  while S is not empty do
      v = S.pop()
      if v is not labeled as discovered then
          label v as discovered
          for all edges from v to w in G.adjacentEdges(v) do
              S.push(w)
*/

func (g *graph.Graph) DepthFirstSearch(vertexID int) {
	zeroth := ll.NewElem(0, g.Vertices[0]) // this is not gonna work, gotta fix NewElem to work with ID first
	dfsStack := stack.NewStack("dfs", zeroth)

	discovered := make(map[int]bool) // TODO: Add size of Graph here once g.Size() implemented

	for !dfsStack.IsEmpty() {
		el := dfsStack.Pop()
		if !discovered[el.ID] {
			discovered[el.ID] = true
			vtx := el.Data
			neighbours := vtx.Vertices
			for k, v := range neighbours {
				vertex := ll.NewElem(k, v) // same as the 1st line of this func
				dfsStack.Push(vertex)
			}
		}
	}
}

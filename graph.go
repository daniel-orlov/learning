package main

import (
	"errors"
	"fmt"
)

// Vertex for the graph
type vertex struct {
	ID       int
	Vertices map[int]*vertex
}

// newVertex is a type vertex constructor
func newVertex(id int) *vertex {
	return &vertex{
		ID:       id,
		Vertices: map[int]*vertex{},
	}
}

// Graph of vertices with a name and an optional direction
type Graph struct {
	Directed bool
	Name     string
	Vertices map[int]*vertex
}

// NewGraph is a type Graph constructor
func NewGraph(name string, isDirected bool) *Graph {
	return &Graph{
		Directed: isDirected,
		Name:     name,
		Vertices: map[int]*vertex{},
	}
}

// AddVertex creates a new Vertex and adds it to the graph
func (g *Graph) AddVertex(id int) {
	vert := newVertex(id)
	g.Vertices[id] = vert
}

// AddEdge creates a connection between two vertices in the graph
func (g *Graph) AddEdge(firstVertexID, secondVertexID int) error {
	vert1 := g.Vertices[firstVertexID]
	vert2 := g.Vertices[secondVertexID]

	if vert1 == nil || vert2 == nil {
		return fmt.Errorf("%wmissing some vertices", ErrCreateEdge)
	}

	if _, ok := vert1.Vertices[vert2.Id]; ok {
		return fmt.Errorf("%wedge already exists", ErrCreateEdge)
	}

	vert1.Vertices[vert2.ID] = vert2
	if !g.Directed && vert1.ID != vert2.ID {
		vert2.Vertices[vert1.ID] = vert1
	}

	return nil
}

// ErrCreateEdge indicates an edge cannot be created
var ErrCreateEdge = errors.New("enable to create an edge: ")

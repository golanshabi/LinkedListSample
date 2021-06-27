package main

import (
	"errors"
	"fmt"

	linkedList "github.com/golanshabi/LinkedList-golang"
)

var (
	errInvalidVertexID = errors.New("error: cannot find vertex number")
	errSameEdgeTwice   = errors.New("error: cant create the same edge twice")
	errConvertToVertex = errors.New("error: problem converting to vertex")
)

const (
	ZERO  = 0
	ONE   = 1
	TWO   = 2
	THREE = 3
	FOUR  = 4
	FIVE  = 5
	SIX   = 6
)

func main() {
	g, err := createGraph()
	if err != nil {
		return
	}

	getDistances(g)
}

func getDistances(g Graph) {
	err := g.ShortestPathsBFS(ZERO)
	if err != nil {
		return
	}

	err = g.ShortestPathsBFS(ONE)
	if err != nil {
		return
	}

	err = g.ShortestPathsBFS(TWO)
	if err != nil {
		return
	}

	err = g.ShortestPathsBFS(FOUR)
	if err != nil {
		return
	}

	err = g.ShortestPathsBFS(SIX)
	if err != nil {
		return
	}
}

type Vertex struct {
	neighbors []int
	distances map[int]int
	id        int
}

type Graph struct{ vertices []Vertex }

func (g *Graph) AddVertex() {
	g.vertices = append(g.vertices, Vertex{
		neighbors: []int{},
		distances: make(map[int]int),
		id:        len(g.vertices),
	})
}

func (g *Graph) AddEdge(first, second int) error {
	if first < 0 || first >= len(g.vertices) {
		return fmt.Errorf("%w: %d", errInvalidVertexID, first)
	}

	if second < 0 || second >= len(g.vertices) {
		return fmt.Errorf("%w: %d", errInvalidVertexID, second)
	}

	if contains(g.vertices[first].neighbors, second) || contains(g.vertices[second].neighbors, first) {
		return errSameEdgeTwice
	}

	g.vertices[first].neighbors = append(g.vertices[first].neighbors, second)

	g.vertices[second].neighbors = append(g.vertices[second].neighbors, first)

	return nil
}

func (g *Graph) ShortestPathsBFS(source int) error {
	if source < 0 || source > len(g.vertices) {
		return fmt.Errorf("%w %d", errInvalidVertexID, source)
	}

	previous := make(map[int]int)

	v := g.vertices[source]
	for i := 0; i < len(g.vertices); i++ {
		v.distances[i] = -1
		previous[i] = -1
	}

	v.distances[source] = 0
	queue := linkedList.NewLinkedList()

	var curVertex Vertex

	var tmp interface{}

	var err error

	var ok bool

	queue.PushBack(v)

	for queue.Len() > 0 {
		tmp, err = queue.PopFront()
		if err != nil {
			return fmt.Errorf("error: failed popping, %w", err)
		}

		curVertex, ok = tmp.(Vertex)
		if !ok {
			return errConvertToVertex
		}

		for _, u := range curVertex.neighbors {
			if v.distances[u] == -1 {
				queue.PushBack(g.vertices[u])

				v.distances[u] = 1 + v.distances[curVertex.id]
			}
		}
	}

	return nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func createGraph() (Graph, error) {
	var g Graph
	for i := 0; i < 7; i++ {
		g.AddVertex()
	}

	err := g.AddEdge(ZERO, ONE)
	if err != nil {
		return Graph{}, err
	}

	err = g.AddEdge(ONE, TWO)
	if err != nil {
		return Graph{}, err
	}

	err = g.AddEdge(ONE, THREE)
	if err != nil {
		return Graph{}, err
	}

	err = g.AddEdge(ONE, FOUR)
	if err != nil {
		return Graph{}, err
	}

	err = g.AddEdge(TWO, THREE)
	if err != nil {
		return Graph{}, err
	}

	err = g.AddEdge(THREE, FOUR)
	if err != nil {
		return Graph{}, err
	}

	err = g.AddEdge(FOUR, FIVE)
	if err != nil {
		return Graph{}, err
	}

	return g, err
}

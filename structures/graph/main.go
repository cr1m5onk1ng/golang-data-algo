package main

import "fmt"

const N = 10

type Graph struct {
	vertices [N]*Vertex
}


func (g *Graph) AddVertex(k int) {
	if k >= len(g.vertices) {
		err := fmt.Errorf("Invalid node key %v", k)
		fmt.Println(err.Error())
	}else if g.vertices[k] != nil {
		err := fmt.Errorf("Node %v already present", k)
		fmt.Println(err.Error())
	} else {
		g.vertices[k] = k
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	if k >= len(g.vertices) {
		return nil
	}
	return g.vertices[k]
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge from %v to %v", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Edge from %v to %v already exists", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func contains(adjecencyList []*Vertex, k int) bool {
	for _, v := range adjecencyList {
		if v.key == k {
			return true
		}
	}
	return false
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (v *Vertex) addAdjecent(k int) {
	v.adjacent = append(v.adjacent, &Vertex{key: k})
}

func (g *Graph) DFS(target int) bool {
	return true
}

func (g *Graph) BFS(target int) []*Vertex {
	path := []*Vertex

}

func main() {
	g = &Graph{}
	
}

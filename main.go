package main

import(
	"fmt"
)

// graph structure
type Graph struct {
	vertices []*Vertex
}
// vertex structure
type Vertex struct {
	key int
	adjacent []*Vertex
}
// add edge
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge %d --> %d", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Edge %d --> %d already exists", from, to)
		fmt.Println(err.Error())
	} else {
	// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}
// add vertex
func (g *Graph) AddVertex(k int) {
	if !contains(g.vertices, k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	} else {
		fmt.Println("Vertex", k, "already exists")
	}
}

// getVertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}
// contains
func contains(v []*Vertex, k int) bool {
	for _, vertex := range v {
		if vertex.key == k {
			return true
		}
	}
	return false
}
// Print 
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex %d: ", v.key)
		for _, vv := range v.adjacent {
			fmt.Printf("%d ", vv.key)
		}
		fmt.Println()
	}
}
func main(){
	test := Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.Print()
}
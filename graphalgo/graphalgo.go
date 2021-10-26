package graphalgo

import (
	"log"
)

type Graph struct {
	verts []*Vertex
}

type Vertex struct {
	key int
	rel []*Vertex
}

// create node
func (g *Graph) addVert(k int) {
	if contain(g.verts, k) {
		log.Fatalf("duplicate key")
	}

	g.verts = append(g.verts, &Vertex{key: k})
}

func contain(g []*Vertex, k int) bool {
	for _, v := range g {
		if k == v.key {
			return true
		}
	}
	return false
}

// create rel
func (g *Graph) addRel(from, to int) {
	if g.check(from) && g.check(to) {
		log.Fatal("vert not exist")
	}
	s:= g.verts[from]
	e:= g.verts[to]

	s.rel = append(s.rel, e)

}

// check node
func (g *Graph) check(key int) bool {
	for _, k := range g.verts {
		if k.key == key {
			return true
		}
	}
	return false
}
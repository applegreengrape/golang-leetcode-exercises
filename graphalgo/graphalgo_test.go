package graphalgo

import (
	"log"
	"testing"
)

func TestAddNode(t *testing.T) {
	g := &Graph{}

	g.addVert(0)
	g.addVert(1)
	g.addVert(2)

	check := contain(g.verts, 1)

	log.Println(check)
}

func TestAddRel(t *testing.T) {
	g := &Graph{}

	g.addVert(0)
	g.addVert(1)

	g.addRel(0, 1)

	log.Println(g.verts, len(g.verts))
	for _, node := range g.verts {
		log.Println(node.key, node.rel)
	}

	g.addRel(2, 1)
}

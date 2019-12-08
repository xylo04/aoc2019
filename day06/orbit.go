package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/goombaio/dag"
	"io/ioutil"
	"strings"
)

type OrbitMap struct {
	graph  *dag.DAG
	bodies mapset.Set
}

func createOrbits(orbitStr string) OrbitMap {
	graph := dag.NewDAG()
	bodies := mapset.NewSet()
	lines := strings.Split(orbitStr, "\n")
	for _, rel := range lines {
		if rel == "" {
			continue
		}
		bod := strings.Split(rel, ")")
		bodies.Add(bod[0])
		bodies.Add(bod[1])
		orbited := getOrCreate(bod[0], graph)
		orbiter := getOrCreate(bod[1], graph)
		_ = graph.AddVertex(orbited)
		_ = graph.AddVertex(orbiter)
		_ = graph.AddEdge(orbited, orbiter)
	}
	return OrbitMap{graph, bodies}
}

func getOrCreate(name string, orbits *dag.DAG) *dag.Vertex {
	vertex, _ := orbits.GetVertex(name)
	if vertex == nil {
		return dag.NewVertex(name, nil)
	}
	return vertex
}

func checksumOrbits(orbitMap OrbitMap) int {
	checksum := 0
	for body := range orbitMap.bodies.Iterator().C {
		checksum += stepsToCOM(body.(string), orbitMap)
	}
	return checksum
}

func stepsToCOM(body string, orbitMap OrbitMap) int {
	if body == "COM" {
		return 0
	} else {
		vertex, _ := orbitMap.graph.GetVertex(body)
		var parentVertex = vertex.Parents.Values()[0].(*dag.Vertex)
		return 1 + stepsToCOM(parentVertex.ID, orbitMap)
	}
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	checksum := checksumOrbits(createOrbits(string(content)))
	fmt.Print(checksum)
}

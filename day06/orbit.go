package main

import (
	"fmt"
	"github.com/twmb/algoimpl/go/graph"
	"io/ioutil"
	"strings"
)

type OrbitMap struct {
	g      *graph.Graph
	bodies map[string]graph.Node
}

func createOrbits(orbitStr string) OrbitMap {
	g := graph.New(graph.Undirected)
	bodies := make(map[string]graph.Node, 0)
	orbitMap := OrbitMap{g, bodies}

	lines := strings.Split(orbitStr, "\n")
	for _, rel := range lines {
		if rel == "" {
			continue
		}
		bod := strings.Split(rel, ")")
		node0 := findOrCreate(bod[0], orbitMap)
		node1 := findOrCreate(bod[1], orbitMap)
		_ = g.MakeEdgeWeight(node0, node1, 1)
	}
	return orbitMap
}

func findOrCreate(name string, orbitMap OrbitMap) graph.Node {
	node, present := orbitMap.bodies[name]
	if !present {
		node = orbitMap.g.MakeNode()
		orbitMap.bodies[name] = node
	}
	return node
}

func checksumOrbits(orbitMap OrbitMap) int {
	checksum := 0
	for _, node := range orbitMap.bodies {
		pathsToAll := orbitMap.g.DijkstraSearch(node)
		checksum += findDistanceToTarget("COM", pathsToAll, orbitMap)
	}
	return checksum
}

func findDistanceToTarget(name string, paths []graph.Path, orbitMap OrbitMap) int {
	targetNode := findOrCreate(name, orbitMap)
	for _, path := range paths {
		edges := len(path.Path)
		if edges > 0 && path.Path[edges-1].End == targetNode {
			return path.Weight
		}
	}
	// assume this is target
	return 0
}

func findDistanceToSanta(orbitMap OrbitMap) interface{} {
	youNode := findOrCreate("YOU", orbitMap)
	paths := orbitMap.g.DijkstraSearch(youNode)
	return findDistanceToTarget("SAN", paths, orbitMap) - 2
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	distance := findDistanceToSanta(createOrbits(string(content)))
	fmt.Print(distance)
}

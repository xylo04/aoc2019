package main

import (
	"fmt"
	"github.com/juliangruber/go-intersect"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Loc struct {
	x int
	y int
}

func closestCrossing(input string) int {
	wireSpecs := strings.Split(input, "\n")
	log.Printf("Found %d wires, building\n", len(wireSpecs))
	wireSet0, dist0 := buildWire(wireSpecs[0])
	wireSet1, dist1 := buildWire(wireSpecs[1])

	log.Printf("Built wires with lengths %d and %d, finding intersections\n", len(wireSet0), len(wireSet1))
	var intersections = intersect.Hash(wireSet0, wireSet1)

	log.Printf("Found %d intersections", len(intersections))
	minDist := math.MaxInt32
	for _, intersection := range intersections {
		i, _ := intersection.(Loc)
		dist := dist0[i] + dist1[i]
		if dist < minDist {
			minDist = dist
		}
	}
	return minDist
}

func buildWire(wireSpec string) ([]Loc, map[Loc]int) {
	x, y, d := 0, 0, 0
	var wireSet = make([]Loc, 0)
	dist := map[Loc]int{}
	segments := strings.Split(wireSpec, ",")
	for _, seg := range segments {
		dir := seg[:1]
		mag, _ := strconv.Atoi(seg[1:])
		for i := 0; i < mag; i++ {
			switch dir {
			case "U":
				y += 1
			case "R":
				x += 1
			case "D":
				y -= 1
			case "L":
				x -= 1
			}
			loc := Loc{x, y}
			wireSet = append(wireSet, loc)
			d++
			dist[loc] = d
		}
	}
	return wireSet, dist
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	fmt.Println(closestCrossing(string(content)))
}

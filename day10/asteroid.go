package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/xylo04/aoc2019/common"
	"math"
)

type point struct {
	x, y int
}

type asteroidField struct {
	asteroids  mapset.Set
	visibility map[point]int
}

func main() {
	lines, err := common.FileToLines("input.txt")
	if err != nil {
		panic(err)
	}
	var field *asteroidField = parseAsteroids(lines)
	m, p := findBestAsteroid(field)
	fmt.Printf("Can see %d other asteroids from %d\n", m, p)
	// 329 @ (25,31)
}

func parseAsteroids(lines []string) *asteroidField {
	field := asteroidField{
		mapset.NewSet(),
		make(map[point]int),
	}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == '#' {
				field.asteroids.Add(point{x, y})
			}
		}
	}
	return &field
}

func findBestAsteroid(field *asteroidField) (int, point) {
	maxVisibility := 0
	var maxVisibilityPoint point
	for p := range field.asteroids.Iterator().C {
		source := p.(point)
		angles := mapset.NewSet()
		for q := range field.asteroids.Iterator().C {
			dest := q.(point)
			if source == dest {
				continue
			}
			angle := angleFrom(source, dest)
			angles.Add(angle)
		}
		if angles.Cardinality() > maxVisibility {
			maxVisibility = angles.Cardinality()
			maxVisibilityPoint = source
		}
	}

	return maxVisibility, maxVisibilityPoint
}

func angleFrom(source point, dest point) float64 {
	return math.Atan2(float64(dest.y-source.y), float64(dest.x-source.x))
}

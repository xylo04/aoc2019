package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/xylo04/aoc2019/common"
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
	// 332 too high, 325 too low, not 328, not 330
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
		for q := range field.asteroids.Iterator().C {
			dest := q.(point)
			if source == dest {
				continue
			}
			if canSeeEachOther(source, dest, field.asteroids) {
				field.visibility[source]++
				if field.visibility[source] > maxVisibility {
					maxVisibility = field.visibility[source]
					maxVisibilityPoint = source
				}
			}
		}
	}

	return maxVisibility, maxVisibilityPoint
}

func canSeeEachOther(source point, dest point, asteroids mapset.Set) bool {
	if source.x == dest.x {
		return canSeeEachOtherVert(source, dest, asteroids)
	}
	// iterate along the x's and see if any y's are whole integers
	minX := min(source.x, dest.x)
	maxX := max(source.x, dest.x)
	slope := float64(dest.y-source.y) / float64(dest.x-source.x)
	b := float64(source.y) - (slope * float64(source.x))
	for x := minX + 1; x < maxX; x++ {
		y := (slope * float64(x)) + b
		if isWholeNumber(y) {
			// if x and y are whole integers, see if there's an asteroid there
			if asteroids.Contains(point{x, int(y)}) {
				return false
			}
		}
	}
	return true
}

func canSeeEachOtherVert(source point, dest point, asteroids mapset.Set) bool {
	// slope is known to be inf, x coords are the same
	// iterate along the y's and see if any x's are whole integers
	x := source.x
	minY := min(source.y, dest.y)
	maxY := max(source.y, dest.y)
	for y := minY + 1; y < maxY; y++ {
		if asteroids.Contains(point{x, y}) {
			return false
		}
	}
	return true
}

func isWholeNumber(n float64) bool {
	if int64(n*128)%128 == 0 {
		return true
	}
	return false
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

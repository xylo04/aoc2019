package main

import "testing"

func TestWires(t *testing.T) {
	testData := map[string]int{
		"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83":               610,
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7": 410,
	}
	for wires, expected := range testData {
		actual := closestCrossing(wires)
		if actual != expected {
			t.Errorf("Manhattan distance of closes crossing should be %d but was %d", expected, actual)
		}
	}
}

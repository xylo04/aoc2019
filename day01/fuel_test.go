package main

import "testing"

func TestSum(t *testing.T) {
	testData := map[int]int{
		12:     2,
		1969:   966,
		100756: 50346}
	for mass, expected := range testData {
		actual := fuel(mass)
		if actual != expected {
			t.Errorf("Mass %d should have fuel %d but it was %d", mass, expected, actual)
		}
	}
}

package main

import "testing"

func TestSum(t *testing.T) {
	testData := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583}
	for mass, expected := range testData {
		actual := fuel(mass)
		if actual != expected {
			t.Errorf("Mass %d should have fuel %d but it was %d", mass, expected, actual)
		}
	}
}

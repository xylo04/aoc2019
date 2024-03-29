package main

import "testing"

func TestOrbitChecksum(t *testing.T) {
	testData := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"
	const expected = 42

	actual := checksumOrbits(createOrbits(testData))
	if actual != expected {
		t.Errorf("Orbit checksum should be %d but was %d", expected, actual)
	}
}

func TestDistanceToSanta(t *testing.T) {
	testData := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN"
	const expected = 4

	actual := findDistanceToSanta(createOrbits(testData))
	if actual != expected {
		t.Errorf("Distance to Santa should be %d but was %d", expected, actual)
	}
}

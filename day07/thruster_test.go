package main

import "testing"

func testAmpSettings(program string, expectedThrust int, t *testing.T, expectedAmpSettings []int) {
	actualThrust, actualAmpSettings := findMaxThrust(program)

	if actualThrust != expectedThrust {
		t.Errorf("Thrust was expected to be %d but was %d", expectedThrust, actualThrust)
	}
	for i := 0; i < 5; i++ {
		if actualAmpSettings[i] != expectedAmpSettings[i] {
			t.Errorf("Amp settings were expected to be %d but was %d", expectedAmpSettings, actualAmpSettings)
			break
		}
	}
}

func TestAmplifier1(t *testing.T) {
	program := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	expectedAmpSettings := []int{4, 3, 2, 1, 0}
	expectedThrust := 43210

	testAmpSettings(program, expectedThrust, t, expectedAmpSettings)
}

func TestAmplifier2(t *testing.T) {
	program := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	expectedThrust := 54321
	expectedAmpSettings := []int{0, 1, 2, 3, 4}

	testAmpSettings(program, expectedThrust, t, expectedAmpSettings)
}

func TestAmplifier3(t *testing.T) {
	program := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	expectedThrust := 65210
	expectedAmpSettings := []int{1, 0, 4, 3, 2}

	testAmpSettings(program, expectedThrust, t, expectedAmpSettings)
}

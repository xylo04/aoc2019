package main

import "testing"

func testAmpSettings(program string, expectedThrust int, t *testing.T, expectedAmpSettings []int, rangeMin int, rangeMax int) {
	actualThrust, actualAmpSettings := findMaxThrust(program, rangeMin, rangeMax)

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

	testAmpSettings(program, expectedThrust, t, expectedAmpSettings, 0, 4)
}

func TestAmplifier2(t *testing.T) {
	program := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	expectedThrust := 54321
	expectedAmpSettings := []int{0, 1, 2, 3, 4}

	testAmpSettings(program, expectedThrust, t, expectedAmpSettings, 0, 4)
}

func TestAmplifier3(t *testing.T) {
	program := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	expectedThrust := 65210
	expectedAmpSettings := []int{1, 0, 4, 3, 2}

	testAmpSettings(program, expectedThrust, t, expectedAmpSettings, 0, 4)
}

func TestFeedback1(t *testing.T) {
	program := "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
	expectedThrust := 139629729
	expectedAmpSettings := []int{9, 8, 7, 6, 5}

	testAmpSettings(program, expectedThrust, t, expectedAmpSettings, 5, 9)
}

func TestFeedback2(t *testing.T) {
	program := "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
	expectedThrust := 18216
	expectedAmpSettings := []int{9, 7, 8, 5, 6}

	testAmpSettings(program, expectedThrust, t, expectedAmpSettings, 5, 9)
}

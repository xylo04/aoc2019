package common

import (
	"testing"
)

func TestCompute(t *testing.T) {
	testData := map[string]string{
		"1,0,0,0,99":          "2,0,0,0,99",
		"2,3,0,3,99":          "2,3,0,6,99",
		"2,4,4,5,99,0":        "2,4,4,5,99,9801",
		"1,1,1,4,99,5,6,0,99": "30,1,1,4,2,5,6,0,99",
	}
	for inputStr, expectedStr := range testData {
		mem, _ := AizuArray(inputStr, ",")
		expected, _ := AizuArray(expectedStr, ",")

		ExecuteIntcode(&mem, 0)
		for i := 0; i < len(mem); i++ {
			if mem[i] != expected[i] {
				t.Errorf("For input %s, final state was expected to be %d but was actually %d",
					inputStr, expected, mem)
				break
			}
		}
	}
}

func TestEcho(t *testing.T) {
	program := "3,0,4,0,99"
	for i := -100; i < 100; i++ {
		mem, _ := AizuArray(program, ",")
		actual := ExecuteIntcode(&mem, i)
		if actual != i {
			t.Errorf("Echo intcode test for %d was actually %d", i, actual)
			break
		}
	}
}

func TestParamMode(t *testing.T) {
	mem, _ := AizuArray("1002,4,3,4,33", ",")
	expected, _ := AizuArray("1002,4,3,4,99", ",")
	ExecuteIntcode(&mem, 0)
	for i := 0; i < len(mem); i++ {
		if mem[i] != expected[i] {
			t.Errorf("Parameter mode intcode test failed, should be %d but was %d", expected, mem)
			break
		}
	}
}

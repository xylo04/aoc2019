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
		state, _ := AizuArray(inputStr, ",")
		expected, _ := AizuArray(expectedStr, ",")

		ExecuteIntcode(&state)
		for i := 0; i < len(state); i++ {
			if state[i] != expected[i] {
				t.Errorf("For input [%s], final state was expected to be [%s] but was actually %d",
					inputStr, expectedStr, state)
				break
			}
		}
	}
}

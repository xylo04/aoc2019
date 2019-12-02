package main

import (
	"github.com/xylo04/aoc2019/common"
	"log"
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
		log.Println(inputStr, expectedStr)
		state, _ := common.AizuArray(inputStr, ",")
		expected, _ := common.AizuArray(expectedStr, ",")

		compute(&state)
		for i := 0; i < len(state); i++ {
			if state[i] != expected[i] {
				t.Errorf("For input [%s], final state was expected to be [%s] but was actually %d",
					inputStr, expectedStr, state)
				break
			}
		}
	}
}

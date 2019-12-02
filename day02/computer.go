package main

import (
	"fmt"
	"github.com/xylo04/aoc2019/common"
)

func main() {
	lines, err := common.FileToLines("input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		state, _ := common.AizuArray(line, ",")
		state[1] = 12
		state[2] = 2
		fmt.Printf("Initial state: %d\n", state)
		compute(&state)
		fmt.Printf("Final state: %d\n", state)
	}
}

func compute(state *[]int) {
	for i := 0; true ; i += 4 {
		opcode := (*state)[i]
		if opcode == 99 {
			return
		}
		src1 := (*state)[i+1]
		src2 := (*state)[i+2]
		dest := (*state)[i+3]
		val1 := (*state)[src1]
		val2 := (*state)[src2]
		if opcode == 1 {
			(*state)[dest] = val1 + val2
		}
		if opcode == 2 {
			(*state)[dest] = val1 * val2
		}
	}
}

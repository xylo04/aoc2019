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
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			state, _ := common.AizuArray(lines[0], ",")
			state[1] = n
			state[2] = v
			compute(&state)
			if state[0] == 19690720 {
				fmt.Printf("%d\n", 100*n+v)
				return
			}
		}
	}
	fmt.Println("Nothing resolved to the given value")
}

func compute(state *[]int) {
	for i := 0; true; i += 4 {
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

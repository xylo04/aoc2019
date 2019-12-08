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
			common.ExecuteIntcode(&state, []int{0})
			if state[0] == 19690720 {
				fmt.Printf("%d\n", 100*n+v)
				return
			}
		}
	}
	fmt.Println("Nothing resolved to the given value")
}

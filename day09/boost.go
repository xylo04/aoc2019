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
	mem, _ := common.AizuArray(lines[0], ",")
	in := make(chan int, 1)
	out := make(chan int, 1)

	// Test input
	in <- 1
	go common.NewIntcode(mem, in, out).Execute()

	for output := range out {
		fmt.Println(output)
	}
}

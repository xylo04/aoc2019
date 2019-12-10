package main

import (
	"fmt"
	"github.com/xylo04/aoc2019/common"
)

func main() {
	const radiator = 5
	lines, err := common.FileToLines("input.txt")
	if err != nil {
		panic(err)
	}
	mem, _ := common.AizuArray(lines[0], ",")
	in := make(chan int, 1)
	out := make(chan int, 1)
	in <- radiator
	common.NewIntcode(mem, in, out).Execute()
	testResult := <-out
	fmt.Printf("%d", testResult)
}

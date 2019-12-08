package main

import (
	"fmt"
	"github.com/xylo04/aoc2019/common"
)

func main() {
	const airConditioner = 1
	lines, err := common.FileToLines("input.txt")
	if err != nil {
		panic(err)
	}
	mem, _ := common.AizuArray(lines[0], ",")
	testResult := common.ExecuteIntcode(&mem, airConditioner)
	fmt.Printf("%d", testResult)
}

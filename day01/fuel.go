package main

import (
	"fmt"
	"github.com/xylo04/aoc2019/common"
	"log"
	"strconv"
)

func fuel(mass int) int {
	f := mass/3 - 2
	if f > 0 {
		return f + fuel(f)
	} else {
		return 0
	}
}

func main() {
	lines, err := common.FileToLines("input.txt")
	if err != nil {
		panic(err)
	}
	total := 0
	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		fuel := fuel(mass)
		total += fuel
		log.Printf("For mass %d, fuel is %d", mass, fuel)
	}
	fmt.Printf("%d\n", total)
}

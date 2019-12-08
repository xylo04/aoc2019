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
	max, maxSetting := findMaxThrust(lines[0])
	fmt.Printf("%d at %d", max, maxSetting)
}

func findMaxThrust(program string) (int, []int) {
	max := 0
	var maxSetting []int
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			if b == a {
				continue
			}
			for c := 0; c < 5; c++ {
				if c == a || c == b {
					continue
				}
				for d := 0; d < 5; d++ {
					if d == a || d == b || d == c {
						continue
					}
					for e := 0; e < 5; e++ {
						if e == a || e == b || e == c || e == d {
							continue
						}
						ampSetting := []int{a, b, c, d, e}
						thrust := thrusterAmplifiers(program, ampSetting)
						if thrust > max {
							max = thrust
							maxSetting = ampSetting
						}
					}
				}
			}
		}
	}
	return max, maxSetting
}

func thrusterAmplifiers(program string, ampSettings []int) int {
	thrust := 0
	for amp := 0; amp < 5; amp++ {
		mem, _ := common.AizuArray(program, ",")
		thrust = common.ExecuteIntcode(&mem, []int{ampSettings[amp], thrust})
	}
	return thrust
}

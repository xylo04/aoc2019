package main

import (
	"fmt"
	"github.com/xylo04/aoc2019/common"
	"sync"
)

func main() {
	lines, err := common.FileToLines("input.txt")
	if err != nil {
		panic(err)
	}
	max, maxSetting := findMaxThrust(lines[0], 5, 9)
	fmt.Printf("%d at %d", max, maxSetting)
}

func findMaxThrust(program string, rangeMin int, rangeMax int) (int, []int) {
	max := 0
	var maxSetting []int
	for a := rangeMin; a <= rangeMax; a++ {
		for b := rangeMin; b <= rangeMax; b++ {
			if b == a {
				continue
			}
			for c := rangeMin; c <= rangeMax; c++ {
				if c == a || c == b {
					continue
				}
				for d := rangeMin; d <= rangeMax; d++ {
					if d == a || d == b || d == c {
						continue
					}
					for e := rangeMin; e <= rangeMax; e++ {
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
	// initialize io channels with amp settings
	var io = make([]chan int, 5)
	for amp := 0; amp < 5; amp++ {
		io[amp] = make(chan int, 5)
		io[amp] <- ampSettings[amp]
	}

	// initialize input thrust to amp A
	io[0] <- 0

	// start amps
	wg := &sync.WaitGroup{}
	for amp := 0; amp < 5; amp++ {
		mem, _ := common.AizuArray(program, ",")
		wg.Add(1)
		go common.NewIntcodeWithWg(mem, io[amp], io[(amp+1)%5], wg).Execute()
	}
	wg.Wait()
	return <-io[0]
}

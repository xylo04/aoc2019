package main

import (
	"fmt"
	"github.com/xylo04/aoc2019/common"
	"math"
	"strconv"
)

func main() {
	wid, hei := 25, 6
	lines, err := common.FileToLines("input.txt")
	if err != nil {
		panic(err)
	}
	image := parseFile(lines[0], wid, hei)
	fmt.Println(checksum(image))
}

type spaceImage struct {
	pixels [][][]int
	// per layer, a map of color value to count of occurrences
	stats []map[int]int
}

func parseFile(input string, wid int, hei int) spaceImage {
	if len(input)%(wid*hei) != 0 {
		panic("File is not an even multiple of given width and height")
	}
	image := spaceImage{}
	layers := len(input) / (hei * wid)
	image.pixels = make([][][]int, layers)
	image.stats = make([]map[int]int, layers)
	for l := 0; l < layers; l++ {
		image.pixels[l] = make([][]int, hei)
		image.stats[l] = make(map[int]int)
		for h := 0; h < hei; h++ {
			image.pixels[l][h] = make([]int, wid)
			for w := 0; w < wid; w++ {
				pixelChar := input[(l*hei*wid)+(h*wid)+w]
				pixelVal, _ := strconv.Atoi(string(pixelChar))
				image.pixels[l][h][w] = pixelVal
				image.stats[l][pixelVal]++
			}
		}
	}
	return image
}

func checksum(image spaceImage) int {
	// find layer with fewest zeros
	minValue := math.MaxInt32
	minLayer := -1
	for l := 0; l < len(image.pixels); l++ {
		if image.stats[l][0] < minValue {
			minValue = image.stats[l][0]
			minLayer = l
		}
	}
	// checksum is number of 1's * number of 2's
	return image.stats[minLayer][1] * image.stats[minLayer][2]
}

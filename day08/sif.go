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
	decoded := decode(image)
	for h := 0; h < len(decoded); h++ {
		for w := 0; w < len(decoded[h]); w++ {
			if decoded[h][w] == 1 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

type spaceImage struct {
	wid, hei, layers int
	pixels           [][][]int
	// per layer, a map of color value to count of occurrences
	stats []map[int]int
}

func parseFile(input string, wid int, hei int) spaceImage {
	if len(input)%(wid*hei) != 0 {
		panic("File is not an even multiple of given width and height")
	}
	image := spaceImage{}
	image.wid, image.hei, image.layers = wid, hei, len(input)/(hei*wid)
	image.pixels = make([][][]int, image.layers)
	image.stats = make([]map[int]int, image.layers)
	for l := 0; l < image.layers; l++ {
		image.pixels[l] = make([][]int, image.hei)
		image.stats[l] = make(map[int]int)
		for h := 0; h < image.hei; h++ {
			image.pixels[l][h] = make([]int, image.wid)
			for w := 0; w < image.wid; w++ {
				pixelChar := input[(l*image.hei*image.wid)+(h*image.wid)+w]
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
	for l := 0; l < image.layers; l++ {
		if image.stats[l][0] < minValue {
			minValue = image.stats[l][0]
			minLayer = l
		}
	}
	// checksum is number of 1's * number of 2's
	return image.stats[minLayer][1] * image.stats[minLayer][2]
}

func decode(image spaceImage) [][]int {
	decoded := make([][]int, image.hei)
	for h := 0; h < image.hei; h++ {
		decoded[h] = make([]int, image.wid)
		for w := 0; w < image.wid; w++ {
			for l := 0; l < image.layers; l++ {
				if image.pixels[l][h][w] != 2 {
					decoded[h][w] = image.pixels[l][h][w]
					break
				}
			}
		}
	}
	return decoded
}

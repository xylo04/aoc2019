package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

// from https://siongui.github.io/2016/04/06/go-readlines-from-file-or-string/
func fileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func main() {
	lines, err := fileToLines("input.txt")
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

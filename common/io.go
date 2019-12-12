package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// from https://siongui.github.io/2016/04/06/go-readlines-from-file-or-string/
func FileToLines(filePath string) (lines []string, err error) {
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

// from https://stackoverflow.com/a/37767467
func AizuArray(arr string, sep string) ([]int, error) {
	const memSize = 100000
	a := strings.Split(arr, sep)
	b := make([]int, memSize)
	var err error
	for i, v := range a {
		b[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}

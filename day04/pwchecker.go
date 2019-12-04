package main

import (
	"fmt"
	"log"
	"strconv"
)

func checkPassword(password string) bool {
	if len(password) != 6 {
		return false
	}
	var repeats []string
	for i := 1; i < len(password); i++ {
		lastChar, thisChar := password[i-1], password[i]
		if lastChar > thisChar {
			return false
		}
		if lastChar == thisChar {
			repeats = append(repeats, string(thisChar))
		}
	}
	if len(repeats) == 0 {
		return false
	}
	if len(repeats) == 1 {
		return true
	}

	repeatCounts := map[string]int{}
	for i := 0; i < len(repeats); i++ {
		repeatCounts[repeats[i]]++
	}
	for _, v := range repeatCounts {
		if v == 1 {
			return true
		}
	}
	return false
}

func main() {
	candidates := 0
	start, end := 248345, 746315
	for i := start; i < end; i++ {
		if i%1000 == 0 {
			log.Printf("Working on %d", i)
		}
		if checkPassword(strconv.Itoa(i)) {
			candidates++
		}
	}
	fmt.Println(candidates)
}

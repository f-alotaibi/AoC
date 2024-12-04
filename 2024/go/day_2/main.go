package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	safeLines := 0
	isIncreasing := false
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		safe := true
		for i := 0; i < len(numbers)-1; i++ {
			x, err := strconv.Atoi(numbers[i])
			if err != nil {
				safe = false
				break
			}
			if i+1 == len(numbers) {
				safe = false
				break
			}
			y, err := strconv.Atoi(numbers[i+1])
			if err != nil {
				safe = false
				break
			}
			diff := (x - y)
			if i == 0 {
				isIncreasing = diff < 0
			}
			if diff == 0 ||
				(isIncreasing && diff < -3) ||
				(!isIncreasing && diff > 3) ||
				(isIncreasing && diff > 0) ||
				(!isIncreasing && diff < 0) {
				safe = false
				break
			}
		}
		if safe && len(numbers) > 1 {
			safeLines++
		}
	}
	println(safeLines)
}

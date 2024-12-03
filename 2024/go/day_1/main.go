package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//go:embed input.txt
var input string

func main() {
	newLineSplit := strings.Split(input, "\n")
	sortedX := make([]int, 0)
	sortedY := make([]int, 0)
	for _, line := range newLineSplit {
		numbers := strings.Split(line, "   ")
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Printf("x: %s\n", numbers[0])
			continue
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Printf("y: %s\n", numbers[1])
			continue
		}
		sortedX = append(sortedX, x)
		sortedY = append(sortedY, y)
		sort.Ints(sortedX)
		sort.Ints(sortedY)
	}
	sum := 0
	for i := 0; i < len(sortedX); i++ {
		sum += Abs(sortedX[i] - sortedY[i])
	}
	println(sum)
}

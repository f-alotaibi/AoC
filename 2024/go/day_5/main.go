package main

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func convertStringArrayToIntArray(arr []string) ([]int, error) {
	result := make([]int, 0, len(arr))
	for _, elem := range arr {
		intElem, err := strconv.Atoi(elem)
		if err != nil {
			return result, err
		}
		result = append(result, intElem)
	}
	return result, nil
}

var rules map[int]int

// Part 2
func comparator(a int, b int) int {
	return rules[a+(b<<8)]
}

func main() {
	rules = make(map[int]int)
	splitInput := strings.Split(input, "\n")
	var i int
	sum := 0
	for i = 0; splitInput[i] != ""; i++ {
		splitPages := strings.Split(splitInput[i], "|")
		pages, err := convertStringArrayToIntArray(splitPages)
		if err != nil {
			panic(err)
		}
		rules[pages[0]+(pages[1]<<8)] = -1
		rules[pages[1]+(pages[0]<<8)] = 1
	}
	for _, updateLine := range splitInput[i+1:] {
		updatesSplit := strings.Split(updateLine, ",")
		updates, err := convertStringArrayToIntArray(updatesSplit)
		if err != nil {
			panic(err)
		}
		ordered := true
		for j, currentUpdate := range updates {
			for k := j + 1; k < len(updates); k++ {
				if rules[updates[k]+(currentUpdate<<8)] == -1 {
					ordered = false
					//break // part 1
				}
			}
		}
		if ordered {
			// Part 1
			//sum += updates[len(updates)/2]
		} else {
			// Part 2
			slices.SortFunc[[]int](updates, comparator)
			sum += updates[len(updates)/2]
		}
	}
	println(sum)
}

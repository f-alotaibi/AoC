package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	re := regexp.MustCompile("(XMAS|SAMX)")
	count := 0
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		chars := strings.Split(line, "")
		for j, char := range chars {
			if char == "X" || char == "S" {
				// Horizontal
				if j+3 < len(chars) {
					if re.MatchString(strings.Join(chars[j:j+4], "")) {
						fmt.Printf("%d horz: %s\n", i, strings.Join(chars[j:j+4], ""))
						count++
					}
				}
				if i+4 > len(lines) {
					continue
				}
				// Vertical
				vTest := char
				// Downward Diagional
				dLeft := char
				dRight := char
				for k := 1; k < 4; k++ {
					kSplit := strings.Split(lines[i+k], "")
					if j-k >= 0 {
						dLeft += kSplit[j-k]
					}
					if j+k < len(lines) {
						dRight += kSplit[j+k]
					}
					vTest += kSplit[j]
				}
				if re.MatchString(vTest) {
					fmt.Printf("%d vTest: %s\n", i, vTest)
					count++
				}
				if re.MatchString(dLeft) {
					fmt.Printf("%d dLeft: %s\n", i, dLeft)
					count++
				}
				if re.MatchString(dRight) {
					fmt.Printf("%d dRight: %s\n", i, dRight)
					count++
				}
			}
		}
	}

	println(count)
}

package main

import (
	_ "embed"
	"regexp"
)

//go:embed input.txt
var input []byte

func main() {
	re := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
	matches := re.FindAll(input, -1)
	sum := 0
	for _, match := range matches {
		match = match[4:]
		match = match[:len(match)-1]
		x, y := 0, 0
		doneFirst := false
		for _, b := range match {
			if b == 44 {
				doneFirst = true
				continue
			}
			number := int(b - 48)
			if doneFirst {
				y = (y * 10) + number
			} else {
				x = (x * 10) + number
			}
		}
		if !doneFirst {
			println("Something wrong happened")
			break
		}
		sum += (x * y)
	}
	println(sum)
}

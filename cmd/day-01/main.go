package main

import (
	"aoc-2021/internal"
	"fmt"
)

var inputTest = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func part01(input []int) int {
	incCount := 0
	for i := 1; i < len(input); i++ {
		if (input[i] > input[i - 1]) {
			incCount += 1
		}
	}
	return incCount
}

func part02(input []int) int {
	incCount := 0
	for i := 3; i < len(input); i++ {
		prevWindow := input[i-1] + input[i-2] + input[i-3]
		currWindow := input[i] + input[i-1] + input[i-2]
		if (currWindow > prevWindow) {
			incCount += 1
		}
	}
	return incCount
}

func main() {
	input := internal.ReadFileSliceInt("cmd/day-01/input.txt")
	fmt.Println(part01(input))
	fmt.Println(part02(input))
}

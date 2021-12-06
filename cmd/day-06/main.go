package main

import (
	"aoc-2021/internal"
	"fmt"
	"strings"
)

var inputRaw = "3,4,3,1,2"

type Line struct {
	x1, y1, x2, y2 int
}

func main() {
	input := internal.ReadFileLines("cmd/day-06/input.txt")[0]
	//input := inputRaw

	lanternfishes := internal.ConvertStringsToInts(strings.Split(input, ","))
	lanternFishesCount := make([]int64, 9)
	for i := 0; i < len(lanternfishes); i++ {
		lanternFishesCount[lanternfishes[i]] += 1
	}

	fmt.Println(lanternFishesCount)

	for i := 0; i < 256; i++ {
		newFishesCount := make([]int64, 9)

		// Spawn
		if lanternFishesCount[0] > 0 {
			newFishesCount[8] += lanternFishesCount[0]
			newFishesCount[6] += lanternFishesCount[0]
			newFishesCount[0] = 0
		}

		for j := 1; j < len(lanternFishesCount); j++ {
			newFishesCount[j-1] += lanternFishesCount[j]
		}

		lanternFishesCount = newFishesCount
	}

	fmt.Println(lanternFishesCount)
	count := int64(0)
	for i := 0; i < len(lanternFishesCount); i++ {
		count += lanternFishesCount[i]
	}

	fmt.Println(count)
}

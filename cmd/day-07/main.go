package main

import (
	"aoc-2021/internal"
	"fmt"
	"math"
	"strings"
)

var inputRaw = "16,1,2,0,4,2,7,1,2,14"

func main() {
	input := internal.ReadFileLines("cmd/day-07/input.txt")[0]
	//input := inputRaw

	crabsPos := internal.ConvertStringsToInts(strings.Split(input, ","))
	furthestCrab := 0
	for _, po := range crabsPos {
		if po > furthestCrab {
			furthestCrab = po
		}
	}

	leastUsedFuel := 9999999999

	for i := 0; i < furthestCrab; i++ {
		fuelForBaseCrab := 0
		for j := 0; j < len(crabsPos); j++ {
			// Part 1
			fuelNeeded := int(math.Abs(float64(crabsPos[j] - i)))
			// Part 2
			fuelNeeded = (fuelNeeded * (1 + fuelNeeded)) / 2

			fuelForBaseCrab += fuelNeeded
		}

		if fuelForBaseCrab < leastUsedFuel {
			leastUsedFuel = fuelForBaseCrab
		}
	}

	fmt.Println(leastUsedFuel)
}

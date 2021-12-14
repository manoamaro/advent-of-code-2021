package main

import (
	"aoc-2021/internal"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := internal.ReadFileLines("cmd/day-14/input.txt")
	//input := strings.Split("NNCB\n\nCH -> B\nHH -> N\nCB -> H\nNH -> C\nHB -> C\nHC -> B\nHN -> C\nNN -> C\nBH -> H\nNC -> B\nNB -> B\nBN -> B\nBB -> N\nBC -> B\nCC -> N\nCN -> C", "\n")

	pairInsertion := map[string]string{}

	polymerTemplate := map[string]uint64{}

	initialPolymerTemplate := ""

	for i, line := range input {
		if i == 0 {
			initialPolymerTemplate = line
			continue
		}

		if i > 1 {
			values := strings.Split(line, " -> ")
			pairInsertion[values[0]] = values[1]
		}
	}

	for i := 0; i < len(initialPolymerTemplate)-1; i++ {
		polymerTemplate[initialPolymerTemplate[i:i+2]] += 1
	}

	fmt.Println(polymerTemplate)

	for step := 0; step < 40; step++ {
		newTemplate := map[string]uint64{}
		for pair, count := range polymerTemplate {
			insertion := pairInsertion[pair]
			newPair1 := pair[0:1] + insertion
			newPair2 := insertion + pair[1:2]
			newTemplate[newPair1] += count
			newTemplate[newPair2] += count
		}

		polymerTemplate = newTemplate
		fmt.Println(polymerTemplate)
	}

	elementsCount := map[string]uint64{}

	maxCount := uint64(0)

	for pair, count := range polymerTemplate {
		elementsCount[pair[0:1]] += count
		elementsCount[pair[1:2]] += count
		if elementsCount[pair[0:1]] > maxCount {
			maxCount = elementsCount[pair[0:1]]
		}
		if elementsCount[pair[1:2]] > maxCount {
			maxCount = elementsCount[pair[1:2]]
		}
	}

	minCount := maxCount
	for _, count := range elementsCount {
		if count < minCount {
			minCount = count
		}
	}

	fmt.Println(elementsCount)
	fmt.Println(maxCount)
	fmt.Println(minCount)
	fmt.Println(maxCount - minCount)
	fmt.Println(uint64(math.Ceil(float64(maxCount-minCount) / 2)))
}

func old(input []string) {
	polymerTemplate := ""
	pairInsertion := map[string]string{}

	for i, line := range input {
		if i == 0 {
			polymerTemplate = line
			continue
		}

		if i > 1 {
			values := strings.Split(line, " -> ")
			pairInsertion[values[0]] = values[1]
		}
	}
	for step := 0; step < 10; step++ {
		newTemplate := ""
		for i := 0; i < len(polymerTemplate)-1; i++ {
			pair := polymerTemplate[i : i+2]
			insertion := pairInsertion[pair]
			newTemplate += pair[:1] + insertion
		}
		newTemplate += polymerTemplate[len(polymerTemplate)-1:]
		polymerTemplate = newTemplate
	}

	elements := map[string]uint64{}

	mostCommonElement, mostCount := "", uint64(0)

	for _, element := range polymerTemplate {
		elements[string(element)] += 1
		if mostCount < elements[string(element)] {
			mostCount = elements[string(element)]
			mostCommonElement = string(element)
		}
	}

	leastCommonElement, leastCount := "", mostCount

	for element, count := range elements {
		if count < leastCount {
			leastCommonElement = element
			leastCount = count
		}
	}

	fmt.Println(mostCommonElement)
	fmt.Println(leastCommonElement)
	fmt.Println(mostCount - leastCount)
	fmt.Println(polymerTemplate)
}

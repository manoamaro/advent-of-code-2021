package main

import (
	"aoc-2021/internal"
	"fmt"
	"sort"
	"strings"
)

var possibleDisplayConfigs = internal.Permutations([]string{"a", "b", "c", "d", "e", "f", "g"})

var displayDigitsPos = [][]int{
	{0, 1, 2, 4, 5, 6},
	{2, 5},
	{0, 2, 3, 4, 6},
	{0, 2, 3, 5, 6},
	{1, 2, 3, 5},
	{0, 1, 3, 5, 6},
	{0, 1, 3, 4, 5, 6},
	{0, 2, 5},
	{0, 1, 2, 3, 4, 5, 6},
	{0, 1, 2, 3, 5, 6},
}

type Entry struct {
	signals []string
	output  []string
}

func parseEntry(raw string) Entry {
	parts := strings.Split(raw, "|")
	signals := strings.Split(strings.Trim(parts[0], " "), " ")
	output := strings.Split(strings.Trim(parts[1], " "), " ")
	return Entry{
		signals: signals,
		output:  output,
	}
}

func (e Entry) convert(displayConfig []string) ([]int, []int) {
	config := strings.Join(displayConfig, "")
	signalsConv := make([]int, 0)
	outputConv := make([]int, 0)
	for _, signal := range e.signals {
		signalConns := make([]int, 0)
		for _, signalWire := range signal {
			signalConns = append(signalConns, strings.Index(config, string(signalWire)))
		}
		sort.Ints(signalConns)

		for i := 0; i < len(displayDigitsPos); i++ {
			if internal.IntsEqual(displayDigitsPos[i], signalConns) {
				signalsConv = append(signalsConv, i)
				break
			}
		}
	}

	for _, output := range e.output {
		outputDigits := make([]int, 0)
		for _, outputWire := range output {
			outputDigits = append(outputDigits, strings.Index(config, string(outputWire)))
		}
		sort.Ints(outputDigits)

		for i := 0; i < len(displayDigitsPos); i++ {
			if internal.IntsEqual(displayDigitsPos[i], outputDigits) {
				outputConv = append(outputConv, i)
				break
			}
		}
	}

	return signalsConv, outputConv
}

func main() {
	input := internal.ReadFileLines("cmd/day-08/input.txt")
	//input := inputRaw

	entries := make([]Entry, 0)
	for _, line := range input {
		entries = append(entries, parseEntry(line))
	}

	count := 0

	for _, entry := range entries {
		for _, out := range entry.output {
			if len(out) == 2 || len(out) == 4 || len(out) == 3 || len(out) == 7 {
				count += 1
			}
		}
	}

	fmt.Println(count)

	countPt2 := 0

	//entries = []Entry{parseEntry("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")}

	for _, entry := range entries {
		for _, config := range possibleDisplayConfigs {
			signals, output := entry.convert(config)
			if len(signals) == len(entry.signals) && len(output) == len(entry.output) {
				countPt2 += internal.SliceToInt(output)
			}
		}
	}
	fmt.Println(countPt2)
}

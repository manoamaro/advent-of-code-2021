package main

import (
	"aoc-2021/internal"
	"fmt"
	"strconv"
	"strings"
)

var inputTest = []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

func part01(input []string) int64 {
	matrix := make([][]string, len(input))
	for i, rawBits := range input {
		matrix[i] = make([]string, len(rawBits))
		for i2, rawBit := range rawBits {
			matrix[i][i2] = string(rawBit)
		}
	}
	bitsPerRow := len(matrix[0])

	gamma := ""
	epsilon := ""

	for i := 0; i < bitsPerRow; i++ {
		bit0Count := 0
		bit1Count := 0
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] == "0" {
				bit0Count += 1
			} else {
				bit1Count += 1
			}
		}
		if bit0Count > bit1Count {
			// 0 most common
			gamma = gamma + "0"
			// 1 less common
			epsilon = epsilon + "1"
		} else {
			// 1 most common
			gamma = gamma + "1"
			// 0 less common
			epsilon = epsilon + "0"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	return gammaInt * epsilonInt
}

type Col struct {
	count0, count1 int
}

func getBitsPerCol(matrix [][]int) []Col {
	colCount := len(matrix[0])
	cols := make([]Col, colCount)
	for i := 0; i < colCount; i++ {
		cols[i] = Col{
			count0: 0,
			count1: 0,
		}
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] == 0 {
				cols[i].count0 += 1
			} else {
				cols[i].count1 += 1
			}
		}
	}
	return cols
}

func filterOxygen(matrix [][]int, colSize int, colPos int) [][]int {
	if colPos < colSize {
		filteredMatrix := make([][]int, 0)
		cols := getBitsPerCol(matrix)
		for _, line := range matrix {
			if (cols[colPos].count0 > cols[colPos].count1 && line[colPos] == 0) || (cols[colPos].count1 >= cols[colPos].count0 && line[colPos] == 1) {
				filteredMatrix = append(filteredMatrix, line)
			}
		}
		if len(filteredMatrix) > 0 {
			return filterOxygen(filteredMatrix, colSize, colPos+1)
		} else {
			return matrix
		}
	} else {
		return matrix
	}
}

func filterCO2(matrix [][]int, colSize int, colPos int) [][]int {
	if colPos < colSize {
		filteredMatrix := make([][]int, 0)
		cols := getBitsPerCol(matrix)
		for _, line := range matrix {
			if (cols[colPos].count0 <= cols[colPos].count1 && line[colPos] == 0) || (cols[colPos].count1 < cols[colPos].count0 && line[colPos] == 1) {
				filteredMatrix = append(filteredMatrix, line)
			}
		}
		if len(filteredMatrix) > 0 {
			return filterCO2(filteredMatrix, colSize, colPos+1)
		} else {
			return matrix
		}
	} else {
		return matrix
	}
}

func part02(input []string) int64 {
	matrix := make([][]int, len(input))
	for i, rawBits := range input {
		matrix[i] = make([]int, len(rawBits))
		for i2, rawBit := range rawBits {
			matrix[i][i2], _ = strconv.Atoi(string(rawBit))
		}
	}

	oxigenInts := filterOxygen(matrix, len(matrix[0]), 0)
	oxigenStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(oxigenInts)), ""), "[]")
	oxygenInt, _ := strconv.ParseInt(oxigenStr, 2, 64)

	co2Ints := filterCO2(matrix, len(matrix[0]), 0)
	co2Str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(co2Ints)), ""), "[]")
	co2Int, _ := strconv.ParseInt(co2Str, 2, 64)

	return oxygenInt * co2Int
}

func main() {
	input := internal.ReadFileLines("cmd/day-03/input.txt")
	fmt.Println(part01(input))
	fmt.Println(part02(input))
}

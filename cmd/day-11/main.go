package main

import (
	"aoc-2021/internal"
	"fmt"
	"strconv"
)

func main() {
	input := internal.ReadFileLines("cmd/day-11/input.txt")
	//input := strings.Split("5483143223\n2745854711\n5264556173\n6141336146\n6357385478\n4167524645\n2176841721\n6882881134\n4846848554\n5283751526", "\n")

	octopuses := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		octopuses[i] = make([]int, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			octopuses[i][j], _ = strconv.Atoi(string(input[i][j]))
		}
	}

	flashCount := 0
	step := 0

	for ; ; step++ {
		printMatrix(octopuses)
		isSynced := true
		for i := 0; i < len(octopuses)-1; i++ {
			if !internal.IntsEqual(octopuses[i], octopuses[i+1]) {
				isSynced = false
			}
		}

		if isSynced {
			break
		}

		// Increment energy
		for i := 0; i < len(octopuses); i++ {
			for j := 0; j < len(octopuses[i]); j++ {
				octopuses[i][j] += 1
			}
		}

		// Flash
		for {
			willFlash := false
			for i := 0; i < len(octopuses); i++ {
				for j := 0; j < len(octopuses[i]); j++ {
					if octopuses[i][j] > 9 {
						willFlash = true
						break
					}
				}
			}

			if !willFlash {
				break
			}

			for i := 0; i < len(octopuses); i++ {
				for j := 0; j < len(octopuses[i]); j++ {
					// Flash
					if octopuses[i][j] > 9 {
						flashCount += 1
						octopuses[i][j] = 0
						// Left
						inc(octopuses, i-1, j)
						// Top-left
						inc(octopuses, i-1, j-1)
						// Top
						inc(octopuses, i, j-1)
						// Top-right
						inc(octopuses, i+1, j-1)
						// Right
						inc(octopuses, i+1, j)
						// Bottom-right
						inc(octopuses, i+1, j+1)
						// Bottom
						inc(octopuses, i, j+1)
						// Bottom-left
						inc(octopuses, i-1, j+1)
					}
				}
			}
		}
	}

	fmt.Println(flashCount)

	fmt.Println(step)
}

func inc(octopuses [][]int, x int, y int) {
	if x >= 0 && x < len(octopuses) && y >= 0 && y < len(octopuses[x]) && octopuses[x][y] > 0 {
		octopuses[x][y] += 1
	}
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

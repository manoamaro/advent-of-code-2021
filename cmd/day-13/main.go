package main

import (
	"aoc-2021/internal"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := internal.ReadFileLines("cmd/day-13/input.txt")
	//input := strings.Split("6,10\n0,14\n9,10\n0,3\n10,4\n4,11\n6,0\n6,12\n4,1\n0,13\n10,12\n3,4\n3,0\n8,4\n1,10\n2,14\n8,10\n9,0\n\nfold along y=7\nfold along x=5", "\n")

	paper := make([][]int, 0)
	readingCoords := true

	for _, line := range input {
		if line == "" {
			readingCoords = false
			fmt.Println("test")
		}

		if readingCoords {

			values := strings.Split(line, ",")
			x, _ := strconv.Atoi(values[0])
			y, _ := strconv.Atoi(values[1])
			paper = append(paper, []int{x, y})

		} else if line != "" {
			fmt.Println(paper)
			var foldLine string
			_, _ = fmt.Sscanf(line, "fold along %s", &foldLine)
			values := strings.Split(foldLine, "=")
			foldLine = values[0]
			coord, _ := strconv.Atoi(values[1])

			if foldLine == "y" {
				for i := 0; i <= coord; i++ {
					// get bottom line - coord * 2
					bottomLine := (coord * 2) - i
					for j := 0; j < len(paper); j++ {
						if paper[j][1] == bottomLine {
							fmt.Printf("%d,%d => %d,%d\n", paper[j][0], bottomLine, paper[j][0], i)
							paper[j][1] = i
						}
					}
				}
			} else {
				for i := 0; i <= coord; i++ {
					rightCol := (coord * 2) - i
					for j := 0; j < len(paper); j++ {
						if paper[j][0] == rightCol {
							fmt.Printf("%d,%d => %d,%d\n", rightCol, paper[j][1], i, paper[j][1])
							paper[j][0] = i
						}
					}
				}
			}
		}
	}

	distinctPaper := make([][]int, 0)

	for i := 0; i < len(paper); i++ {
		c := paper[i]
		exists := false
		for j := 0; j < len(distinctPaper); j++ {
			if distinctPaper[j][0] == c[0] && distinctPaper[j][1] == c[1] {
				exists = true
				break
			}
		}
		if !exists {
			distinctPaper = append(distinctPaper, c)
		}
	}

	fmt.Println(distinctPaper)
	fmt.Println(len(distinctPaper))

	maxX, maxY := 0, 0

	for i := 0; i < len(distinctPaper); i++ {
		if distinctPaper[i][0] > maxX {
			maxX = distinctPaper[i][0]
		}
		if distinctPaper[i][0] > maxY {
			maxY = distinctPaper[i][1]
		}
	}

	paperMatrix := make([][]bool, maxY+2)
	for i := 0; i < len(paperMatrix); i++ {
		paperMatrix[i] = make([]bool, maxX+2)
	}

	for _, p := range distinctPaper {
		paperMatrix[p[1]][p[0]] = true
	}

	for i := 0; i < len(paperMatrix); i++ {
		for j := 0; j < len(paperMatrix[i]); j++ {
			if paperMatrix[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

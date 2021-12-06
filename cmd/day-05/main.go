package main

import (
	"aoc-2021/internal"
	"fmt"
	"math"
)

var inputRaw = "0,9 -> 5,9\n8,0 -> 0,8\n9,4 -> 3,4\n2,2 -> 2,1\n7,0 -> 7,4\n6,4 -> 2,0\n0,9 -> 2,9\n3,4 -> 1,4\n0,0 -> 8,8\n5,5 -> 8,2"

type Line struct {
	x1, y1, x2, y2 int
}

func main() {
	input := internal.ReadFileLines("cmd/day-05/input.txt")
	//input := strings.Split(inputRaw, "\n")

	lines := make([]Line, len(input))
	vhLines := make([]Line, 0)

	maxX := 0
	maxY := 0

	for i, line := range input {
		fmt.Sscanf(line, "%d,%d -> %d,%d",
			&lines[i].x1,
			&lines[i].y1,
			&lines[i].x2,
			&lines[i].y2)

		if lines[i].x1 == lines[i].x2 || lines[i].y1 == lines[i].y2 {
			vhLines = append(vhLines, lines[i])
		}

		if lines[i].x1 > maxX {
			maxX = lines[i].x1
		}
		if lines[i].x2 > maxX {
			maxX = lines[i].x2
		}
		if lines[i].y1 > maxY {
			maxY = lines[i].y1
		}
		if lines[i].y2 > maxY {
			maxY = lines[i].y2
		}
	}

	diagram := make([]int, (maxX+1)*(maxY+1))

	for _, line := range lines {

		// Vertical
		if line.x1 == line.x2 || line.y1 == line.y2 {
			for y := math.Min(float64(line.y1), float64(line.y2)); y <= math.Max(float64(line.y1), float64(line.y2)); y++ {
				for x := math.Min(float64(line.x1), float64(line.x2)); x <= math.Max(float64(line.x1), float64(line.x2)); x++ {
					diagram[int(x)+(int(y)*maxY)] += 1
				}
			}
		} else {
			dx := int(float64(line.x2-line.x1) / math.Abs(float64(line.x2-line.x1)))
			dy := int(float64(line.y2-line.y1) / math.Abs(float64(line.y2-line.y1)))

			y := line.y1
			x := line.x1

			for y != line.y2+dy {
				diagram[x+(y*maxY)] += 1
				y += dy
				x += dx
			}
		}
	}

	count := 0
	for _, pos := range diagram {
		if pos > 1 {
			count += 1
		}
	}

	fmt.Println(count)
}

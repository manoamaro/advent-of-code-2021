package main

import (
	"aoc-2021/internal"
	"fmt"
	"sort"
	"strconv"
)

type Point struct {
	x, y  int
	value int
}

func (p Point) name() {

}

func main() {
	input := internal.ReadFileLines("cmd/day-09/input.txt")
	//input := []string{"2199943210", "3987894921", "9856789892", "8767896789", "9899965678"}

	heightmap := make([][]int, len(input))

	for i, line := range input {
		heightmap[i] = make([]int, len(line))
		for j, char := range line {
			heightmap[i][j], _ = strconv.Atoi(string(char))
		}
	}

	lowestPoints := make([]Point, 0)

	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			corners := make([]int, 0)
			if i > 0 {
				corners = append(corners, heightmap[i-1][j])
			}

			if j > 0 {
				corners = append(corners, heightmap[i][j-1])
			}

			if i < len(heightmap)-1 {
				corners = append(corners, heightmap[i+1][j])
			}

			if j < len(heightmap[i])-1 {
				corners = append(corners, heightmap[i][j+1])
			}

			sort.Ints(corners)
			if heightmap[i][j] < corners[0] {
				lowestPoints = append(lowestPoints, Point{
					x:     i,
					y:     j,
					value: heightmap[i][j],
				})
			}
		}
	}

	// Part 1
	risk := 0
	for _, point := range lowestPoints {
		risk += point.value + 1
	}

	fmt.Println(risk)

	// Part 2

	basinPoints := lowestPoints[:]
	basinsSizes := make([]int, len(lowestPoints))
	for i, point := range lowestPoints {
		basin := findBasin(heightmap, point.x, point.y, []Point{})
		basinsSizes[i] = len(basin) + 1
		basinPoints = append(basinPoints, basin...)
	}

	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			isLowerPoint := false
			for _, point := range lowestPoints {
				if point.x == i && point.y == j {
					isLowerPoint = true
					break
				}
			}
			isInBasin := false
			for _, point := range basinPoints {
				if point.x == i && point.y == j {
					isInBasin = true
					break
				}
			}
			if heightmap[i][j] == 9 {
				fmt.Printf("\u001b[34m#")
			} else if isLowerPoint {
				fmt.Printf("\u001b[31m%d", heightmap[i][j])
			} else if isInBasin {
				fmt.Printf("\u001b[32m%d", heightmap[i][j])
			} else {
				fmt.Printf("\u001b[0m%d", heightmap[i][j])
			}
		}
		fmt.Println("")
	}
}

func union(a, b []Point) []Point {
	m := make(map[Point]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}

func findBasin(heightmap [][]int, x int, y int, points []Point) []Point {
	currentValue := heightmap[x][y]

	// move top
	nextPoints := make([]Point, 0)
	if x-1 >= 0 && heightmap[x-1][y] < 9 && heightmap[x-1][y] > currentValue {
		nextPoints = append(nextPoints, Point{
			x:     x - 1,
			y:     y,
			value: heightmap[x-1][y],
		})
	}

	// move down
	if x+1 < len(heightmap) && heightmap[x+1][y] < 9 && heightmap[x+1][y] > currentValue {
		nextPoints = append(nextPoints, Point{
			x:     x + 1,
			y:     y,
			value: heightmap[x+1][y],
		})
	}

	// move left
	if y-1 >= 0 && heightmap[x][y-1] < 9 && heightmap[x][y-1] > currentValue {
		nextPoints = append(nextPoints, Point{
			x:     x,
			y:     y - 1,
			value: heightmap[x][y-1],
		})
	}

	// move right
	if y+1 < len(heightmap[x]) && heightmap[x][y+1] < 9 && heightmap[x][y+1] > currentValue {
		nextPoints = append(nextPoints, Point{
			x:     x,
			y:     y + 1,
			value: heightmap[x][y+1],
		})
	}

	visitedPoints := union(points, nextPoints)

	for _, point := range nextPoints {
		visitedPoints = union(visitedPoints, findBasin(heightmap, point.x, point.y, visitedPoints))
	}

	return visitedPoints
}

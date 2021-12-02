package main

import (
	"aoc-2021/internal"
	"fmt"
	"strconv"
	"strings"
)

var inputTest = []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

func part01(input []string) int {
	horizontal := 0
	depth := 0
	for _, rawCommand := range input {
		commandParts := strings.Split(rawCommand, " ")
		commandValue, _ := strconv.Atoi(commandParts[1])
		if commandParts[0] == "forward" {
			horizontal += commandValue
		} else if commandParts[0] == "down" {
			depth += commandValue
		} else if commandParts[0] == "up" {
			depth -= commandValue
		}
	}
	return horizontal * depth
}

func part02(input []string) int {
	horizontal := 0
	depth := 0
	aim := 0
	for _, rawCommand := range input {
		commandParts := strings.Split(rawCommand, " ")
		commandValue, _ := strconv.Atoi(commandParts[1])
		if commandParts[0] == "forward" {
			horizontal += commandValue
			depth += aim * commandValue
		} else if commandParts[0] == "down" {
			aim += commandValue
		} else if commandParts[0] == "up" {
			aim -= commandValue
		}
	}
	return horizontal * depth
}

func main() {
	input := internal.ReadFileLines("cmd/day-02/input.txt")
	fmt.Println(part01(input))
	fmt.Println(part02(input))
}

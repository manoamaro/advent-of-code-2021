package main

import (
	"aoc-2021/internal"
	"fmt"
	"strings"
)

var inputRaw = "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1\n\n22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19\n\n 3 15  0  2 22\n 9 18 13 17  5\n19  8  7 25 23\n20 11 10 24  4\n14 21 16 12  6\n\n14 21 17 24  4\n10 16 15  9 19\n18  8 23 26 20\n22 11 13  6  5\n 2  0 12  3  7"

type Board struct {
	board            []int
	markedNumbers    []int
	sumMarkedNumbers int
	win              bool
}

func (b Board) sumAllNumbers() (sum int) {
	for _, i := range b.board {
		sum += i
	}
	return
}

func (b *Board) mark(n int) int {
	if !b.win && internal.Contains(b.board, n) && !internal.Contains(b.markedNumbers, n) {
		b.markedNumbers = append(b.markedNumbers, n)
		b.sumMarkedNumbers = b.sumMarkedNumbers + n
		if b.check() {
			b.win = true
			return b.sumAllNumbers() - b.sumMarkedNumbers
		}
	}
	return 0
}

func (b Board) check() bool {
	for i := 0; i < 5; i++ {
		matchAllLine := true
		matchAllColumn := true

		for j := 0; j < 5; j++ {
			if !internal.Contains(b.markedNumbers, b.board[j+(5*i)]) {
				matchAllLine = false
				break
			}
		}

		for j := 0; j < 5; j++ {
			if !internal.Contains(b.markedNumbers, b.board[(j*5)+i]) {
				matchAllColumn = false
				break
			}
		}

		if matchAllLine || matchAllColumn {
			return true
		}
	}

	return false
}

func part01(randomNumbers []int, boards []Board) int {
	for i := 0; i < len(randomNumbers); i++ {
		for j := 0; j < len(boards); j++ {
			r := (&boards[j]).mark(randomNumbers[i])
			if r > 0 {
				return r * randomNumbers[i]
			}
		}
	}
	return 0
}

func part02(randomNumbers []int, boards []Board) int {
	wonBoards := make([]Board, 0)
	for i := 0; i < len(randomNumbers); i++ {
		for j := 0; j < len(boards); j++ {
			if (&boards[j]).mark(randomNumbers[i]) > 0 {
				wonBoards = append(wonBoards, boards[j])
			}
		}
	}

	lastBoard := wonBoards[len(wonBoards)-1]

	return (lastBoard.sumAllNumbers() - lastBoard.sumMarkedNumbers) * lastBoard.markedNumbers[len(lastBoard.markedNumbers)-1]
}

func main() {
	input := internal.ReadFileLines("cmd/day-04/input.txt")
	//input := strings.Split(inputRaw, "\n")

	numbers := internal.ConvertStringsToInts(strings.Split(input[0], ","))

	boards := make([]Board, 0)

	for _, line := range input[1:] {
		if len(line) <= 1 {
			boards = append(boards, Board{win: false})
		} else {
			board := &boards[len(boards)-1]
			boardLineNumbers := internal.ConvertStringsToInts(strings.Split(line, " "))
			board.board = append(board.board, boardLineNumbers...)
		}
	}

	fmt.Println(part01(numbers, boards))
	fmt.Println(part02(numbers, boards))
}

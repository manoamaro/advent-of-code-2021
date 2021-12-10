package main

import (
	"aoc-2021/internal"
	"fmt"
	"sort"
)

type Point struct {
	x, y  int
	value int
}

func (p Point) name() {

}

func main() {
	input := internal.ReadFileLines("cmd/day-10/input.txt")
	//input := strings.Split("[({(<(())[]>[[{[]{<()<>>\n[(()[<>])]({[<{<<[]>>(\n{([(<{}[<>[]}>{[]{[(<()>\n(((({<>}<{<{<>}{[]{[]{}\n[[<[([]))<([[{}[[()]]]\n[{[{({}]{}}([{[{{{}}([]\n{<[[]]>}<{[{[{[]{()[[[]\n[<(<(<(<{}))><([]([]()\n<{([([[(<>()){}]>(<<{{\n<{([{{}}[<[[[<>{}]]]>[]]", "\n")

	corruptedScore := 0
	incompleteScores := make([]int, 0)

	for _, line := range input {
		currentChunk := &internal.Node{}
		var firstIllegalChar rune

		for _, char := range line {
			if isOpen(char) {
				// Opening chunk
				currentChunk = currentChunk.AddChild(char)
			} else if isClose(char) && closingChar(currentChunk.Data.(rune)) == char {
				// Closing chunk
				currentChunk = currentChunk.Parent
			} else if isClose(char) && closingChar(currentChunk.Data.(rune)) != char {
				// corrupted line - closing with another char
				firstIllegalChar = char
				break
			}
		}

		if firstIllegalChar > 0 {
			// Corrupted line - calculate corruptedScore
			corruptedScore += illegalCharPoint(firstIllegalChar)
		} else if currentChunk.Parent != nil {
			// Incomplete line, complete it and corruptedScore
			incompleteScores = append(incompleteScores, recur(currentChunk, 0))
		}
	}

	sort.Ints(incompleteScores)

	fmt.Println(corruptedScore)
	fmt.Println(incompleteScores)
	fmt.Println(incompleteScores[len(incompleteScores)/2])
}

func recur(node *internal.Node, score int) int {
	value := missingCharPoint(closingChar(node.Data.(rune)))
	if node.Parent.Data == nil {
		return (score * 5) + value
	} else {
		return recur(node.Parent, (score*5)+value)
	}
}

func isOpen(char rune) bool {
	return char == '(' || char == '[' || char == '{' || char == '<'
}

func closingChar(openChar rune) rune {
	if openChar == '(' {
		return ')'
	} else if openChar == '[' {
		return ']'
	} else if openChar == '{' {
		return '}'
	} else if openChar == '<' {
		return '>'
	} else {
		return 0
	}
}

func illegalCharPoint(char rune) int {
	if char == ')' {
		return 3
	} else if char == ']' {
		return 57
	} else if char == '}' {
		return 1197
	} else if char == '>' {
		return 25137
	} else {
		return 0
	}
}

func missingCharPoint(char rune) int {
	if char == ')' {
		return 1
	} else if char == ']' {
		return 2
	} else if char == '}' {
		return 3
	} else if char == '>' {
		return 4
	} else {
		return 0
	}
}

func isClose(char rune) bool {
	return !isOpen(char)
}

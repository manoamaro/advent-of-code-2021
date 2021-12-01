package internal

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open")
	}
	return file
}

func ReadFileLines(path string) (r []string) {
	file := OpenFile(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
	file.Close()
	return
}

func ReadFileSliceInt(path string) (r []int) {
	lines := ReadFileLines(path)

	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		r = append(r, value)
	}

	return
}

func ConvertStringsToInts(i []string) (r []int) {
	for _, s := range i {
		v, err := strconv.Atoi(s)
		if err == nil {
			r = append(r, v)
		}
	}
	return
}
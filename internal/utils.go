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
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
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

func Contains(a []int, v int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == v {
			return true
		}
	}
	return false
}

// ContainsAll check if all items of a are contained in b
func ContainsAll(a, b []int) bool {
	for _, ai := range a {
		if !Contains(b, ai) {
			return false
		}
	}

	return true
}

func RotateMatrix(matrix [][]int) [][]int {

	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return matrix
}

func IntsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, aItem := range a {
		if aItem != b[i] {
			return false
		}
	}

	return true
}

func Permutations(arr []string) [][]string {
	var helper func([]string, int)
	var res [][]string

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func SliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

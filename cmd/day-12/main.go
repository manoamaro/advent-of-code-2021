package main

import (
	"aoc-2021/internal"
	"fmt"
	"strings"
)

func main() {
	input := internal.ReadFileLines("cmd/day-12/input.txt")
	//input := strings.Split("start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end", "\n")

	tunnels := internal.NewUndirectedGraph()

	for _, line := range input {
		path := strings.Split(line, "-")
		tunnels.AddVertex(path[0])
		tunnels.AddVertex(path[1])
		tunnels.AddEdge(path[0], path[1])
	}

	paths := make([][]string, 0)
	findPaths(tunnels, func(v []string) {
		//fmt.Println(v)
		paths = append(paths, v)
	})
	fmt.Println(len(paths))
}

func findPaths(graph *internal.Graph, pathFound func(s []string)) {
	_findPaths(graph, graph.Vertices["start"], map[string]int{}, []string{"start"}, pathFound)
}

func isLower(v string) bool {
	return strings.ToLower(v) == v
}

func smallCaveVisitedTwice(visited map[string]int) bool {
	for t, count := range visited {
		if isLower(t) && count > 1 {
			return true
		}
	}
	return false
}

func canVisitCave(visited map[string]int, cave string) bool {
	visitedCount := visited[cave]
	r := visitedCount == 0 || !isLower(cave) || ((cave == "start" || cave == "end") && visitedCount == 0) || (cave != "start" && cave != "end" && isLower(cave) && !smallCaveVisitedTwice(visited))

	return r
}

func _findPaths(graph *internal.Graph, start *internal.Vertex, visited map[string]int, path []string, pathFound func(s []string)) {

	if start.Key == "end" {
		pathFound(path)
		return
	}

	visited[start.Key]++

	for _, vertex := range start.Vertices {
		if canVisitCave(visited, vertex.Key) {
			path = append(path, vertex.Key)
			_findPaths(graph, vertex, visited, path, pathFound)
			path = path[:len(path)-1]
		}
	}

	visited[start.Key]--
}

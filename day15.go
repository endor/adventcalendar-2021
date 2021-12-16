package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var gridSize = 100

type Position struct {
	x, y int
}

func unvisitedNeighbours(pos Position, unvisited map[Position]int) []Position {
	u := []Position{}

	neighbours := []Position{
		{x: pos.x - 1, y: pos.y},
		{x: pos.x + 1, y: pos.y},
		{x: pos.x, y: pos.y - 1},
		{x: pos.x, y: pos.y + 1},
	}
	for _, n := range neighbours {
		if (n.x < 0) || (n.x >= gridSize*5) || (n.y < 0) || (n.y >= gridSize*5) {
			continue
		}
		if _, ok := unvisited[n]; ok {
			u = append(u, n)
		}
	}

	return u
}

func smallestUnvisited(unvisited map[Position]int) Position {
	smallestWeight := math.MaxInt
	smallestPosition := Position{x: 0, y: 0}

	for pos, weight := range unvisited {
		if weight < smallestWeight {
			smallestWeight = weight
			smallestPosition = pos
		}
	}

	return smallestPosition
}

func wrap(n int) int {
	if n <= 9 {
		return n
	} else {
		return n - 9
	}
}

func resize(field [][]int) [][]int {
	newField := make([][]int, len(field)*5)
	for n := range []int{1, 2, 3, 4, 5} {
		for m := range []int{1, 2, 3, 4, 5} {
			for y, row := range field {
				if m == 0 {
					newField[y+len(field)*n] = make([]int, len(field)*5)
				}
				for x, val := range row {
					newField[y+len(field)*n][x+len(field)*m] = wrap(val + m + n)
				}
			}
		}
	}
	return newField
}

func main() {
	bytes, err := os.ReadFile("day15.input")
	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}

	lines := strings.Split(string(bytes), "\n")
	field := make([][]int, len(lines))

	for idxLine, line := range lines {
		parts := strings.Split(line, "")
		weights := make([]int, len(parts))

		for idxPart, part := range parts {
			weight, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Cannot parse %s \n", part)
			}
			weights[idxPart] = weight
		}

		field[idxLine] = weights
	}

	field = resize(field)

	unvisited := make(map[Position]int)
	for y, row := range field {
		for x := range row {
			unvisited[Position{x, y}] = math.MaxInt
		}
	}
	unvisited[Position{x: 0, y: 0}] = 0

	source := Position{x: 0, y: 0}
	destination := Position{x: gridSize*5 - 1, y: gridSize*5 - 1}
	current := source

	for true {
		for _, n := range unvisitedNeighbours(current, unvisited) {
			if unvisited[current]+field[n.y][n.x] < unvisited[n] {
				unvisited[n] = unvisited[current] + field[n.y][n.x]
			}
		}
		delete(unvisited, current)

		if unvisited[destination] != math.MaxInt {
			fmt.Println(unvisited[destination])
			return
		}

		current = smallestUnvisited(unvisited)
	}
}

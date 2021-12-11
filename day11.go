package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var gridSize = 10

type Position struct {
	x, y int
}

type FlashMap = map[Position]struct{}

func readOctos(lines []string) [][]int {
	octos := make([][]int, len(lines))

	for idxLine, line := range lines {
		parts := strings.Split(line, "")
		numbers := make([]int, len(parts))

		for idxPart, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Cannot parse %s \n", part)
			}
			numbers[idxPart] = number
		}

		octos[idxLine] = numbers
	}

	return octos
}

func increaseEnergy(octos [][]int) ([]Position, [][]int) {
	needsFlash := []Position{}
	for y, row := range octos {
		for x := range row {
			octos[y][x] += 1
			if (octos[y][x]) > 9 {
				needsFlash = append(needsFlash, Position{x, y})
			}
		}
	}
	return needsFlash, octos
}

func surrounding(pos Position) []Position {
	s := []Position{}
	n := []int{-1, 0, 1}
	for _, y := range n {
		for _, x := range n {
			if (x == 0 && y == 0) || (pos.x+x < 0) || (pos.x+x >= gridSize) || (pos.y+y < 0) || (pos.y+y >= gridSize) {
				continue
			}
			s = append(s, Position{x: pos.x + x, y: pos.y + y})
		}
	}
	return s
}

func flash(octos [][]int, needsFlash []Position) (FlashMap, [][]int) {
	flashed := map[Position]struct{}{}

	for len(needsFlash) > 0 {
		pos := needsFlash[0]
		_, wasFlashed := flashed[pos]
		needsFlash = needsFlash[1:]

		if wasFlashed {
			continue
		}

		flashed[pos] = struct{}{}

		for _, s := range surrounding(pos) {
			octos[s.y][s.x] += 1
			if octos[s.y][s.x] > 9 {
				needsFlash = append(needsFlash, s)
			}
		}
	}

	return flashed, octos
}

func reset(octos [][]int, flashed FlashMap) [][]int {
	for f := range flashed {
		octos[f.y][f.x] = 0
	}
	return octos
}

func main() {
	bytes, err := os.ReadFile("day11.input")
	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}

	lines := strings.Split(string(bytes), "\n")
	octos := readOctos(lines)
	flashes := 0
	flashed := map[Position]struct{}{}
	needsFlash := []Position{}

	for i := 0; i < math.MaxInt; i++ {
		needsFlash, octos = increaseEnergy(octos)
		flashed, octos = flash(octos, needsFlash)
		if i < 100 {
			flashes += len(flashed)
		} else if len(flashed) == gridSize*gridSize {
			fmt.Println(i + 1)
			break
		}
		octos = reset(octos, flashed)
	}

	fmt.Println(flashes)
}

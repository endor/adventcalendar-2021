package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Minimum struct {
	x   int
	y   int
	val int
}

func valueOrMax(m [][]int, y int, x int) int {
	if len(m) == 0 {
		return math.MaxInt
	} else if y < 0 || y >= len(m) {
		return math.MaxInt
	} else if x < 0 || x >= len(m[0]) {
		return math.MaxInt
	} else {
		return m[y][x]
	}
}

func product(numbers []int) int {
	result := 1
	for _, n := range numbers {
		result *= n
	}
	return result
}

func riskLevel(minima []Minimum) int {
	result := 0
	for _, m := range minima {
		result += m.val + 1
	}
	return result
}

func findBasin(m [][]int, y int, x int, basin map[string]struct{}) map[string]struct{} {
	key := fmt.Sprintf("%d-%d", y, x)
	value := valueOrMax(m, y, x)

	if _, ok := basin[key]; ok || value >= 9 {
		return basin
	} else {
		basin[key] = struct{}{}
		basin = findBasin(m, y-1, x, basin)
		basin = findBasin(m, y+1, x, basin)
		basin = findBasin(m, y, x-1, basin)
		basin = findBasin(m, y, x+1, basin)
		return basin
	}
}

func main() {
	bytes, err := os.ReadFile("day09.input")
	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}
	lines := strings.Split(string(bytes), "\n")
	m := make([][]int, len(lines))

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

		m[idxLine] = numbers
	}

	minima := []Minimum{}
	for y, row := range m {
		for x, val := range row {
			north := valueOrMax(m, y-1, x)
			south := valueOrMax(m, y+1, x)
			east := valueOrMax(m, y, x-1)
			west := valueOrMax(m, y, x+1)

			if val < north && val < south && val < east && val < west {
				minima = append(minima, Minimum{x, y, val})
			}
		}
	}

	// Part 1
	fmt.Println(riskLevel(minima))

	basins := make([]map[string]struct{}, len(minima))
	for idx, min := range minima {
		basins[idx] = findBasin(m, min.y, min.x, map[string]struct{}{})
	}

	basinSizes := []int{}
	for _, basin := range basins {
		basinSizes = append(basinSizes, len(basin))
	}

	// Part 2
	sort.Ints(basinSizes)
	fmt.Println(product(basinSizes[len(basinSizes)-3:]))
}

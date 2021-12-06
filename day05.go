package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("day05.input")
	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}

	lines := strings.Split(string(bytes), "\n")

	field := make([][]int, 1000)
	for y := range field {
		row := make([]int, 1000)
		for x := range row {
			row[x] = 0
		}
		field[y] = row
	}

	for _, line := range lines {
		n := strings.Split(line, " -> ")

		n0 := strings.Split(n[0], ",")
		x1, _ := strconv.Atoi(n0[0])
		y1, _ := strconv.Atoi(n0[1])

		n1 := strings.Split(n[1], ",")
		x2, _ := strconv.Atoi(n1[0])
		y2, _ := strconv.Atoi(n1[1])

		dx := -1
		if x2 > x1 {
			dx = 1
		} else if x2 == x1 {
			dx = 0
		}

		dy := -1
		if y2 > y1 {
			dy = 1
		} else if y2 == y1 {
			dy = 0
		}

		y := y1
		x := x1
		for y != y2 || x != x2 {
			field[y][x]++
			x += dx
			y += dy
		}
		field[y][x]++
	}

	overlap := 0
	for _, y := range field {
		for x := range y {
			if y[x] > 1 {
				overlap++
			}
		}
	}

	fmt.Printf("%d \n", overlap)
}

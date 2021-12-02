package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(numbers []int) {
	current := math.MaxInt
	increases := 0
	for _, number := range numbers {
		if current < number {
			increases++
		}
		current = number
	}

	fmt.Println(increases)
}

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func part2(numbers []int) {
	window := []int{}
	increases := 0

	for _, number := range numbers {
		window = append(window, number)

		if len(window) < 4 {
			continue
		} else if len(window) > 4 {
			window = window[1:]
		}

		if sum(window[:3]) < sum(window[1:]) {
			increases++
		}
	}

	fmt.Println(increases)
}

func main() {
	bytes, err := os.ReadFile("day01.input")
	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}
	lines := strings.Split(string(bytes), "\n")
	numbers := []int{}

	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Cannot parse %s \n", line)
		}
		numbers = append(numbers, number)
	}

	part1(numbers)
	part2(numbers)
}

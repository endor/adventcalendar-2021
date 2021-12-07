package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func findEdges(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func calculateFuelConstant(pos int, positions []int) int {
	fuel := 0
	for _, p := range positions {
		fuel += int(math.Abs(float64(p) - float64(pos)))
	}
	return fuel
}

func calculateFuelLinear(pos int, positions []int) int {
	fuel := 0
	for _, p := range positions {
		distance := int(math.Abs(float64(p) - float64(pos)))
		pFuel := 0
		for i := 1; i <= distance; i++ {
			pFuel += i
		}
		fuel += pFuel
	}
	return fuel
}

func main() {
	bytes, err := os.ReadFile("day07.input")
	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}

	positions := []int{}

	for _, p0 := range strings.Split(string(bytes), ",") {
		p1, err := strconv.Atoi(p0)
		if err != nil {
			fmt.Printf("Cannot parse %s \n", p0)
		}
		positions = append(positions, p1)
	}

	min, max := findEdges((positions))

	minFuelConstant := math.MaxInt
	minFuelLinear := math.MaxInt
	for pos := min; pos <= max; pos++ {
		fuelConstant := calculateFuelConstant(pos, positions)
		if fuelConstant < minFuelConstant {
			minFuelConstant = fuelConstant
		}
		fuelLinear := calculateFuelLinear(pos, positions)
		if fuelLinear < minFuelLinear {
			minFuelLinear = fuelLinear
		}
	}

	fmt.Println(minFuelConstant)
	fmt.Println(minFuelLinear)
}

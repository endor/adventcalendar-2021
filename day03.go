package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string, line_len int) {
	gamma_rate := ""
	epsilon_rate := ""

	for i := 0; i < line_len; i++ {
		count_ones := 0

		for _, line := range lines {
			if line[i] == '1' {
				count_ones++
			}
		}

		if count_ones > len(lines)/2 {
			gamma_rate += "1"
			epsilon_rate += "0"
		} else {
			gamma_rate += "0"
			epsilon_rate += "1"
		}
	}

	gamma_rate_number, _ := strconv.ParseInt(gamma_rate, 2, 64)
	epsilon_rate_number, _ := strconv.ParseInt(epsilon_rate, 2, 64)

	fmt.Println(gamma_rate_number * epsilon_rate_number)
}

func oxygen_rating(ratings []string, line_len int) string {
	for i := 0; i < line_len; i++ {
		oxygen_ratings := split(ratings, i)

		if len(oxygen_ratings.ones) == len(oxygen_ratings.zeros) {
			ratings = oxygen_ratings.ones
		} else if len(oxygen_ratings.ones) > len(oxygen_ratings.zeros) {
			ratings = oxygen_ratings.ones
		} else {
			ratings = oxygen_ratings.zeros
		}

		if len(ratings) == 1 {
			break
		}
	}

	return ratings[0]
}

type split_strings struct {
	ones  []string
	zeros []string
}

func split(strings []string, pos int) split_strings {
	ones := []string{}
	zeros := []string{}

	for _, string := range strings {
		if string[pos] == '1' {
			ones = append(ones, string)
		} else {
			zeros = append(zeros, string)
		}
	}

	return split_strings{
		ones, zeros,
	}
}

func co2_rating(ratings []string, line_len int) string {
	for i := 0; i < line_len; i++ {
		co2_ratings := split(ratings, i)

		if len(co2_ratings.ones) == len(co2_ratings.zeros) {
			ratings = co2_ratings.zeros
		} else if len(co2_ratings.ones) > len(co2_ratings.zeros) {
			ratings = co2_ratings.zeros
		} else {
			ratings = co2_ratings.ones
		}

		if len(ratings) == 1 {
			break
		}
	}

	return ratings[0]
}

func part2(lines []string, line_len int) {
	oxygen_rating := oxygen_rating(lines, line_len)
	co2_rating := co2_rating(lines, line_len)

	oxygen_rating_number, _ := strconv.ParseInt(oxygen_rating, 2, 64)
	co2_rating_number, _ := strconv.ParseInt(co2_rating, 2, 64)

	fmt.Println(oxygen_rating_number * co2_rating_number)
}

func main() {
	bytes, err := os.ReadFile("day03.input")
	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}
	lines := strings.Split(string(bytes), "\n")
	line_len := len(lines[0])

	part1(lines, line_len)
	part2(lines, line_len)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// It seems like the submarine can take a series of commands like forward 1,
	// down 2, or up 3:
	//
	// - forward X increases the horizontal position by X units.
	// - down X increases the depth by X units.
	// - up X decreases the depth by X units.
	lines := readInputLines()

	// Calculate the horizontal position and depth you would have after following
	// the planned course.
	fmt.Println(problemA(lines))

	// The commands also mean something entirely different than you first thought:
	//
	// - down X increases your aim by X units.
	// - up X decreases your aim by X units.
	// - forward X does two things:
	//   - It increases your horizontal position by X units.
	//   - It increases your depth by your aim multiplied by X.
	fmt.Println(problemB(lines))
}

func readInputLines() (lines []string) {
	file, err := os.Open("input/day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

// What do you get if you multiply your final horizontal position by your
// final depth?
func problemA(input []string) int {
	horisontal := 0
	depth := 0
	for _, command := range input {
		parts := strings.Fields(command)
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch parts[0] {
		case "forward":
			horisontal += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}
	return horisontal * depth
}

// What do you get if you multiply your final horizontal position by your
// final depth?
func problemB(input []string) int {
	horisontal := 0
	depth := 0
	aim := 0
	for _, command := range input {
		parts := strings.Fields(command)
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch parts[0] {
		case "forward":
			horisontal += amount
			depth += aim * amount
		case "up":
			aim -= amount
		case "down":
			aim += amount
		}
	}
	return horisontal * depth
}

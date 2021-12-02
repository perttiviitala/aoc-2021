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
	// down 2, or up 3
	lines := readInputLines()

	// What do you get if you multiply your final horizontal position by your
	// final depth?
	if horisontal, depth := problemA(lines); true {
		fmt.Println(horisontal * depth)
	}

	// Based on your calculations, the planned course doesn't seem to make any
	// sense. You find the submarine manual and discover that the process is
	// actually slightly more complicated.
	if horisontal, depth := problemB(lines); true {
		fmt.Println(horisontal * depth)
	}
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

// Calculate the horizontal position and depth you would have after following
// the planned course.
//
// - forward X increases the horizontal position by X units.
// - down X increases the depth by X units.
// - up X decreases the depth by X units.
func problemA(input []string) (horisontal, depth int) {
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
	return
}

// Calculate the horizontal position and depth you would have after following
// the planned course.
//
// - down X increases your aim by X units.
// - up X decreases your aim by X units.
// - forward X does two things:
//   - It increases your horizontal position by X units.
//   - It increases your depth by your aim multiplied by X.
func problemB(input []string) (horisontal, depth int) {
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
	return
}

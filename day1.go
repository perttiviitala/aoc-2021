package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Each line is a measurement of the sea floor depth as the sweep looks further
	// and further away from the submarine.
	lines := readInputLines()

	// The first order of business is to figure out how quickly the depth
	// increases, just so you know what you're dealing with.
	//
	// To do this, count the number of times a depth measurement increases from
	// the previous measurement. (There is no measurement before the first
	// measurement.)
	fmt.Println("Increases", problemA(lines))

	// Your goal now is to count the number of times the sum of measurements in
	// this sliding window increases from the previous sum. So, compare A with B,
	// then compare B with C, then C with D, and so on. Stop when there aren't
	// enough measurements left to create a new three-measurement sum.
	fmt.Println("After grouping", problemB(lines))
}

func readInputLines() (lines []int) {
	file, err := os.Open("input/day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		lines = append(lines, int(value))
	}
	return
}

// How many measurements are larger than the previous measurement?
func problemA(input []int) (counter int) {
	previous := 0
	for _, measurement := range input {
		// First iteration is not expected to match.
		if previous > 0 && measurement > previous {
			counter++
		}
		previous = measurement
	}
	return
}

// How many sums are larger than the previous sum?
func problemB(input []int) (counter int) {
	previous := 0
	for index, measurement := range input {
		// Quarantees that index-1 and index-2 exists.
		if index < 2 {
			continue
		}
		sum := measurement + input[index-1] + input[index-2]
		// First iteration is not expected to match.
		if previous > 0 && sum > previous {
			counter++
		}
		previous = sum
	}
	return
}

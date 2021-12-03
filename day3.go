package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// The diagnostic report consists of a list of binary numbers which, when
	// decoded properly, can tell you many useful things about the conditions of
	// the submarine.
	lines := readInputLines()

	// What is the power consumption of the submarine?
	gamma, epsilon := problemA(lines)
	fmt.Println(gamma * epsilon)

	// Next, you should verify the life support rating.
	oxygen, scrubber := problemB(lines)
	fmt.Println(oxygen * scrubber)
}

func readInputLines() (lines []string) {
	file, err := os.Open("input/day3.txt")
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

// Each bit in the gamma rate can be determined by finding the most common bit in
// the corresponding position of all numbers in the diagnostic report.
//
// The epsilon rate is calculated in a similar way; rather than use the most
// common bit, the least common bit from each position is used.
func problemA(input []string) (gamma, epsilon int) {
	for index := range input[0] {
		// Move bits to left.
		gamma <<= 1
		epsilon <<= 1
		// Check wether to increase gamma or epsilon.
		most, least := mostAndLeastCommonByIndex(input, index)
		gamma |= int(most - '0')
		epsilon |= int(least - '0')
	}
	return
}

// Both the oxygen generator rating and the CO2 scrubber rating are values that
// can be found in your diagnostic report - finding them is the tricky part. Both
// values are located using a similar process that involves filtering out values
// until only one remains. Before searching for either rating value, start with
// the full list of binary numbers from your diagnostic report and consider just
// the first bit of those numbers.
func problemB(input []string) (oxygen, scrubber int64) {
	var err error

	oxygen, err = strconv.ParseInt(filterRating(input, 0, true), 2, 64)
	if err != nil {
		panic(err)
	}

	scrubber, err = strconv.ParseInt(filterRating(input, 0, false), 2, 64)
	if err != nil {
		panic(err)
	}

	return
}

func mostAndLeastCommonByIndex(input []string, index int) (rune, rune) {
	var ones, zeros int
	for _, line := range input {
		if line[index] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	if ones >= zeros {
		return '1', '0'
	}
	return '0', '1'
}

func filterRating(lines []string, index int, useMostCommon bool) string {
	most, least := mostAndLeastCommonByIndex(lines, index)
	wanted := least
	if useMostCommon {
		wanted = most
	}

	filtered := make([]string, 0)
	for _, line := range lines {
		if rune(line[index]) == wanted {
			filtered = append(filtered, line)
		}
	}

	// As there is only one number left, stop.
	if len(filtered) == 1 {
		return filtered[0]
	}
	// Keep filtering, panics if no single result found.
	return filterRating(filtered, index+1, useMostCommon)
}

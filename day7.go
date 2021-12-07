package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	positions := readInputLine()
	fmt.Println(problemA(positions))
	fmt.Println(problemB(positions))
}

func readInputLine() (positions []int) {
	file, err := os.Open("input/day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, text := range strings.Split(scanner.Text(), ",") {
			position, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			positions = append(positions, position)
		}
	}
	return
}

func problemA(positions []int) int {
	max := 0
	for _, position := range positions {
		if position > max {
			max = position
		}
	}
	sums := make(map[int]int, max)
	for i := 0; i < max; i++ {
		for _, position := range positions {
			sums[i] += int(math.Abs(float64(position - i)))
		}
	}
	min := 0
	for _, sum := range sums {
		if min == 0 || sum < min {
			min = sum
		}
	}
	return min
}

func problemB(positions []int) int {
	max := 0
	for _, position := range positions {
		if position > max {
			max = position
		}
	}
	sums := make(map[int]int, max)
	for i := 0; i < max; i++ {
		for _, position := range positions {
			distance := int(math.Abs(float64(position - i)))
			sums[i] += distance * (distance + 1) / 2
		}
	}
	min := 0
	for _, sum := range sums {
		if min == 0 || sum < min {
			min = sum
		}
	}
	return min
}

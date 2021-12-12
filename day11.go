package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(problemA(readInputLines()))
	fmt.Println(problemB(readInputLines()))
}

func readInputLines() (lines grid) {
	file, err := os.Open("input/day11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := make([]int, 0)
		for _, char := range scanner.Text() {
			val, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, val)
		}
		lines = append(lines, line)
	}
	return
}

func problemA(grid grid) (total int) {
	for i := 0; i < 100; i++ {
		total += grid.Step()
	}
	return
}

func problemB(grid grid) int {
	return stepUntilFlashing(grid, 1)
}

func stepUntilFlashing(grid grid, index int) int {
	grid.Step()
	if grid.IsAllFlashing() {
		return index
	}
	return stepUntilFlashing(grid, index+1)
}

type grid [][]int

func (grid grid) Step() int {
	initial := grid.IncreaseGrid(0, 0, len(grid))
	flashes := grid.Flash(initial, map[[2]int]bool{})
	for point := range flashes {
		grid[point[0]][point[1]] = 0
	}
	return len(flashes)
}

func (grid grid) IncreaseGrid(y, x, size int) (flashes [][2]int) {
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			if y+j < 0 || x+i < 0 {
				continue
			}
			if y+j > len(grid)-1 || x+i > len(grid[0])-1 {
				continue
			}
			val := grid[y+j][x+i]
			grid[y+j][x+i]++
			if val == 9 {
				flashes = append(flashes, [2]int{y + j, x + i})
			}
		}
	}
	return
}

func (grid grid) Flash(current [][2]int, flashes map[[2]int]bool) map[[2]int]bool {
	next := make([][2]int, 0)
	for _, point := range current {
		flashes[[2]int{point[0], point[1]}] = true
		for _, found := range grid.IncreaseGrid(point[0]-1, point[1]-1, 3) {
			if _, ok := flashes[[2]int{found[0], found[1]}]; ok {
				// Already flashed
				continue
			}
			next = append(next, [2]int{found[0], found[1]})
		}
	}
	if len(next) > 0 {
		return grid.Flash(next, flashes)
	}
	return flashes
}

func (grid grid) IsAllFlashing() bool {
	for _, line := range grid {
		for _, octopus := range line {
			if octopus > 0 {
				return false
			}
		}
	}
	return true
}

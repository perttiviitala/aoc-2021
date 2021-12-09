package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	cube := readInputLines()
	fmt.Println(problemA(cube))
	fmt.Println(problemB(cube))
}

func readInputLines() (lines [][]int) {
	file, err := os.Open("input/day9.txt")
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

func problemA(cube cube) (total int) {
	for _, point := range cube.SolveLowPoints() {
		total += point.value + 1
	}
	return
}

func problemB(cube cube) (total int) {
	sizes := cube.SolveBasinSizes()
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	total += sizes[0]
	for _, size := range sizes[1:3] {
		total *= size
	}
	return
}

type cube [][]int

func (cube cube) MaxY() int {
	return len(cube) - 1
}

func (cube cube) MaxX() int {
	return len(cube[0]) - 1
}

type point struct {
	y, x, value int
}

func (point point) String() string {
	return fmt.Sprintf("%v,%v", point.y, point.x)
}

func (cube cube) SolveLowPoints() (lowPoints []point) {
	for y, line := range cube {
		for x, value := range line {
			if x > 0 && cube[y][x-1] <= value {
				continue
			}
			if x < cube.MaxX() && cube[y][x+1] <= value {
				continue
			}
			if y > 0 && cube[y-1][x] <= value {
				continue
			}
			if y < cube.MaxY() && cube[y+1][x] <= value {
				continue
			}
			lowPoints = append(lowPoints, point{y, x, value})
		}
	}
	return
}

func (cube cube) SolveBasinSizes() (basins []int) {
	for _, lowPoint := range cube.SolveLowPoints() {
		acc := make(map[string]int)
		_, acc = addBasins(cube, []point{lowPoint}, acc)
		basins = append(basins, len(acc))
	}
	return
}

func addBasins(cube cube, basins []point, acc map[string]int) ([]point, map[string]int) {
	expanded := make([]point, 0)
	for _, basin := range basins {
		acc[basin.String()] = basin.value

		if basin.x > 0 {
			if left := cube[basin.y][basin.x-1]; left > basin.value && left < 9 {
				expanded = addUnique(expanded, point{basin.y, basin.x - 1, 0}, acc)
			}
		}
		if basin.x < cube.MaxX() {
			if right := cube[basin.y][basin.x+1]; right > basin.value && right < 9 {
				expanded = addUnique(expanded, point{basin.y, basin.x + 1, 0}, acc)
			}
		}
		if basin.y > 0 {
			if left := cube[basin.y-1][basin.x]; left > basin.value && left < 9 {
				expanded = addUnique(expanded, point{basin.y - 1, basin.x, 0}, acc)
			}
		}
		if basin.y < cube.MaxY() {
			if right := cube[basin.y+1][basin.x]; right > basin.value && right < 9 {
				expanded = addUnique(expanded, point{basin.y + 1, basin.x, 0}, acc)
			}
		}
	}
	if len(expanded) > 0 {
		return addBasins(cube, expanded, acc)
	}
	return expanded, acc
}

func addUnique(points []point, point point, acc map[string]int) []point {
	if _, ok := acc[point.String()]; !ok {
		points = append(points, point)
	}
	return points
}

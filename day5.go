package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Each line of vents is given as a line segment in the format x1,y1 -> x2,y2
	// where x1,y1 are the coordinates of one end the line segment and x2,y2 are
	// the coordinates of the other end. These line segments include the points
	// at both ends.
	lines := readInputLines()

	// Only counting horizontal and vertial lines.
	fmt.Println(problemA(lines))

	// Counting all lines.
	fmt.Println(problemB(lines))
}

func readInputLines() (lines []string) {
	file, err := os.Open("input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

type Point struct {
	x, y int
}

func makePoint(pair string) Point {
	cords := strings.Split(pair, ",")
	x, err := strconv.Atoi(string(cords[0]))
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(string(cords[1]))
	if err != nil {
		log.Fatal(err)
	}
	return Point{x, y}
}

type Line struct {
	start, end Point
}

func (line *Line) IsStraight() bool {
	return line.start.x == line.end.x || line.start.y == line.end.y
}

func (line *Line) Points() (points []Point) {
	if line.IsStraight() {
		x1, x2 := minMaxInts(line.start.x, line.end.x)
		y1, y2 := minMaxInts(line.start.y, line.end.y)
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				points = append(points, Point{x, y})
			}
		}
		return
	}
	// There has to be a better way!
	x1, x2 := line.start.x, line.end.x
	y1, y2 := line.start.y, line.end.y
	if x1 > x2 {
		if y1 > y2 {
			for x, y := x1, y1; x >= x2 && y >= y2; x, y = x-1, y-1 {
				points = append(points, Point{x, y})
			}
		} else {
			for x, y := x1, y1; x >= x2 && y <= y2; x, y = x-1, y+1 {
				points = append(points, Point{x, y})
			}
		}
	} else {
		if y1 > y2 {
			for x, y := x1, y1; x <= x2 && y >= y2; x, y = x+1, y-1 {
				points = append(points, Point{x, y})
			}
		} else {
			for x, y := x1, y1; x <= x2 && y <= y2; x, y = x+1, y+1 {
				points = append(points, Point{x, y})
			}
		}
	}
	return
}

func makeLine(line string) Line {
	sides := strings.Fields(line)
	return Line{
		makePoint(sides[0]),
		makePoint(sides[2]),
	}
}

type Grid struct {
	values [][]int
}

func makeGrid(size int) Grid {
	grid := Grid{make([][]int, size)}
	for i := range grid.values {
		grid.values[i] = make([]int, size)
	}
	return grid
}

func (grid *Grid) MarkPath(line Line) {
	for _, point := range line.Points() {
		grid.values[point.y][point.x] += 1
	}
}

func (grid *Grid) OverlapCount() (counter int) {
	for _, line := range grid.values {
		for _, number := range line {
			if number > 1 {
				counter++
			}
		}
	}
	return
}

func minMaxInts(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func problemA(input []string) int {
	grid := makeGrid(1000)
	for _, line := range input {
		ventLine := makeLine(line)
		if ventLine.IsStraight() {
			grid.MarkPath(ventLine)
		}
	}
	return grid.OverlapCount()
}

func problemB(input []string) int {
	grid := makeGrid(1000)
	for _, line := range input {
		ventLine := makeLine(line)
		grid.MarkPath(ventLine)
	}
	return grid.OverlapCount()
}

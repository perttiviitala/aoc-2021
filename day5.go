package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func readInputLines() (lines []Line) {
	file, err := os.Open("input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, Line{
			Point{x1, y1},
			Point{x2, y2},
		})
	}
	return
}

type Point struct {
	x, y int
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
	x1, x2 := line.start.x, line.end.x
	y1, y2 := line.start.y, line.end.y
	if x1 > x2 && y1 > y2 {
		for i := 0; i <= x1-x2; i++ {
			points = append(points, Point{x1 - i, y1 - i})
		}
	} else if x1 > x2 && y1 < y2 {
		for i := 0; i <= x1-x2; i++ {
			points = append(points, Point{x1 - i, y1 + i})
		}
	} else if x1 < x2 && y1 > y2 {
		for i := 0; i <= x2-x1; i++ {
			points = append(points, Point{x1 + i, y1 - i})
		}
	} else if x1 < x2 && y1 < y2 {
		for i := 0; i <= x2-x1; i++ {
			points = append(points, Point{x1 + i, y1 + i})
		}
	}
	return
}

type Grid struct {
	values [][]int
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

func makeGrid(lines []Line) Grid {
	maxX, maxY := 0, 0
	for _, line := range lines {
		if _, x := minMaxInts(line.start.x, line.end.x); x > maxX {
			maxX = x
		}
		if _, y := minMaxInts(line.start.y, line.end.y); y > maxY {
			maxY = y
		}
	}
	grid := Grid{make([][]int, maxY+1)}
	for i := range grid.values {
		grid.values[i] = make([]int, maxX+1)
	}
	return grid
}

func minMaxInts(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func problemA(lines []Line) int {
	grid := makeGrid(lines)
	for _, line := range lines {
		if line.IsStraight() {
			grid.MarkPath(line)
		}
	}
	return grid.OverlapCount()
}

func problemB(lines []Line) int {
	grid := makeGrid(lines)
	for _, line := range lines {
		grid.MarkPath(line)
	}
	return grid.OverlapCount()
}

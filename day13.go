package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(problemA(readInputLines()))
	fmt.Println(problemB(readInputLines()))
}

func readInputLines() (paper paper, folds []fold) {
	file, err := os.Open("input/day13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	breakpoint := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			breakpoint = true
			continue
		}

		if breakpoint {
			var fold fold
			fmt.Sscanf(line, "fold along %1s=%d", &fold.dir, &fold.amount)
			folds = append(folds, fold)
		} else {
			var x, y int
			fmt.Sscanf(line, "%d,%d", &x, &y)
			paper = append(paper, [2]int{x, y})
		}
	}
	return
}

func problemA(paper paper, folds []fold) int {
	paper.Fold(folds[0])
	return paper.PointCount()
}

func problemB(paper paper, folds []fold) string {
	for _, fold := range folds {
		paper.Fold(fold)
	}
	return paper.String()
}

type paper [][2]int

func (paper paper) String() string {
	maxX, maxY := 0, 0
	exists := make(map[[2]int]bool)
	for _, point := range paper {
		if point[0] > maxX {
			maxX = point[0]
		}
		if point[1] > maxY {
			maxY = point[1]
		}
		exists[point] = true
	}
	out := ""
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if _, ok := exists[[2]int{x, y}]; ok {
				out += "#"
				continue
			}
			out += "."
		}
		out += "\n"
	}
	return out
}

func (paper paper) Fold(fold fold) {
	switch fold.dir {
	case "x":
		paper.FoldX(fold.amount)
	case "y":
		paper.FoldY(fold.amount)
	}
}

func (paper paper) FoldX(at int) {
	for i, point := range paper {
		if point[0] > at {
			paper[i][0] -= 2 * (point[0] - at)
		}
	}
}

func (paper paper) FoldY(at int) {
	for i, point := range paper {
		if point[1] > at {
			paper[i][1] -= 2 * (point[1] - at)
		}
	}
}

func (paper paper) PointCount() int {
	unique := make(map[[2]int]bool)
	for _, point := range paper {
		unique[point] = true
	}
	return len(unique)
}

type fold struct {
	amount int
	dir    string
}

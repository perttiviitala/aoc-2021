package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	lines := readInputLines()
	fmt.Println(problemA(lines))
	fmt.Println(problemB(lines))
}

var closing = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}

type line string

func (l line) IsValid() (bool, rune, rune) {
	expected := make([]rune, 0)
	for _, char := range l {
		switch char {
		case '(', '[', '{', '<':
			expected = append(expected, closing[char])
		case ')', ']', '}', '>':
			i := len(expected) - 1
			if char != expected[i] {
				return false, expected[i], char
			}
			expected = expected[:i]
		}
	}
	return true, rune(0), rune(0)
}

func (l line) Expected() string {
	expected := make([]rune, 0)
	for _, char := range l {
		switch char {
		case '(', '[', '{', '<':
			expected = append(expected, closing[char])
		case ')', ']', '}', '>':
			i := len(expected) - 1
			if char == expected[i] {
				expected = expected[:i]
			}
		}
	}
	out := ""
	for _, char := range string(expected) {
		out = string(char) + out
	}
	return out
}

func readInputLines() (lines []line) {
	file, err := os.Open("input/day10.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, line(scanner.Text()))
	}
	return
}

func problemA(lines []line) (total int) {
	points := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	for _, line := range lines {
		if valid, _, found := line.IsValid(); !valid {
			total += points[found]
		}
	}
	return
}

func problemB(lines []line) int {
	points := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	scores := make([]int, 0)
	for _, line := range lines {
		if valid, _, _ := line.IsValid(); !valid {
			continue
		}
		score := 0
		for _, char := range line.Expected() {
			score = (score * 5) + points[char]
		}
		scores = append(scores, score)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	return scores[len(scores)/2]
}

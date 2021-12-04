package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines()

	// What will your final score be if you choose first winning board?
	fmt.Println(problemA(lines))

	// What will your final score be if you choose last winning board?
	fmt.Println(problemB(lines))
}

func readInputLines() (lines []string) {
	file, err := os.Open("input/day4.txt")
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

func problemA(input []string) int {
	board, number := winningBoard(
		parseBoards(input[1:]),
		mapAsInt(strings.Split(input[0], ",")),
	)
	return board.Score() * number
}

func problemB(input []string) int {
	board, number := filterWinning(
		parseBoards(input[1:]),
		mapAsInt(strings.Split(input[0], ",")),
	)
	return board.Score() * number
}

type Board struct {
	lines [5]BoardLine
}
type BoardLine struct {
	numbers [5]LineNumber
}
type LineNumber struct {
	number  int
	checked bool
}

func (board *Board) MarkNumber(value int) {
	for i, line := range board.lines {
		for j, number := range line.numbers {
			if value == number.number {
				board.lines[i].numbers[j].checked = true
			}
		}
	}
	return
}

func (board *Board) IsWinner() bool {
	// One line is fully checked.
	for _, line := range board.lines {
		counter := 0
		for _, number := range line.numbers {
			if !number.checked {
				break
			}
			counter++
		}
		if counter == 5 {
			return true
		}
	}
	// One column is fully checked.
	for i := 0; i < 5; i++ {
		counter := 0
		for _, line := range board.lines {
			if !line.numbers[i].checked {
				break
			}
			counter++
		}
		if counter == 5 {
			return true
		}
	}
	return false
}

func (board *Board) Score() (score int) {
	for _, line := range board.lines {
		for _, number := range line.numbers {
			if !number.checked {
				score += number.number
			}
		}
	}
	return
}

func mapAsInt(input []string) (output []int) {
	for _, str := range input {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		output = append(output, number)
	}
	return
}

func parseBoards(input []string) (boards []Board) {
	// Before each board there is an empty line.
	boards = make([]Board, len(input)/6)
	for i := range boards {
		for j, line := range input[i*6+1 : i*6+6] {
			for k, number := range mapAsInt(strings.Fields(line)) {
				boards[i].lines[j].numbers[k].number = number
			}
		}
	}
	return
}

func winningBoard(boards []Board, draw []int) (winner Board, number int) {
	for _, number := range draw {
		for i, board := range boards {
			board.MarkNumber(number)
			if board.IsWinner() {
				return board, number
			}
			boards[i] = board
		}
	}
	panic("No winner")
}

func filterWinning(boards []Board, draw []int) (Board, int) {
	if len(boards) == 1 {
		// We know winning board but at what number.
		board := boards[0]
		for _, number := range draw {
			board.MarkNumber(number)
			if board.IsWinner() {
				return board, number
			}
		}
	}
	return filterWinning(
		filterWinners(boards, draw[0]),
		draw[1:],
	)
	panic("No single last winner")
}

func filterWinners(boards []Board, number int) (filtered []Board) {
	for _, board := range boards {
		board.MarkNumber(number)
		if !board.IsWinner() {
			filtered = append(filtered, board)
		}
	}
	return
}

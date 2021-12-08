package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	segments := readInputLines()
	fmt.Println(problemA(segments))
	fmt.Println(problemB(segments))
}

func readInputLines() (segments [][]string) {
	file, err := os.Open("input/day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		line := make([]string, 0)
		line = append(line, fields[:10]...)
		line = append(line, fields[11:]...)
		segments = append(segments, line)
	}
	return
}

func problemA(segments [][]string) (total int) {
	for _, segment := range segments {
		for _, digit := range segment[10:] {
			switch len(digit) {
			case 2, 3, 4, 7:
				total++
			}
		}
	}
	return
}

func problemB(segments [][]string) (total int) {
	for _, segment := range segments {
		sorted := make([]string, 0)
		for _, digit := range segment {
			sorted = append(sorted, sortDigit(digit))
		}
		total += segmentOutput(sorted[:10], sorted[10:])
	}
	return
}

func segmentOutput(input, output []string) int {
	digits := solveDigits(input)

	totalStr := ""
	for _, digit := range output {
		totalStr += strconv.Itoa(digits[digit])
	}
	total, err := strconv.Atoi(totalStr)
	if err != nil {
		log.Fatal(err)
	}
	return total
}

func solveDigits(digits []string) map[string]int {
	var zero, one, two, three, four, five, six, seven, eight, nine string

	for _, digit := range digits {
		switch len(digit) {
		case 2:
			one = digit
		case 4:
			four = digit
		case 3:
			seven = digit
		case 7:
			eight = digit
		}
	}
	for _, digit := range digits {
		switch len(digit) {
		case 6:
			if overlapCount(digit, one) == 1 {
				six = digit
				continue
			}
			if overlapCount(digit, four) == 3 {
				zero = digit
				continue
			}
			nine = digit
		}
	}
	for _, digit := range digits {
		switch len(digit) {
		case 5:
			if overlapCount(digit, one) == 2 {
				three = digit
				continue
			}
			if overlapCount(digit, nine) == 5 {
				five = digit
				continue
			}
			two = digit
		}
	}

	return map[string]int{
		zero:  0,
		one:   1,
		two:   2,
		three: 3,
		four:  4,
		five:  5,
		six:   6,
		seven: 7,
		eight: 8,
		nine:  9,
	}
}

func sortDigit(input string) string {
	stringArr := make([]string, 0)
	for _, rune := range input {
		stringArr = append(stringArr, string(rune))
	}
	sort.Sort(sort.StringSlice(stringArr))
	return strings.Join(stringArr, "")
}

func overlapCount(digit, compare string) (count int) {
	for _, c := range compare {
		if strings.ContainsRune(digit, c) {
			count++
		}
	}
	return
}

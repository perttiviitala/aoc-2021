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
	fishAges := readInputLine()
	fmt.Println(problemA(fishAges))
	fmt.Println(problemB(fishAges))
}

func readInputLine() (lines []int) {
	file, err := os.Open("input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, text := range strings.Split(scanner.Text(), ",") {
			age, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			lines = append(lines, age)
		}
	}
	return
}

func problemA(fishAges []int) (total int) {
	counts := make(map[int]int)
	for _, age := range fishAges {
		counts[age] += 1
	}
	counts = nextNDays(counts, 80)
	for _, count := range counts {
		total += count
	}
	return
}

func problemB(fishAges []int) (total int) {
	counts := make(map[int]int)
	for _, age := range fishAges {
		counts[age] += 1
	}
	counts = nextNDays(counts, 256)
	for _, count := range counts {
		total += count
	}
	return
}

func nextNDays(counts map[int]int, days int) map[int]int {
	for day := 0; day < days; day++ {
		counts = nextMorning(counts)
	}
	return counts
}

func nextMorning(today map[int]int) map[int]int {
	tomorrow := make(map[int]int)
	for day, count := range today {
		if count == 0 {
			continue
		}
		if day == 0 {
			tomorrow[6] += count
			tomorrow[8] += count
			continue
		}
		tomorrow[day-1] += count
	}
	return tomorrow
}

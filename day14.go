package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(problemA(readPolymer()))
	fmt.Println(problemB(readPolymer()))
}

func readPolymer() polymer {
	file, err := os.Open("input/day14.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pairs := make(map[string]int)
	rules := make([]rule, 0)
	counts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(pairs) == 0 {
			counts[string(line[0])]++
			for i := 1; i < len(line); i++ {
				pairs[string(line[i-1])+string(line[i])]++
				counts[string(line[i])]++
			}
			continue
		}
		if line == "" {
			continue
		}

		var rule rule
		fmt.Sscanf(line, "%c%c -> %c", &rule.left, &rule.right, &rule.insert)
		rules = append(rules, rule)
	}

	return polymer{rules, pairs, counts}
}

func problemA(polymer polymer) int {
	for i := 0; i < 10; i++ {
		polymer.Step()
	}
	min, max := polymer.MinMax()
	return max - min
}

func problemB(polymer polymer) int {
	for i := 0; i < 40; i++ {
		polymer.Step()
	}
	min, max := polymer.MinMax()
	return max - min
}

type rule struct {
	left   rune
	right  rune
	insert rune
}

func (rule rule) Pair() (output string) {
	output += string(rule.left)
	output += string(rule.right)
	return
}

type polymer struct {
	rules  []rule
	pairs  map[string]int
	counts map[string]int
}

func (polymer polymer) Step() {
	counts := make(map[string]int)
	for _, rule := range polymer.rules {
		if count, ok := polymer.pairs[rule.Pair()]; ok {
			delete(polymer.pairs, rule.Pair())
			counts[string(rule.left)+string(rule.insert)] += count
			counts[string(rule.insert)+string(rule.right)] += count
			polymer.counts[string(rule.insert)] += count
		}
	}
	for pair, count := range counts {
		polymer.pairs[pair] = count
	}
}

func (polymer polymer) MinMax() (min, max int) {
	for _, count := range polymer.counts {
		if count < min || min == 0 {
			min = count
		}
		if count > max {
			max = count
		}
	}
	return
}

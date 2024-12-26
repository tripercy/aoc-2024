package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafe(levels []int) bool {
	asc, desc := true, true
	for i := range len(levels) - 1 {
		level := levels[i]
		next := levels[i+1]
		if level < next {
			desc = false
		}
		if level > next {
			asc = false
		}
		diff := level - next
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			desc = false
			asc = false
			break
		}
	}
	return asc || desc
}

func partOne(input [][]int) {
	res := 0

	for _, levels := range input {
		if isSafe(levels) {
			res++
		}
	}

	fmt.Println("Part 1:", res)
}

func partTwo(input [][]int) {
	res := 0

	for _, levels := range input {
		for i := range levels {
			newLevels := make([]int, i)
			copy(newLevels, levels[:i])
			newLevels = append(newLevels, levels[i+1:]...)
			if isSafe(newLevels) {
				res++
				break
			}
		}
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		levelStr := strings.Fields(line)
		levels := make([]int, len(levelStr))
		for i, lev := range levelStr {
			if level, err := strconv.Atoi(lev); err != nil {
				log.Fatal(err)
			} else {
				levels[i] = level
			}
		}
		input = append(input, levels)
	}
	partOne(input)
	partTwo(input)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var dr []int = []int{0, 0, 1, -1}
var dc []int = []int{1, -1, 0, 0}

func inBound(r, c, nr, nc int) bool {
	return r >= 0 && c >= 0 && r < nr && c < nc
}

func walk(input [][]byte, visited [][]bool, r, c int, checkRate bool) int {
	nr, nc := len(input), len(input[0])
	val := input[r][c]
	if val == 9 {
		if !visited[r][c] || checkRate {
			visited[r][c] = true
			return 1
		}
		return 0
	}
	res := 0
	for d := range 4 {
		tr, tc := r+dr[d], c+dc[d]
		if inBound(tr, tc, nr, nc) && input[tr][tc] == val+1 {
			res += walk(input, visited, tr, tc, checkRate)
		}
	}

	return res
}

func partOne(input [][]byte) {
	res := 0
	nr, nc := len(input), len(input[0])

	for r, line := range input {
		for c, val := range line {
			if val == 0 {
				visited := make([][]bool, nr)
				for i := range nr {
					visited[i] = make([]bool, nc)
				}
				score := walk(input, visited, r, c, false)
				res += score
			}
		}
	}

	fmt.Println("Part 1:", res)
}

func partTwo(input [][]byte) {
	res := 0
	nr, nc := len(input), len(input[0])

	for r, line := range input {
		for c, val := range line {
			if val == 0 {
				visited := make([][]bool, nr)
				for i := range nr {
					visited[i] = make([]bool, nc)
				}
				score := walk(input, visited, r, c, true)
				res += score
			}
		}
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([][]byte, 0)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		for i, x := range line {
			line[i] = x - '0'
		}
		input = append(input, line)
	}

	partOne(input)
	partTwo(input)
}

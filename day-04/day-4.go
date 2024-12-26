package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func partOne(input []string) {
	res := 0

	dr := []int{0, 0, 1, 1, 1, -1, -1, -1}
	dc := []int{1, -1, 1, 0, -1, 1, 0, -1}

	for r := range len(input) {
		for c := range len(input[r]) {
			if input[r][c] != 'X' {
				continue
			}

			for d := range 8 {
				flag := true
				for i, ch := range "XMAS" {
					nr := r + dr[d]*i
					nc := c + dc[d]*i
					if nr < 0 || nc < 0 || nr >= len(input) || nc >= len(input[r]) || input[nr][nc] != byte(ch) {
						flag = false
						break
					}
				}
				if flag {
					res++
				}
			}
		}
	}

	fmt.Println("Part 1:", res)
}

func partTwo(input []string) {
	res := 0
	for r := 1; r < len(input)-1; r++ {
		for c := 1; c < len(input[r])-1; c++ {
			if input[r][c] != 'A' {
				continue
			}

			diag1 := string([]byte{input[r-1][c-1], input[r][c], input[r+1][c+1]})
			diag2 := string([]byte{input[r-1][c+1], input[r][c], input[r+1][c-1]})
			if diag1 != "MAS" && diag1 != "SAM" {
				continue
			}
			if diag2 != "MAS" && diag2 != "SAM" {
				continue
			}
			res++
		}
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	partOne(input)
	partTwo(input)
}

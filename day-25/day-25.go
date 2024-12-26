package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func overlap(lock, key []int) bool {
	for i := range lock {
		if lock[i]+key[i] > 5 {
			return true
		}
	}
	return false
}

func partOne(locks, keys [][]int) {
	res := 0

	for _, lock := range locks {
		for _, key := range keys {
			if !overlap(lock, key) {
				res++
			}
		}
	}

	fmt.Println("Part 1:", res)
}

func main() {
	file, err := os.Open("day-25.txt")
	// file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	locks := make([][]int, 0)
	keys := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		isLock := true
		if scanner.Text()[0] == '.' {
			isLock = false
		}

		heights := make([]int, 5)
		for range 5 {
			scanner.Scan()
			for i, ch := range scanner.Text() {
				if ch == '#' {
					heights[i]++
				}
			}
		}
		scanner.Scan()
		scanner.Scan()

		if isLock {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	partOne(locks, keys)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const mod = 16777216

func stoi64(s string) uint64 {
	res, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func partOne(input []uint64) {

	res := uint64(0)
	for _, num := range input {
		x := num
		for range 2000 {
			x = (x ^ (x << 6)) % mod
			x = (x ^ (x >> 5)) % mod
			x = (x ^ (x << 11)) % mod
		}

		res += x
	}
	fmt.Println("Part 1:", res)
}

func hash(seq []int) int {
	res := 0
	for _, x := range seq {
		x += 9
		res = res*18 + x
	}
	return res
}

func partTwo(input []uint64) {
	prices := make([]int, 2001)
	diff := make([]int, 2001)

	bananas := make(map[int]int)

	for _, num := range input {
		x := num
		prices[0] = int(x % 10)
		for i := range 2000 {
			x = (x ^ (x << 6)) % mod
			x = (x ^ (x >> 5)) % mod
			x = (x ^ (x << 11)) % mod
			prices[i+1] = int(x % 10)
			diff[i+1] = prices[i+1] - prices[i]
		}
		appeared := make([]bool, 2e6)
		for i := 5; i < 2001; i++ {
			seq := diff[i-4 : i]
			h := hash(seq)
			if !appeared[h] {
				appeared[h] = true
				bananas[h] += prices[i-1]
			}
		}
	}
	res := 0
	for _, cnt := range bananas {
		res = max(res, cnt)
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-22.txt")
	// file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([]uint64, 0)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		input = append(input, stoi64(scanner.Text()))
	}

	partOne(input)
	partTwo(input)
}

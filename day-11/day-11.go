package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func stoi(s string) int {
	if num, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
	} else {
		return num
	}

	return -1
}

func cntDigit(n int) int {
	res := 0
	for n > 0 {
		res++
		n /= 10
	}
	return res
}

func tenPow(x int) int {
	res := 1
	for x > 0 {
		x--
		res *= 10
	}
	return res
}

func blink(freq map[int]int) {
	adding := make(map[int]int)
	for num, cnt := range freq {
		if cnt == 0 {
			continue
		}
		if num == 0 {
			adding[0] -= cnt
			adding[1] += cnt
		} else if digits := cntDigit(num); digits%2 == 0 {
			a := num / tenPow(digits/2)
			b := num % tenPow(digits/2)
			adding[a] += cnt
			adding[b] += cnt
			adding[num] -= cnt
		} else {
			adding[num] -= cnt
			adding[num*2024] += cnt
		}
	}

	for num, amount := range adding {
		freq[num] += amount
	}
}

func partOne(input []int) {
	freq := make(map[int]int)

	for _, num := range input {
		freq[num]++
	}

	for range 25 {
		blink(freq)
	}

	res := 0
	for _, cnt := range freq {
		res += cnt
	}

	fmt.Println("Part 1:", res)
}

func partTwo(input []int) {
	freq := make(map[int]int)

	for _, num := range input {
		freq[num]++
	}

	for range 75 {
		blink(freq)
	}

	res := 0
	for _, cnt := range freq {
		res += cnt
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.Fields(scanner.Text())
	input := make([]int, len(line))
	for i, s := range line {
		input[i] = stoi(s)
	}

	partOne(input)
	partTwo(input)
}

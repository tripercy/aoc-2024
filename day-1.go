package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func partOne(leftArr, rightArr []int) {
	slices.Sort(leftArr)
	slices.Sort(rightArr)
	res := 0

	for i := range leftArr {
		res += int(math.Abs(float64(leftArr[i]) - float64(rightArr[i])))
	}

	fmt.Println("Part 1: ", res)
}

func partTwo(leftArr, rightArr []int) {
	freq := make(map[int]int)

	for _, num := range rightArr {
		freq[num]++
	}

	res := 0
	for _, num := range leftArr {
		res += num * freq[num]
	}
	fmt.Println("Part 2: ", res)
}

func main() {
	file, err := os.Open("input/day-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	leftArr := make([]int, len(input))
	rightArr := make([]int, len(input))

	for i, line := range input {
		nums := strings.Fields(line)
		if len(nums) != 2 {
			log.Fatal("Split failed with input " + line)
		}

		if a, err := strconv.Atoi(nums[0]); err != nil {
			log.Fatal(err)
		} else {
			leftArr[i] = a
		}

		if b, err := strconv.Atoi(nums[1]); err != nil {
			log.Fatal(err)
		} else {
			rightArr[i] = b
		}
	}

	partOne(leftArr, rightArr)
	partTwo(leftArr, rightArr)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
	} else {
		return i
	}
	return -1
}

func parseLine(line string) (int, []int) {
	tokens := strings.Split(line, ": ")
	le := atoi(tokens[0])
	re := strings.Fields(tokens[1])
	ls := make([]int, len(re))
	for i, s := range re {
		ls[i] = atoi(s)
	}

	return le, ls
}

func canCalibrate(res int, terms []int) bool {
	if len(terms) == 1 {
		return terms[0] == res
	}
	n := len(terms)
	flag := canCalibrate(res-terms[n-1], terms[:n-1])
	if !flag && res%terms[n-1] == 0 {
		flag = canCalibrate(res/terms[n-1], terms[:n-1])
	}
	return flag
}

func partOne(lhs []int, rhs [][]int) {
	res := 0

	for i := range lhs {
		if canCalibrate(lhs[i], rhs[i]) {
			res += lhs[i]
		}
	}

	fmt.Println("Part 1:", res)
}

func isSuffix(x, target int) bool {
	x -= target
	for target > 0 {
		if x%10 != 0 {
			return false
		}
		x /= 10
		target /= 10
	}
	return true
}

func countDigits(n int) int {
	res := 0
	for n > 0 {
		n /= 10
		res++
	}
	return res
}

func trimInt(n, length int) int {
	for ; length > 0; length-- {
		n /= 10
	}
	return n
}

func canCalibrateConcat(res int, terms []int) bool {
	if len(terms) == 1 {
		return terms[0] == res
	}

	n := len(terms)
	flag := canCalibrateConcat(res-terms[n-1], terms[:n-1])

	if !flag && res%terms[n-1] == 0 {
		flag = canCalibrateConcat(res/terms[n-1], terms[:n-1])
	}

	if !flag && isSuffix(res, terms[n-1]) {
		flag = canCalibrateConcat(trimInt(res, countDigits(terms[n-1])), terms[:n-1])
	}
	return flag
}

func partTwo(lhs []int, rhs [][]int) {
	var res uint64 = 0

	for i := range lhs {
		if canCalibrateConcat(lhs[i], rhs[i]) {
			res += uint64(lhs[i])
		}
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lhs := make([]int, 0)
	rhs := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		le, re := parseLine(line)
		lhs = append(lhs, le)
		rhs = append(rhs, re)
	}

	partOne(lhs, rhs)
	partTwo(lhs, rhs)
}

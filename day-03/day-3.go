package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func partOne(input []string) {
	res := 0
	re, err := regexp.Compile("mul\\(([0-9]+),([0-9]+)\\)")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range input {
		matches := re.FindAllSubmatch([]byte(line), -1)
		for _, match := range matches {
			a, err := strconv.Atoi(string(match[1]))
			if err != nil {
				log.Fatal(err)
			}
			b, err := strconv.Atoi(string(match[2]))
			if err != nil {
				log.Fatal(err)
			}
			res += a * b
		}
	}

	fmt.Println("Part 1:", res)
}

func partTwo(input []string) {
	res := 0
	reMul, errMul := regexp.Compile("mul\\(([0-9]+),([0-9]+)\\)")

	if errMul != nil {
		log.Fatal(errMul)
	}
	reDo, errDo := regexp.Compile("do\\(\\)")
	if errDo != nil {
		log.Fatal(errDo)
	}
	reDont, errDont := regexp.Compile("don't\\(\\)")
	if errDont != nil {
		log.Fatal(errDont)
	}

	last := 0
	tokens := make([][]int, 0)
	for _, line := range input {
		mulMatches := reMul.FindAllSubmatch([]byte(line), -1)
		mulIndexes := reMul.FindAllSubmatchIndex([]byte(line), -1)
		for i, match := range mulMatches {
			a, err := strconv.Atoi(string(match[1]))
			if err != nil {
				log.Fatal(err)
			}
			b, err := strconv.Atoi(string(match[2]))
			if err != nil {
				log.Fatal(err)
			}
			index := mulIndexes[i][0]
			tokens = append(tokens, []int{last + index, 2, a * b})
		}

		doIndexes := reDo.FindAllSubmatchIndex([]byte(line), -1)
		dontIndexes := reDont.FindAllSubmatchIndex([]byte(line), -1)
		for _, index := range doIndexes {
			tokens = append(tokens, []int{last + index[0], 1})
		}
		for _, index := range dontIndexes {
			tokens = append(tokens, []int{last + index[0], 0})
		}
		last += len(line)
	}

	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i][0] < tokens[j][0]
	})

	doing := true
	for _, token := range tokens {
		if token[1] == 1 {
			doing = true
		} else if token[1] == 0 {
			doing = false
		} else if doing {
			res += token[2]
		}
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-3.txt")
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

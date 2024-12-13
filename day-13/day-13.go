package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Pair struct {
	x int
	y int
}

type Arcade struct {
	btnA  Pair
	btnB  Pair
	prize Pair
}

func stoi(s string) int {
	if res, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
	} else {
		return res
	}

	return -1
}

var re, _ = regexp.Compile("([0-9]+)")

func getIntPair(s string) Pair {
	matches := re.FindAllSubmatch([]byte(s), -1)
	res := Pair{
		x: stoi(string(matches[0][1])),
		y: stoi(string(matches[1][1])),
	}

	return res
}

func createArcade(btnA, btnB, prize string) Arcade {
	return Arcade{
		btnA:  getIntPair(btnA),
		btnB:  getIntPair(btnB),
		prize: getIntPair(prize),
	}
}

func findMinToken(arcade Arcade, broken bool) int {
	xA, yA := arcade.btnA.x, arcade.btnA.y
	xB, yB := arcade.btnB.x, arcade.btnB.y
	xP, yP := arcade.prize.x, arcade.prize.y
	if broken {
		xP += 10000000000000
		yP += 10000000000000
	}

	bNom := xP*yA - yP*xA
	bDenom := xB*yA - yB*xA

	if bNom%bDenom != 0 {
		return 0
	}
	b := bNom / bDenom

	aNom := yP - yB*b
	aDenom := yA

	if aNom%aDenom != 0 {
		return 0
	}
	a := aNom / aDenom

	return a*3 + b*1
}

func partOne(arcades []Arcade) {
	res := 0

	for _, arcade := range arcades {
		res += findMinToken(arcade, false)
	}

	fmt.Println("Part 1:", res)
}

func partTwo(arcades []Arcade) {
	res := 0

	for _, arcade := range arcades {
		res += findMinToken(arcade, true)
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	arcades := make([]Arcade, 0)
	for scanner.Scan() {
		btnA := scanner.Text()
		scanner.Scan()
		btnB := scanner.Text()
		scanner.Scan()
		prize := scanner.Text()

		arcades = append(arcades, createArcade(btnA, btnB, prize))

		// Newline
		scanner.Scan()
	}

	partOne(arcades)
	partTwo(arcades)
}

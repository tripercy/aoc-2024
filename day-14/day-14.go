package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const nrows = 103
const ncols = 101

type Pair struct {
	x int
	y int
}

type Robot struct {
	pos Pair
	vel Pair
}

func stoi(s string) int {
	if res, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
	} else {
		return res
	}

	return -1
}

var re, _ = regexp.Compile("p=(.*),(.*) v=(.*),(.*)")

func parseRobot(robotStr string) Robot {
	match := re.FindSubmatch([]byte(robotStr))
	return Robot{
		pos: Pair{
			x: stoi(string(match[1])),
			y: stoi(string(match[2])),
		},
		vel: Pair{
			x: stoi(string(match[3])),
			y: stoi(string(match[4])),
		},
	}
}

func getQuadrant(pos Pair) int {
	if pos.x == (ncols-1)/2 || pos.y == (nrows-1)/2 {
		return -1
	}
	if pos.y < (nrows-1)/2 {
		if pos.x > (ncols-1)/2 {
			return 0
		} else {
			return 1
		}
	} else {
		if pos.x > (ncols-1)/2 {
			return 3
		} else {
			return 2
		}
	}
}

func (r Robot) jump(steps int) Pair {
	finalPos := r.pos
	finalPos.x += r.vel.x * steps
	finalPos.y += r.vel.y * steps
	finalPos.x = (finalPos.x%ncols + ncols) % ncols
	finalPos.y = (finalPos.y%nrows + nrows) % nrows

	return finalPos
}

func partOne(robots []Robot) {
	var quadrants [4]int

	for _, robot := range robots {
		finalPos := robot.jump(100)
		quad := getQuadrant(finalPos)
		if quad > -1 {
			quadrants[quad]++
		}
	}

	res := 1
	for _, quad := range quadrants {
		res *= quad
	}

	fmt.Println("Part 1:", res)
}

func match(mp, pattern [][]bool) bool {
	for r := 0; r < len(mp)-len(pattern); r++ {
		for c := 0; c < len(mp[0])-len(pattern[0]); c++ {
			flag := true
			for r0, prow := range pattern {
				for c0 := range prow {
					if mp[r+r0][c+c0] != pattern[r0][c0] {
						flag = false
						break
					}
				}
				if !flag {
					break
				}
			}
			if flag {
				return true
			}
		}
	}

	return false
}

func dumpMap(mp [][]bool) {
	for _, row := range mp {
		for _, col := range row {
			if col {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func partTwo(robots []Robot) {
	pattern := [][]bool{
		{false, false, true, false, false},
		{false, true, true, true, false},
		{true, true, true, true, true},
	}

	for i := range int(1e4) {
		mp := make([][]bool, nrows)
		for j := range mp {
			mp[j] = make([]bool, ncols)
		}
		for _, robot := range robots {
			finalPos := robot.jump(i)
			mp[finalPos.y][finalPos.x] = true
		}
		if match(mp, pattern) {
			fmt.Println("Iteration:", i)
			dumpMap(mp)
			fmt.Println("====")
		}
	}
}

func main() {
	file, err := os.Open("day-14.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	robots := make([]Robot, 0)
	for scanner.Scan() {
		line := scanner.Text()
		robots = append(robots, parseRobot(line))
	}

	partOne(robots)
	partTwo(robots)
}

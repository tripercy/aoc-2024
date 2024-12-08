package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func getAttenas(input [][]byte) map[byte][][]int {
	attenas := make(map[byte][][]int)

	for r, line := range input {
		for c, attena := range line {
			if attena != '.' {
				attenas[attena] = append(attenas[attena], []int{r, c})
			}
		}
	}

	return attenas
}

func isInBound(r, c, nr, nc int) bool {
	return r >= 0 && c >= 0 && r < nr && c < nc
}

func partOne(input [][]byte) {
	nr := len(input)
	nc := len(input[0])
	attenas := getAttenas(input)

	affected := make([][]bool, nr)
	for i := range affected {
		affected[i] = make([]bool, nc)
	}

	for _, positions := range attenas {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				posA := positions[i]
				posB := positions[j]
				rA, cA := posA[0], posA[1]
				rB, cB := posB[0], posB[1]

				rDis := abs(rA - rB)
				cDis := abs(cA - cB)

				r1 := rA - rDis
				r2 := rB + rDis
				c1 := cA - cDis
				c2 := cB + cDis
				if cA > cB {
					c1 = cA + cDis
					c2 = cB - cDis
				}
				// fmt.Println(rA, cA, rB, cB, r1, c1, r2, c2)
				if isInBound(r1, c1, nr, nc) {
					affected[r1][c1] = true
				}
				if isInBound(r2, c2, nr, nc) {
					affected[r2][c2] = true
				}
			}
		}
	}

	res := 0
	for _, row := range affected {
		for _, pos := range row {
			if pos {
				res++
			}
		}
	}

	fmt.Println("Part 1:", res)
}

func partTwo(input [][]byte) {
	nr := len(input)
	nc := len(input[0])
	attenas := getAttenas(input)

	affected := make([][]bool, nr)
	for i := range affected {
		affected[i] = make([]bool, nc)
	}

	for _, positions := range attenas {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				posA := positions[i]
				posB := positions[j]
				rA, cA := posA[0], posA[1]
				rB, cB := posB[0], posB[1]
				rDis := abs(rA - rB)
				cDis := abs(cA - cB)

				for isInBound(rA, cA, nr, nc) {
					affected[rA][cA] = true
					rA -= rDis
					if cA < cB {
						cA -= cDis
					} else {
						cA += cDis
					}
				}
				for isInBound(rB, cB, nr, nc) {
					affected[rB][cB] = true
					rB += rDis
					if cB < cA {
						cB -= cDis
					} else {
						cB += cDis
					}
				}
			}
		}
	}

	res := 0
	for _, row := range affected {
		for _, pos := range row {
			if pos {
				res++
			}
		}
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []byte(line))
	}

	partOne(input)
	partTwo(input)
}

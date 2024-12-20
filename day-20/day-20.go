package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Position struct {
	r int
	c int
}

var dr = []int{0, 0, 1, -1}
var dc = []int{1, -1, 0, 0}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func manhattanDist(a, b Position) int {
	r := abs(a.r - b.r)
	c := abs(a.c - b.c)

	return r + c
}

func djikstra(mp [][]byte, r, c int) [][]int {
	nr, nc := len(mp), len(mp[0])
	dist := make([][]int, nr)
	for i := range nr {
		dist[i] = make([]int, nc)
		for j := range nc {
			dist[i][j] = int(1e9)
		}
	}
	dist[r][c] = 0

	posQueue := make([]Position, 0)
	posQueue = append(posQueue, Position{r, c})

	for i := 0; i < len(posQueue); i++ {
		pos := posQueue[i]
		cr, cc := pos.r, pos.c
		cdist := dist[cr][cc]
		for d := range 4 {
			tr, tc := cr+dr[d], cc+dc[d]
			if mp[tr][tc] == '#' {
				continue
			}
			ndist := cdist + 1
			if dist[tr][tc] <= ndist {
				continue
			}
			dist[tr][tc] = ndist
			posQueue = append(posQueue, Position{tr, tc})
		}
	}

	return dist
}

func getPos(mp [][]byte, x byte) Position {
	for r, row := range mp {
		for c, col := range row {
			if col == x {
				return Position{r, c}
			}
		}
	}
	return Position{-1, -1}
}

func countCheats(mp [][]byte, cheatLim, saveLeast int) int {
	start := getPos(mp, 'S')
	end := getPos(mp, 'E')

	startDist := djikstra(mp, start.r, start.c)
	endDist := djikstra(mp, end.r, end.c)

	maxDist := startDist[end.r][end.c]
	res := 0
	for r, row := range mp {
		for c, col := range row {
			if col == '#' {
				continue
			}
			for r2 := r - cheatLim; r2 <= r+cheatLim; r2++ {
				if r2 < 0 {
					continue
				}
				if r2 >= len(mp) {
					break
				}
				for c2 := c - cheatLim; c2 <= c+cheatLim; c2++ {
					if c2 < 0 {
						continue
					}
					if c2 >= len(row) {
						break
					}
					if mp[r2][c2] == '#' {
						continue
					}
					dist := manhattanDist(Position{r, c}, Position{r2, c2})
					if dist > cheatLim {
						continue
					}
					if startDist[r][c]+endDist[r2][c2]+dist-1 < maxDist-saveLeast {
						res++
					}
				}
			}
		}
	}

	return res
}

func partOne(mp [][]byte) {
	fmt.Println("Part 1:", countCheats(mp, 2, 100))
}

func partTwo(mp [][]byte) {
	fmt.Println("Part 2:", countCheats(mp, 20, 100))
}

func main() {
	file, err := os.Open("day-20.txt")
	// file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mp := make([][]byte, 0)
	for scanner.Scan() {
		mp = append(mp, []byte(scanner.Text()))
	}

	partOne(mp)
	partTwo(mp)
}

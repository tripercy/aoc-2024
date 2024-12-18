package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	nrows = 71
	ncols = 71
)

type Byte struct {
	x int
	y int
}

func stoi(s string) int {
	if res, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
	} else {
		return res
	}
	return -1
}

func createByte(s string) Byte {
	tokens := strings.Split(s, ",")
	return Byte{
		x: stoi(tokens[0]),
		y: stoi(tokens[1]),
	}
}

type queueItem struct {
	row int
	col int
	len int
}

func findPath(grid [nrows][ncols]bool) int {
	dr := []int{0, 0, 1, -1}
	dc := []int{1, -1, 0, 0}

	queue := make([]queueItem, 0)
	queue = append(queue, queueItem{0, 0, 0})

	res := 0
	for i := 0; i < len(queue) && res == 0; i++ {
		item := queue[i]
		r, c := item.row, item.col
		ln := item.len
		for d := range 4 {
			nr, nc := r+dr[d], c+dc[d]
			if nr < 0 || nc < 0 || nr >= nrows || nc >= ncols || grid[nr][nc] {
				continue
			}
			if nr == nrows-1 && nc == ncols-1 {
				res = ln + 1
				return res
			}
			grid[nr][nc] = true
			queue = append(queue, queueItem{nr, nc, ln + 1})
		}
	}

	return -1
}

func partOne(bytes []Byte) {
	var grid [nrows][ncols]bool
	for _, b := range bytes[:1024] {
		grid[b.y][b.x] = true
	}
	grid[0][0] = true

	res := findPath(grid)

	fmt.Println("Part 1:", res)
}

func partTwo(bytes []Byte) {
	var grid [nrows][ncols]bool
	grid[0][0] = true

	var res Byte
	for _, b := range bytes {
		grid[b.y][b.x] = true
		var tempGrid [nrows][ncols]bool
		copy(tempGrid[:][:], grid[:][:])
		if findPath(tempGrid) == -1 {
			res = b
			break
		}
	}
	fmt.Printf("Part 2: %d,%d\n", res.x, res.y)
}

func main() {
	file, err := os.Open("day-18.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes := make([]Byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bytes = append(bytes, createByte(line))
	}

	partOne(bytes)
	partTwo(bytes)
}

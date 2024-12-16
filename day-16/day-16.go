package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var dr = []int{0, -1, 0, 1}
var dc = []int{1, 0, -1, 0}

type CellState struct {
	row int
	col int
	dir int
}

func addCell(dist map[CellState]int, par map[CellState][]CellState, cellQueue *[]CellState, source, cell CellState, cost int) {
	if oldDist, exist := dist[cell]; !exist || oldDist > cost {
		dist[cell] = cost
		*cellQueue = append(*cellQueue, cell)
		par[cell] = []CellState{source}
	} else if oldDist == cost {
		par[cell] = append(par[cell], source)
	}
}

func solve(maze [][]byte) {
	startCell := CellState{
		row: len(maze) - 2,
		col: 1,
		dir: 0,
	}

	dist := make(map[CellState]int)
	par := make(map[CellState][]CellState)
	dist[startCell] = 0
	cellQueue := []CellState{startCell}

	for i := 0; i < len(cellQueue); i++ {
		cell := cellQueue[i]
		cost := dist[cell]
		r, c := cell.row, cell.col
		dir := cell.dir
		nr, nc := r+dr[dir], c+dc[dir]

		if maze[nr][nc] != '#' {
			addCell(dist, par, &cellQueue, cell, CellState{row: nr, col: nc, dir: dir}, cost+1)
		}
		addCell(dist, par, &cellQueue, cell, CellState{row: r, col: c, dir: (dir + 1) % 4}, cost+1000)
		addCell(dist, par, &cellQueue, cell, CellState{row: r, col: c, dir: (dir - 1 + 4) % 4}, cost+1000)
	}

	// Set up for part 2
	seats := make([][]bool, len(maze))
	for i := range seats {
		seats[i] = make([]bool, len(maze[0]))
	}
	visited := make(map[CellState]bool)
	cellQueue = cellQueue[:0]

	res := int(1e12)
	for d := range 4 {
		endCell := CellState{row: 1, col: len(maze[0]) - 2, dir: d}
		if dist[endCell] < res {
			res = dist[endCell]
			visited[endCell] = true
			cellQueue = append(cellQueue, endCell)
		}
	}

	fmt.Println("Part 1:", res)

	for i := 0; i < len(cellQueue); i++ {
		cell := cellQueue[i]

		r, c := cell.row, cell.col
		seats[r][c] = true

		for _, p := range par[cell] {
			if !visited[p] {
				cellQueue = append(cellQueue, p)
				visited[p] = true
			}
		}
	}

	res = 0
	for _, row := range seats {
		for _, seat := range row {
			if seat {
				res++
			}
		}
	}
	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-16.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maze := make([][]byte, 0)
	for scanner.Scan() {
		maze = append(maze, []byte(scanner.Text()))
	}

	solve(maze)
}

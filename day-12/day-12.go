package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var dr []int = []int{1, -1, 0, 0}
var dc []int = []int{0, 0, 1, -1}

func checkRegion(input []string, visited [][]bool, r, c int) (int, int) {
	visited[r][c] = true
	plant := input[r][c]
	peri, area := 0, 1
	for d := range 4 {
		nr, nc := r+dr[d], c+dc[d]
		if nr < 0 || nc < 0 || nr >= len(input) || nc >= len(input[0]) {
			peri++
			continue
		}
		nPlant := input[nr][nc]
		if plant != nPlant {
			peri++
			continue
		}
		if plant == nPlant && !visited[nr][nc] {
			nPeri, nArea := checkRegion(input, visited, nr, nc)
			peri += nPeri
			area += nArea
			continue
		}
	}

	return peri, area
}

func checkRegionDiscount(input []string, visited, shape [][]bool, r, c int) int {
	visited[r][c] = true
	shape[r+1][c+1] = true
	plant := input[r][c]
	area := 1
	for d := range 4 {
		nr, nc := r+dr[d], c+dc[d]
		if nr < 0 || nc < 0 || nr >= len(input) || nc >= len(input[0]) {
			continue
		}
		nPlant := input[nr][nc]
		if plant != nPlant {
			continue
		}
		if plant == nPlant && !visited[nr][nc] {
			nArea := checkRegionDiscount(input, visited, shape, nr, nc)
			area += nArea
			continue
		}
	}

	return area
}

func partOne(input []string) {
	nrows, ncols := len(input), len(input[0])
	visited := make([][]bool, nrows)
	for i := range nrows {
		visited[i] = make([]bool, ncols)
	}

	res := 0
	for r := range nrows {
		for c := range ncols {
			if !visited[r][c] {
				peri, area := checkRegion(input, visited, r, c)
				res += peri * area
			}
		}
	}
	fmt.Println("Part 1:", res)
}

func getNumEdges(shape [][]bool) int {
	nrows, ncols := len(shape), len(shape[0])
	res := 0
	for horLine := 1; horLine < nrows; horLine++ {
		for col := 1; col < ncols-1; col++ {
			if shape[horLine-1][col] != shape[horLine][col] {
				res++
				for shape[horLine-1][col] != shape[horLine][col] && shape[horLine][col+1] == shape[horLine][col] {
					col++
				}
			}
		}
	}
	for verLine := 1; verLine < ncols; verLine++ {
		for row := 1; row < nrows-1; row++ {
			if shape[row][verLine-1] != shape[row][verLine] {
				res++
				for shape[row][verLine-1] != shape[row][verLine] && shape[row+1][verLine] == shape[row][verLine] {
					row++
				}
			}
		}
	}

	return res
}

func partTwo(input []string) {
	nrows, ncols := len(input), len(input[0])
	visited := make([][]bool, nrows)
	for i := range nrows {
		visited[i] = make([]bool, ncols)
	}

	res := 0
	for r := range nrows {
		for c := range ncols {
			if !visited[r][c] {
				shape := make([][]bool, nrows+2)
				for i := range nrows + 2 {
					shape[i] = make([]bool, ncols+2)
				}
				area := checkRegionDiscount(input, visited, shape, r, c)
				edges := getNumEdges(shape)
				// fmt.Println(string(input[r][c]), edges)
				res += area * edges
			}
		}
	}
	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	partOne(input)
	partTwo(input)
}

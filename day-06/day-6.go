package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInitPos(input []string) (int, int) {
	r, c := -1, -1
	for i, line := range input {
		for j, ch := range line {
			if ch == '^' {
				r, c = i, j
				break
			}
		}
		if r != -1 {
			break
		}
	}

	return r, c
}

func toByte2D(input []string) [][]byte {
	arr := make([][]byte, len(input))
	for i, line := range input {
		arr[i] = []byte(line)
	}

	return arr
}

func getMp(input []string) [][]byte {
	r, c := getInitPos(input)
	mp := toByte2D(input)

	dr, dc := -1, 0
	for {
		mp[r][c] = 'X'
		nr, nc := r+dr, c+dc
		if nr < 0 || nc < 0 || nr >= len(mp) || nc >= len(mp[0]) {
			break
		}
		for mp[nr][nc] == '#' {
			dr, dc = dc, dr
			if dr == 0 {
				dc = -dc
			}
			nr, nc = r+dr, c+dc
		}
		r, c = nr, nc
	}

	return mp
}

func dirId(dr, dc int) int {
	if dr == -1 {
		// UP
		return 0
	}
	if dr == 1 {
		// DOWN
		return 2
	}
	if dc == 1 {
		// RIGHT
		return 1
	}
	// LEFT
	return 3
}

func mark(marked [][]byte, r, c, dir int) {
	marked[r][c] |= 1 << dir
}

func isMarked(marked [][]byte, r, c, dir int) bool {
	return ((marked[r][c] >> dir) & 1) == 1
}

func hasLoop(mp [][]byte, r, c int) bool {
	marked := make([][]byte, len(mp))
	for i := range marked {
		marked[i] = make([]byte, len(mp[i]))
	}

	dr, dc := -1, 0

	for {
		if isMarked(marked, r, c, dirId(dr, dc)) {
			return true
		}
		mark(marked, r, c, dirId(dr, dc))
		nr, nc := r+dr, c+dc
		if nr < 0 || nc < 0 || nr >= len(mp) || nc >= len(mp[0]) {
			return false
		}
		for mp[nr][nc] == '#' {
			dr, dc = dc, dr
			if dr == 0 {
				dc = -dc
			}
			nr, nc = r+dr, c+dc
		}
		r, c = nr, nc
	}
}

func partOne(input []string) {
	res := 0
	mp := getMp(input)

	for _, row := range mp {
		for _, ch := range row {
			if ch == 'X' {
				res++
			}
		}
	}

	fmt.Println("Part 1:", res)
}

func partTwo(input []string) {
	res := 0

	mp := getMp(input)
	modified := toByte2D(input)
	r0, c0 := getInitPos(input)

	for r, row := range mp {
		for c, ch := range row {
			if (r != r0 || c != c0) && ch == 'X' {
				modified[r][c] = '#'
				if hasLoop(modified, r0, c0) {
					res++
				}
				modified[r][c] = '.'
			}
		}
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-6.txt")
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

	// tmp := []string{
	// 	"....#.....",
	// 	"....+#--+#",
	// 	"....|...|.",
	// 	"..#.|...|.",
	// 	"....|..#|.",
	// 	"....|...|.",
	// 	".#.#^---+.",
	// 	"........#.",
	// 	"#.........",
	// 	"......#...",
	// }
	// mp := toByte2D(tmp)
	// fmt.Println(hasLoop(mp, 6, 4))
}

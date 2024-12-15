package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func copyMap(mp [][]byte) [][]byte {
	newMap := make([][]byte, len(mp))
	for i, row := range mp {
		newMap[i] = make([]byte, len(row))
		copy(newMap[i], row)
	}
	return newMap
}

func getRobotPos(mp [][]byte) (int, int) {
	for i, row := range mp {
		for j, col := range row {
			if col == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func opToDir(op rune) int {
	switch op {
	case '>':
		return 0
	case '<':
		return 1
	case 'v':
		return 2
	case '^':
		return 3
	}
	return -1
}

func partOne(mp [][]byte, ops string) {
	mp = copyMap(mp)
	r, c := getRobotPos(mp)
	if r == -1 {
		log.Fatal("Unable to locate robot")
	}

	dr := []int{0, 0, 1, -1}
	dc := []int{1, -1, 0, 0}

	for _, op := range ops {
		d := opToDir(op)
		nr, nc := r+dr[d], c+dc[d]
		for mp[nr][nc] == 'O' {
			nr += dr[d]
			nc += dc[d]
		}
		if mp[nr][nc] == '#' {
			continue
		}
		if nr == r+dr[d] && nc == c+dc[d] {
			mp[nr][nc] = '@'
			mp[r][c] = '.'
			r, c = nr, nc
		} else {
			mp[nr][nc] = 'O'
			mp[r][c] = '.'
			r += dr[d]
			c += dc[d]
			mp[r][c] = '@'
		}
	}

	res := 0
	for r, row := range mp {
		for c, col := range row {
			if col == 'O' {
				res += r*100 + c
			}
		}
	}

	fmt.Println("Part 1:", res)
}

func expandMap(mp [][]byte) [][]byte {
	newMp := make([][]byte, len(mp))
	for r, row := range mp {
		newMp[r] = make([]byte, 2*len(row))
		for c, col := range row {
			var newCells []byte
			switch col {
			case '#':
				newCells = []byte("##")
			case '.':
				newCells = []byte("..")
			case 'O':
				newCells = []byte("[]")
			case '@':
				newCells = []byte("@.")
			}
			newMp[r][c*2] = newCells[0]
			newMp[r][c*2+1] = newCells[1]
		}
	}

	return newMp
}

func getHalfBoxPos(r, c int, half byte) (int, int) {
	if half == '[' {
		return r, c + 1
	}
	return r, c - 1
}

func tryPush(mp [][]byte, r, c, dr, dc int) bool {
	visited := make([][]bool, len(mp))
	for i := range visited {
		visited[i] = make([]bool, len(mp[0]))
	}

	r2, c2 := getHalfBoxPos(r, c, mp[r][c])
	visited[r][c] = true
	visited[r2][c2] = true
	boxes := [][]int{{r, c, r2, c2}}

	for i := 0; i < len(boxes); i++ {
		tr1, tc1 := boxes[i][0], boxes[i][1]
		tr2, tc2 := boxes[i][2], boxes[i][3]

		nr1, nc1 := tr1+dr, tc1+dc
		if mp[nr1][nc1] == '#' {
			return false
		}

		if mp[nr1][nc1] != '.' && !visited[nr1][nc1] {
			nr2, nc2 := getHalfBoxPos(nr1, nc1, mp[nr1][nc1])
			visited[nr1][nc1] = true
			visited[nr2][nc2] = true
			boxes = append(boxes, []int{nr1, nc1, nr2, nc2})
		}

		nr1, nc1 = tr2+dr, tc2+dc
		if mp[nr1][nc1] == '#' {
			return false
		}

		if mp[nr1][nc1] != '.' && !visited[nr1][nc1] {
			nr2, nc2 := getHalfBoxPos(nr1, nc1, mp[nr1][nc1])
			visited[nr1][nc1] = true
			visited[nr2][nc2] = true
			boxes = append(boxes, []int{nr1, nc1, nr2, nc2})
		}
	}

	sort.Slice(boxes, func(i, j int) bool {
		if dr == 0 {
			if dc > 0 {
				return boxes[i][1] > boxes[j][1]
			}
			return boxes[i][1] < boxes[j][1]
		}
		if dr > 0 {
			return boxes[i][0] > boxes[j][0]
		}
		return boxes[i][0] < boxes[j][0]
	})

	for _, box := range boxes {
		r1, c1, r2, c2 := box[0], box[1], box[2], box[3]
		nr1, nc1 := r1+dr, c1+dc
		nr2, nc2 := r2+dr, c2+dc
		sym1, sym2 := mp[r1][c1], mp[r2][c2]
		mp[r1][c1] = '.'
		mp[r2][c2] = '.'
		mp[nr1][nc1] = sym1
		mp[nr2][nc2] = sym2
	}

	return true
}

func partTwo(mp [][]byte, ops string) {
	mp = expandMap(mp)
	r, c := getRobotPos(mp)

	dr := []int{0, 0, 1, -1}
	dc := []int{1, -1, 0, 0}

	for _, op := range ops {
		d := opToDir(op)
		nr, nc := r+dr[d], c+dc[d]
		if mp[nr][nc] == '#' {
			continue
		}
		if mp[nr][nc] == '.' || tryPush(mp, nr, nc, dr[d], dc[d]) {
			mp[r][c] = '.'
			mp[nr][nc] = '@'
			r, c = nr, nc
		}

	}

	res := 0
	for r, row := range mp {
		for c, col := range row {
			if col == '[' {
				res += r*100 + c
			}
		}
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-15.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mp := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for line := scanner.Text(); len(line) > 1; _, line = scanner.Scan(), scanner.Text() {
		mp = append(mp, []byte(line))
	}
	ops := ""
	for scanner.Scan() {
		ops += scanner.Text()
	}
	partOne(mp, ops)
	partTwo(mp, ops)
}

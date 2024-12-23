package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func stoi(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

type Position struct {
	row int
	col int
}

type DpadMove struct {
	to   byte
	from byte
}

var dpad = []string{
	" ^A",
	"<v>",
}

var npad = []string{
	"789",
	"456",
	"123",
	" 0A",
}

var buttonPos = make(map[byte]Position)
var numPos = make(map[byte]Position)
var padTrans = make(map[DpadMove][]string)

func movePad(from, to byte, pad []string, pos map[byte]Position) []string {
	start := pos[from]
	end := pos[to]

	hor, ver := "", ""
	for range abs(end.row - start.row) {
		if start.row < end.row {
			ver += "v"
		} else {
			ver += "^"
		}
	}

	for range abs(end.col - start.col) {
		if start.col < end.col {
			hor += ">"
		} else {
			hor += "<"
		}
	}

	if len(hor) == 0 || len(ver) == 0 {
		return []string{hor + ver + "A"}
	}
	res := make([]string, 0)

	if pad[start.row][end.col] != ' ' {
		res = append(res, hor+ver+"A")
	}
	if pad[end.row][start.col] != ' ' {
		res = append(res, ver+hor+"A")
	}

	return res
}

func solve(input []string) {
	dp := make([][][]uint64, 26)
	for i := range 26 {
		dp[i] = make([][]uint64, 500)
		for j := range 500 {
			dp[i][j] = make([]uint64, 500)
			for k := range 500 {
				dp[i][j][k] = uint64(5e15)
			}
		}
	}

	for _, c1 := range "A^v><" {
		for _, c2 := range "A^v><" {
			dp[0][c1][c2] = 1
		}
	}

	for t := 1; t < 26; t++ {
		for _, c1 := range []byte("A^v><") {
			for _, c2 := range []byte("A^v><") {
				combs := padTrans[DpadMove{c1, c2}]
				for _, seq := range combs {
					l := uint64(0)
					seq = "A" + seq
					for i := 1; i < len(seq); i++ {
						l += dp[t-1][byte(seq[i-1])][byte(seq[i])]
					}
					dp[t][c1][c2] = min(dp[t][c1][c2], l)
				}
			}
		}
	}

	res1 := uint64(0)
	res2 := uint64(0)

	for _, password := range input {
		password = "A" + password
		l1 := uint64(0)
		l2 := uint64(0)

		for i := 1; i < len(password); i++ {
			c1, c2 := byte(password[i-1]), byte(password[i])
			combs := padTrans[DpadMove{c1, c2}]
			l_task1 := uint64(5e15)
			l_task2 := uint64(5e15)
			for _, seq := range combs {
				tmp1 := uint64(0)
				tmp2 := uint64(0)
				seq = "A" + seq
				for i := 1; i < len(seq); i++ {
					tmp1 += dp[2][byte(seq[i-1])][byte(seq[i])]
					tmp2 += dp[25][byte(seq[i-1])][byte(seq[i])]
				}
				l_task1 = min(l_task1, tmp1)
				l_task2 = min(l_task2, tmp2)
			}
			l1 += l_task1
			l2 += l_task2
			// fmt.Println(string(c1), string(c2), combs, l_task1)
		}
		// fmt.Println(password[1:], l)
		res1 += l1 * uint64(stoi(password[1:len(password)-1]))
		res2 += l2 * uint64(stoi(password[1:len(password)-1]))
	}

	fmt.Println("Part 1:", res1)
	fmt.Println("Part 2:", res2)
}

func main() {
	file, err := os.Open("day-21.txt")
	// file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	for r, row := range dpad {
		for c, ch := range []byte(row) {
			buttonPos[ch] = Position{r, c}
		}
	}

	for r, row := range npad {
		for c, ch := range []byte(row) {
			numPos[ch] = Position{r, c}
		}
	}

	for _, c1 := range []byte("A^v><") {
		for _, c2 := range []byte("A^v><") {
			padTrans[DpadMove{c1, c2}] = movePad(c1, c2, dpad, buttonPos)
		}
	}

	for _, c1 := range []byte("0123456789A") {
		for _, c2 := range []byte("0123456789A") {
			padTrans[DpadMove{c1, c2}] = movePad(c1, c2, npad, numPos)
		}
	}

	solve(input)
}

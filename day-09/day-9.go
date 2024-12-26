package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func atoi(a rune) int {
	return int(a) - int('0')
}

func partOne(input string) {
	filesys := make([]int, 0)
	id := 0
	for i, c := range input {
		cnt := atoi(c)
		adding := make([]int, cnt)

		var val int
		if i%2 == 0 {
			val = id
			id++
		} else {
			val = -1
		}

		for i := range cnt {
			adding[i] = val
		}
		filesys = append(filesys, adding...)
	}

	l, r := 0, len(filesys)-1
	for l < r {
		for filesys[l] != -1 {
			l++
		}
		for filesys[r] == -1 {
			r--
		}
		for filesys[l] == -1 && filesys[r] != -1 && l < r {
			filesys[l] = filesys[r]
			filesys[r] = -1
			l++
			r--
		}
	}

	var res uint64 = 0

	for i, x := range filesys {
		if x > -1 {
			res += uint64(i * x)
		}
	}

	fmt.Println("Part 1:", res)
}

func partTwo(input string) {
	files := make([][]int, 0)
	gaps := make([][]int, 0)
	id := 0
	last := 0
	for i, c := range input {
		cnt := atoi(c)
		if i%2 == 0 {
			files = append(files, []int{last, cnt})
			id++
		} else {
			gaps = append(gaps, []int{last, cnt})
		}
		last += cnt
	}

	id--
	for ; id > 0; id-- {
		file := files[id]
		i, cnt := file[0], file[1]
		for _, gap := range gaps {
			if gap[0] > i {
				break
			}
			if gap[1] >= cnt {
				file[0] = gap[0]
				gap[0] += file[1]
				gap[1] -= file[1]
				break
			}
		}
	}
	var res uint64 = 0
	for id, file := range files {
		a := file[0]
		b := file[0] + file[1] - 1
		sum := (b*(b+1) - a*(a-1)) / 2
		res += uint64(sum * id)
	}
	fmt.Println("Part 2:", res)
}

func run(sol func(string), input string, wg *sync.WaitGroup) {
	defer wg.Done()
	sol(input)
}

func main() {
	file, err := os.Open("day-9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go run(partOne, input, wg)
	go run(partTwo, input, wg)
	wg.Wait()
}

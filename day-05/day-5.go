package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func atoi(s string) int {
	if x, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
	} else {
		return x
	}

	return -1
}

func checkOrder(edges [][]int, ls []int) bool {
	for i, u := range ls {
		for j := i + 1; j < len(ls); j++ {
			v := ls[j]
			if edges[v][u] == 1 {
				return false
			}
		}
	}

	return true
}

func partOne(edges, lists [][]int) {
	res := 0

	for _, ls := range lists {
		if checkOrder(edges, ls) {
			res += ls[len(ls)/2]
		}
	}

	fmt.Println("Part 1:", res)
}

func partTwo(edges, lists [][]int) {
	res := 0

	for _, ls := range lists {
		if checkOrder(edges, ls) {
			continue
		}

		sort.Slice(ls, func(i, j int) bool {
			u, v := ls[i], ls[j]
			return edges[u][v] == 1
		})

		res += ls[len(ls)/2]
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-5.txt")

	if err != nil {
		log.Fatal("err")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	edges := make([][]int, 100)
	for i := range 100 {
		edges[i] = make([]int, 100)
	}

	for scanner.Scan() && len(scanner.Text()) > 1 {
		line := scanner.Text()
		pages := strings.Split(line, "|")
		u, v := atoi(pages[0]), atoi(pages[1])
		edges[u][v] = 1
	}

	lists := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		strList := strings.Split(line, ",")
		intList := make([]int, len(strList))
		for i, x := range strList {
			intList[i] = atoi(x)
		}
		lists = append(lists, intList)
	}

	partOne(edges, lists)
	partTwo(edges, lists)
}

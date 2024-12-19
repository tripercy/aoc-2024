package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type TrieNode struct {
	isWord   bool
	children map[byte]*TrieNode
}

func addWord(root *TrieNode, word string) {
	curr := root
	for _, c := range []byte(word) {
		node, exists := curr.children[c]
		if !exists {
			node = &TrieNode{isWord: false, children: make(map[byte]*TrieNode)}
			curr.children[c] = node
		}
		curr = node
	}
	curr.isWord = true
}

var valid map[string]bool = make(map[string]bool)

func canMake(root *TrieNode, word string) bool {
	// fmt.Println(word)
	curr := root
	if len(word) == 0 {
		return true
	}
	if res, exists := valid[word]; exists {
		return res
	}

	for i, c := range []byte(word) {
		node, exists := curr.children[c]
		if !exists {
			valid[word] = false
			return false
		}
		curr = node
		if curr.isWord {
			canMakeSubstr := canMake(root, word[i+1:])
			if canMakeSubstr {
				valid[word] = true
				return true
			}
		}
	}
	valid[word] = false
	return false
}

func partOne(patterns, requirements []string) {
	root := &TrieNode{isWord: false, children: make(map[byte]*TrieNode)}
	for _, pattern := range patterns {
		addWord(root, pattern)
	}

	res := 0
	for _, pattern := range requirements {
		if canMake(root, pattern) {
			// fmt.Println(pattern)
			res++
		}
	}

	fmt.Println("Part 1:", res)
}

var ways map[string]int = make(map[string]int)

func countWays(root *TrieNode, word string) int {
	// fmt.Println(word)
	curr := root
	if len(word) == 0 {
		return 1
	}
	if res, exists := ways[word]; exists {
		return res
	}

	for i, c := range []byte(word) {
		node, exists := curr.children[c]
		if !exists {
			break
		}
		curr = node
		if curr.isWord {
			numWays := countWays(root, word[i+1:])
			ways[word] += numWays
		}
	}
	return ways[word]
}

func partTwo(patterns, requirements []string) {
	root := &TrieNode{isWord: false, children: make(map[byte]*TrieNode)}
	for _, pattern := range patterns {
		addWord(root, pattern)
	}

	res := 0
	for _, pattern := range requirements {
		res += countWays(root, pattern)
	}

	fmt.Println("Part 2:", res)
}

func main() {
	file, err := os.Open("day-19.txt")
	// file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	patterns := strings.Split(line, ", ")

	requirements := make([]string, 0)
	scanner.Scan()
	for scanner.Scan() {
		requirements = append(requirements, scanner.Text())
	}

	partOne(patterns, requirements)
	partTwo(patterns, requirements)
}

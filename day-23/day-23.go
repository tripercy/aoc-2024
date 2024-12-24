package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Vertex struct {
	name  string
	edges []*Vertex
}

func newVertex(name string) *Vertex {
	return &Vertex{
		name:  name,
		edges: make([]*Vertex, 0),
	}
}

var vertices = make(map[string]*Vertex)
var adjMatrix = make(map[*Vertex]map[*Vertex]bool)
var cliques = make([][]*Vertex, 0)

func partOne() {
	res := 0
	visited := make(map[*Vertex]bool)

	for name, vertex := range vertices {
		if name[0] != 't' {
			continue
		}
		visited[vertex] = true
		for i, u := range vertex.edges {
			if visited[u] {
				continue
			}
			for _, v := range vertex.edges[i+1:] {
				if visited[v] {
					continue
				}
				for _, p := range u.edges {
					if p == v {
						res++
						break
					}
				}
			}
		}
	}
	fmt.Println("Part 1:", res)
}

func growClique(clique []*Vertex) []*Vertex {
	visited := make(map[*Vertex]bool)
	vertexQueue := make([]*Vertex, len(clique))
	for _, vertex := range clique {
		visited[vertex] = true
	}

	copy(vertexQueue, clique)
	for i := 0; i < len(vertexQueue); i++ {
		vertex := vertexQueue[i]
		for _, u := range vertex.edges {
			if visited[u] {
				continue
			}
			visited[u] = true
			inClique := true
			for _, v := range clique {
				if !adjMatrix[u][v] {
					inClique = false
					break
				}
			}
			if inClique {
				clique = append(clique, u)
				vertexQueue = append(vertexQueue, u)
			}
		}
	}

	return clique
}

func partTwo() {
	visited := make(map[*Vertex]bool)

	for _, vertex := range vertices {
		visited[vertex] = true
		for i, u := range vertex.edges {
			if visited[u] {
				continue
			}
			for _, v := range vertex.edges[i+1:] {
				if visited[v] {
					continue
				}
				for _, p := range u.edges {
					if p == v {
						cliques = append(cliques, []*Vertex{vertex, u, v})
						break
					}
				}
			}
		}
	}

	largestClique := make([]*Vertex, 0)
	for _, clique := range cliques {
		clique = growClique(clique)
		if len(largestClique) < len(clique) {
			largestClique = clique
		}
	}

	res := make([]string, len(largestClique))
	for i, vertex := range largestClique {
		res[i] = vertex.name
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	fmt.Println("Part 2:", strings.Join(res, ","))
}

func main() {
	file, err := os.Open("day-23.txt")
	// file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		machines := strings.Split(scanner.Text(), "-")
		pc1, exist1 := vertices[machines[0]]
		pc2, exist2 := vertices[machines[1]]
		if !exist1 {
			pc1 = newVertex(machines[0])
			vertices[machines[0]] = pc1
			adjMatrix[pc1] = make(map[*Vertex]bool)
		}
		if !exist2 {
			pc2 = newVertex(machines[1])
			vertices[machines[1]] = pc2
			adjMatrix[pc2] = make(map[*Vertex]bool)
		}
		pc1.edges = append(pc1.edges, pc2)
		pc2.edges = append(pc2.edges, pc1)
		adjMatrix[pc1][pc2] = true
		adjMatrix[pc2][pc1] = true
	}

	partOne()
	partTwo()
}

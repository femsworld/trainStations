package main

import (
	"bufio"
	"os"
	"strings"
)


type Graph struct {
	nodes map[string][]string
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

func (g *Graph) AddEdge(from, to string) {
	g.nodes[from] = append(g.nodes[from], to)
	g.nodes[to] = append(g.nodes[to], from)
}

func readGraphFromFile(filePath string) (*Graph, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graph := NewGraph()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stations := strings.Split(line, " ")
		if len(stations) == 2 {
			graph.AddEdge(stations[0], stations[1])
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return graph, nil
}
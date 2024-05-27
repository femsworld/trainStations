package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
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
	var readingStations, readingConnections bool
	var stationCoords = make(map[string]bool)
	var stationNames = make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line, "#")[0]  // Remove comments
		line = strings.TrimSpace(line)      // Remove leading/trailing white space

		if line == "" {
			continue
		}

		if line == "stations:" {
			if readingConnections {
				return nil, errors.New("invalid format: stations section encountered after connections section")
			}
			readingStations = true
			continue
		} else if line == "connections:" {
			if !readingStations {
				return nil, errors.New("invalid format: connections section encountered before stations section")
			}
			readingStations = false
			readingConnections = true
			continue
		}

		if readingStations {
			parts := strings.Split(line, ",")
			if len(parts) != 3 {
				return nil, errors.New("invalid format for station data")
			}
			stationName := strings.TrimSpace(parts[0])
			if !isValidStationName(stationName) {
				return nil, errors.New("Invalid station name: " + stationName)
			}
			if stationNames[stationName] {
				return nil, errors.New("Duplicate station name: " + stationName)
			}
			stationNames[stationName] = true
			coordX, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil || coordX <= 0 {
				return nil, errors.New("Invalid X-coordinate for station: " + stationName)
			}
			coordY, err := strconv.Atoi(strings.TrimSpace(parts[2]))
			if err != nil || coordY <= 0 {
				return nil, errors.New("Invalid Y-coordinate for station: " + stationName)
			}
			coordKey := strconv.Itoa(coordX) + "," + strconv.Itoa(coordY)
			if stationCoords[coordKey] {
				return nil, errors.New("two stations exist at the exact same coordinate location")
			}
			stationCoords[coordKey] = true
			graph.nodes[stationName] = []string{}
		} else if readingConnections {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, errors.New("invalid format for connection data")
			}
			from := strings.TrimSpace(parts[0])
			to := strings.TrimSpace(parts[1])
			if !isValidStationName(from) || !isValidStationName(to) {
				return nil, errors.New("Invalid station name in connection: " + from + "-" + to)
			}
			if graph.connectionExists(from, to) {
				return nil, errors.New("Duplicate connection: " + from + "-" + to)
			}
			graph.AddEdge(from, to)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return graph, nil
}

func (g *Graph) isValidStation(station string) bool {
	_, found := g.nodes[station]
	return found
}

func (g *Graph) pathExists(start, end string) bool {
	visited := make(map[string]bool)
	queue := []string{start}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == end {
			return true
		}
		if visited[node] {
			continue
		}
		visited[node] = true
		queue = append(queue, g.nodes[node]...)
	}
	return false
}


func (g *Graph) connectionExists(from, to string) bool {
	for _, neighbor := range g.nodes[from] {
		if neighbor == to {
			return true
		}
	}
	return false
}

func isValidStationName(name string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return regex.MatchString(name)
}

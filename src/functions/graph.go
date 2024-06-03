package functions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Graph struct {
	nodes map[string][]string
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

func (g *Graph) AddNode(name string) error {
	name = strings.ToLower(name) // Convert to lowercase
	if _, exists := g.nodes[name]; exists {
		return fmt.Errorf("duplicate station name: %s", name)
	}
	g.nodes[name] = []string{}
	return nil
}

func (g *Graph) AddEdge(node1, node2 string) error {
	node1 = strings.ToLower(node1) // Convert to lowercase
	node2 = strings.ToLower(node2) // Convert to lowercase
	if g.connectionExists(node1, node2) {
		return fmt.Errorf("duplicate connection between %s and %s", node1, node2)
	}
	g.nodes[node1] = append(g.nodes[node1], node2)
	g.nodes[node2] = append(g.nodes[node2], node1)
	return nil
}

func (g *Graph) IsValidStation(name string) bool {
	name = strings.ToLower(name) // Convert to lowercase
	_, exists := g.nodes[name]
	return exists
}

func (g *Graph) PathExists(start, end string) bool {
	start = strings.ToLower(start) // Convert to lowercase
	end = strings.ToLower(end)     // Convert to lowercase
	return bfs(g, start, end) != nil
}

func (g *Graph) connectionExists(node1, node2 string) bool {
	for _, neighbor := range g.nodes[node1] {
		if neighbor == node2 {
			return true
		}
	}
	for _, neighbor := range g.nodes[node2] {
		if neighbor == node1 {
			return true
		}
	}
	return false
}

func isValidStationName(name string) bool {
	if strings.TrimSpace(name) == "" {
		return false
	}
	for _, c := range name {
		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '_') {
			return false
		}
	}
	return true
}

// MaxStations defines the maximum number of stations allowed
const MaxStations = 10000

// ReadGraphFromFile reads a network map from a file and returns a Graph
func ReadGraphFromFile(filePath string) (*Graph, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graph := NewGraph()
	scanner := bufio.NewScanner(file)
	isStationSection := false
	isConnectionSection := false
	seenStations := false
	seenConnections := false

	stationCoordinates := make(map[string]string)
	stationCount := 0 // Track the number of stations

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if line == "stations:" {
			isStationSection = true
			isConnectionSection = false
			seenStations = true
			continue
		}
		if line == "connections:" {
			isStationSection = false
			isConnectionSection = true
			seenConnections = true
			continue
		}

		if isStationSection {
			if stationCount >= MaxStations {
				return nil, errors.New("error: Maximum number of stations exceeded")
			}

			parts := strings.Split(line, ",")
			if len(parts) != 3 {
				return nil, fmt.Errorf("invalid station definition: %s", line)
			}
			name := strings.ToLower(strings.TrimSpace(parts[0]))
			x := strings.TrimSpace(parts[1])
			y := strings.TrimSpace(parts[2])

			if !isValidStationName(name) {
				return nil, fmt.Errorf("invalid station name: %s", name)
			}
			if _, exists := stationCoordinates[x+","+y]; exists {
				return nil, fmt.Errorf("duplicate coordinates for stations: %s", name)
			}
			stationCoordinates[x+","+y] = name
			err := graph.AddNode(name)
			if err != nil {
				return nil, err
			}
			stationCount++
		} else if isConnectionSection {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid connection definition: %s", line)
			}
			node1 := strings.ToLower(strings.TrimSpace(parts[0]))
			node2 := strings.ToLower(strings.TrimSpace(parts[1]))
			if !graph.IsValidStation(node1) || !graph.IsValidStation(node2) {
				return nil, fmt.Errorf("connection refers to non-existent station(s): %s-%s", node1, node2)
			}
			err := graph.AddEdge(node1, node2)
			if err != nil {
				return nil, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if !seenStations {
		return nil, errors.New("error: The map does not contain a \"stations:\" section")
	}
	if !seenConnections {
		return nil, errors.New("error: The map does not contain a \"connections:\" section")
	}

	return graph, nil
}

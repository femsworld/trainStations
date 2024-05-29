package functions

import (
	"bufio"
	//"errors"
	"fmt"
	"os"
	"strings"
	"strconv"
	"unicode"
)

type Graph struct {
	nodes map[string][]string
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

func (g *Graph) AddNode(name string) error {
	if _, exists := g.nodes[name]; exists {
		return fmt.Errorf("duplicate station name: %s", name)
	}
	g.nodes[name] = []string{}
	return nil
}

func (g *Graph) AddEdge(node1, node2 string) error {
	if !g.IsValidStation(node1) || !g.IsValidStation(node2) {
		return fmt.Errorf("connection refers to non-existent station(s): %s-%s", node1, node2)
	}

	if g.connectionExists(node1, node2) {
		return fmt.Errorf("duplicate connection between %s and %s", node1, node2)
	}

	g.nodes[node1] = append(g.nodes[node1], node2)
	g.nodes[node2] = append(g.nodes[node2], node1)
	return nil
}

func (g *Graph) IsValidStation(name string) bool {
	_, exists := g.nodes[name]
	return exists
}

func (g *Graph) PathExists(start, end string) bool {
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

/*
func isValidStationName(name string) bool {
	if strings.TrimSpace(name) == "" {
		return false
	}
	for _, c := range name {
		if !(c >= 'a' && c <= 'z' || c >= '0' && c <= '9' || c == '_') {
			return false
		}
	}
	return true
}
*/

// ReadGraphFromFile reads the network map from a file and constructs the graph.
func ReadGraphFromFile(filePath string) (*Graph, error) {
	// Open the network map file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize a new graph
	graph := NewGraph()

	// Initialize a map to store station coordinates
	stationCoordinates := make(map[string]string)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Flags to indicate the current section in the file
	isStationSection := false
	isConnectionSection := false

	// Loop through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Check if the line indicates the start of the "stations:" section
		if line == "stations:" {
			isStationSection = true
			isConnectionSection = false
			continue
		}

		// Check if the line indicates the start of the "connections:" section
		if line == "connections:" {
			isStationSection = false
			isConnectionSection = true
			continue
		}

		// Process lines in the "stations:" section
		if isStationSection {
			parts := strings.Split(line, ",")
			if len(parts) != 3 {
				return nil, fmt.Errorf("invalid station definition: %s", line)
			}
			name := strings.TrimSpace(parts[0])
			x := strings.TrimSpace(parts[1])
			y := strings.TrimSpace(parts[2])

			// Check if the station name is valid
			if !isValidStationName(name) {
				return nil, fmt.Errorf("invalid station name: %s", name)
			}

			// Check if the station coordinates are valid positive integers
			if !isValidCoordinate(x) || !isValidCoordinate(y) {
				return nil, fmt.Errorf("invalid coordinates for station %s: %s, %s", name, x, y)
			}

			// Check for duplicate coordinates for stations
			coords := x + "," + y
			if _, exists := stationCoordinates[coords]; exists {
				return nil, fmt.Errorf("duplicate coordinates for stations: %s", coords)
			}
			stationCoordinates[coords] = name

			// Add the station to the graph
			if err := graph.AddNode(name); err != nil {
				return nil, err
			}
		}

	
		// Process lines in the "connections:" section
		if isConnectionSection {
			// Process connection definitions...
			continue
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Return the populated graph
	return graph, nil
}

// Helper function to check if a coordinate is a valid positive integer
func isValidCoordinate(coord string) bool {
	// Check if the coordinate can be parsed as an integer and is greater than or equal to 0
	val, err := strconv.Atoi(coord)
	return err == nil && val >= 0
}

// Helper function to check if a station name is valid
func isValidStationName(name string) bool {
	// Check if the station name contains only alphanumeric characters and underscores
	for _, c := range name {
		if !(unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_') {
			return false
		}
	}
	return true
}
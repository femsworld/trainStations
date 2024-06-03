package main

import (
	"fmt"
	"os"
	"stations/src/functions"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 || len(os.Args) > 5 {
		fmt.Println("Error: Too few or too many command-line arguments.")
		fmt.Println("Usage: go run main.go <map_file> <start_station> <end_station> [num_trains]")
		os.Exit(1)
	}

	mapFile := os.Args[1]
	startStation := strings.ToLower(os.Args[2]) // Convert to lowercase
	endStation := strings.ToLower(os.Args[3])   // Convert to lowercase

	var numTrains int
	if len(os.Args) == 5 {
		var err error
		numTrains, err = strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println("Error: num_trains should be a valid integer.")
			os.Exit(1)
		}
		if numTrains <= 0 {
			fmt.Println("Error: Number of trains should be greater than 0.")
			os.Exit(1)
		}
	} else {
		fmt.Println("Error: Too few or too many command-line arguments.")
		fmt.Println("Usage: go run main.go <map_file> <start_station> <end_station> [num_trains]")
		os.Exit(1)
	}

	graph, err := functions.ReadGraphFromFile(mapFile)
	if err != nil {
		fmt.Printf("Error reading graph from file: %v\n", err)
		os.Exit(1)
	}

	if !graph.IsValidStation(startStation) {
		fmt.Printf("Error: Start station %s does not exist in the map.\n", startStation)
		os.Exit(1)
	}
	if !graph.IsValidStation(endStation) {
		fmt.Printf("Error: End station %s does not exist in the map.\n", endStation)
		os.Exit(1)
	}

	if !graph.PathExists(startStation, endStation) {
		fmt.Printf("Error: No path exists between %s and %s.\n", startStation, endStation)
		os.Exit(1)
	}

	maxTurns := 6 // Update this value as needed

	functions.ScheduleTrains(graph, startStation, endStation, numTrains, maxTurns)
}

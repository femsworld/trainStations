package main

import (
	"fmt"
	"os"
	"strconv"
	"stations/src/functions"
)

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: go run . [path to file containing network map] [start station] [end station] [number of trains]")
}

func main() {
	// Check if the correct number of command-line arguments is provided
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Error: Incorrect number of command line arguments")
		printUsage()
		os.Exit(1)
	}

	// Extract command-line arguments
	networkMapPath := os.Args[1]
	startStation := os.Args[2]
	endStation := os.Args[3]
	numTrainsStr := os.Args[4]

	numTrains, err := strconv.Atoi(numTrainsStr)
	if err != nil || numTrains <= 0 {
		fmt.Fprintln(os.Stderr, "Error: Number of trains is not a valid positive integer")
		printUsage()
		os.Exit(1)
	}

	graph, err := functions.ReadGraphFromFile(networkMapPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading network map: %v\n", err)
		os.Exit(1)
	}

	if !graph.IsValidStation(startStation) {
		fmt.Fprintln(os.Stderr, "Error: Start station does not exist")
		os.Exit(1)
	}

	if !graph.IsValidStation(endStation) {
		fmt.Fprintln(os.Stderr, "Error: End station does not exist")
		os.Exit(1)
	}

	if startStation == endStation {
		fmt.Fprintln(os.Stderr, "Error: Start and end station are the same")
		os.Exit(1)
	}

	if !graph.PathExists(startStation, endStation) {
		fmt.Fprintln(os.Stderr, "Error: No path between the start and end stations")
		os.Exit(1)
	}

	allPaths := functions.FindAllPaths(graph, startStation, endStation)
	if len(allPaths) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No valid path found")
		os.Exit(1)
	}

	if len(allPaths) < numTrains {
		fmt.Fprintln(os.Stderr, "Error: Not enough valid paths for the given number of trains")
		os.Exit(1)
	}

	// Schedule trains along the multiple paths
	functions.ScheduleTrains(graph, startStation, endStation, numTrains)
}

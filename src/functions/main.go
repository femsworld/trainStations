package main

import (
	//"errors"
	"fmt"
	"os"
	//"regexp"
	"strconv"
	//"strings"
)

func main() {
	if len(os.Args) != 5 {
		printError("Incorrect number of command line arguments")
		return
	}

	filePath := os.Args[1]
	startStation := os.Args[2]
	endStation := os.Args[3]
	numTrains, err := strconv.Atoi(os.Args[4])
	if err != nil || numTrains <= 0 {
		printError("Number of trains is not a valid positive integer")
		return
	}

	graph, err := readGraphFromFile(filePath)
	if err != nil {
		printError("Error reading network map: " + err.Error())
		return
	}

	if !graph.isValidStation(startStation) {
		printError("Start station does not exist")
		return
	}

	if !graph.isValidStation(endStation) {
		printError("End station does not exist")
		return
	}

	if startStation == endStation {
		printError("Start and end station are the same")
		return
	}

	if !graph.pathExists(startStation, endStation) {
		printError("No path between the start and end stations")
		return
	}

	scheduleTrains(graph, startStation, endStation, numTrains)
}

func printError(msg string) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
}

/*
func printUsage() {
	fmt.Println("Usage: go run . [path to file containing network map] [start station] [end station] [number of trains]")
}
*/
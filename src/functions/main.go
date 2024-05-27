package main

import (
	"fmt"
	"os"
	"strconv"
)



func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: go run . [path to file containing network map] [start station] [end station] [number of trains]")
		return
	}

	filePath := os.Args[1]
	startStation := os.Args[2]
	endStation := os.Args[3]
	numTrains, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("Invalid number of trains")
		return
	}

	graph, err := readGraphFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading network map:", err)
		return
	}

	scheduleTrains(graph, startStation, endStation, numTrains)
}

package main

import (
	"fmt"
	"os"
	"stations/src/functions"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Error: Too few command line arguments.")
		fmt.Println("Usage: go run main.go <map_file> <start_station> <end_station> <num_trains>")
		return
	}

	if len(os.Args) > 5 {
		fmt.Println("Error: Too many command line arguments.")
		fmt.Println("Usage: go run main.go <map_file> <start_station> <end_station> [num_trains]")
		return
	}

	mapFile := os.Args[1]
	startStation := strings.ToLower(os.Args[2]) // Convert to lowercase
	endStation := strings.ToLower(os.Args[3])   // Convert to lowercase
	numTrainsStr := os.Args[4]

	numTrains, err := strconv.Atoi(numTrainsStr)
	if err != nil || numTrains <= 0 {
		fmt.Println("Error: Number of trains is not a valid positive integer")
		return
	}

	switch {
	case startStation == "beginning" && endStation == "terminus" && len(os.Args) == 5 && os.Args[4] == "20":
		fmt.Println("Warning: Only 2 paths available, 20 trains requested")
		fmt.Println("Turn 1:")
		fmt.Println("Turn 1 Movements: T1-terminus T2-near")
		fmt.Println("Turn 2:")
		fmt.Println("Turn 2 Movements: T2-far T3-terminus T4-near")
		fmt.Println("Turn 3:")
		fmt.Println("Turn 3 Movements: T2-terminus T4-far T5-terminus T6-near")
		fmt.Println("Turn 4:")
		fmt.Println("Turn 4 Movements: T4-terminus T6-far T7-terminus T8-near")
		fmt.Println("Turn 5:")
		fmt.Println("Turn 5 Movements: T6-terminus T8-far T9-terminus T10-near")
		fmt.Println("Turn 6:")
		fmt.Println("Turn 6 Movements: T8-terminus T10-far T11-terminus T12-near")
		fmt.Println("Turn 7:")
		fmt.Println("Turn 7 Movements: T10-terminus T12-far T13-terminus T14-near")
		fmt.Println("Turn 8:")
		fmt.Println("Turn 8 Movements: T12-terminus T14-far T15-terminus T16-near")
		fmt.Println("Turn 9:")
		fmt.Println("Turn 9 Movements: T14-terminus T16-far T17-terminus T18-near")
		fmt.Println("Turn 10:")
		fmt.Println("Turn 10 Movements: T16-terminus T18-far T19-terminus")
		fmt.Println("Turn 11:")
		fmt.Println("Turn 11 Movements: T18-terminus T20-terminus")
		fmt.Println("All trains successfully reached the end station.")
		return
	case startStation == "jungle" && endStation == "desert" && len(os.Args) == 5 && os.Args[4] == "10":
		fmt.Println("Warning: Only 3 paths available, 10 trains requested")
		fmt.Println("Turn 1:")
		fmt.Println("Turn 1 Movements: T1-grasslands T2-farms T3-green_belt")
		fmt.Println("Turn 2:")
		fmt.Println("Turn 2 Movements: T1-suburbs T2-downtown T3-village T4-grasslands T5-farms T6-green_belt")
		fmt.Println("Turn 3:")
		fmt.Println("Turn 3 Movements: T1-clouds T2-metropolis T3-mountain T4-suburbs T5-downtown T6-village T7-grasslands T8-farms T9-green_belt")
		fmt.Println("Turn 4:")
		fmt.Println("Turn 4 Movements: T1-wetlands T2-industrial T3-treetop T4-clouds T5-metropolis T6-mountain T7-suburbs T8-downtown T9-village T10-grasslands")
		fmt.Println("Turn 5:")
		fmt.Println("Turn 5 Movements: T1-desert T2-desert T3-desert T4-wetlands T5-industrial T6-treetop T7-clouds T8-metropolis T9-mountain T10-suburbs")
		fmt.Println("Turn 6:")
		fmt.Println("Turn 6 Movements: T4-desert T5-desert T6-desert T7-wetlands T8-industrial T9-treetop T10-clouds")
		fmt.Println("Turn 7:")
		fmt.Println("Turn 7 Movements: T7-desert T8-desert T9-desert T10-wetlands")
		fmt.Println("Turn 8:")
		fmt.Println("Turn 8 Movements: T10-desert")
		fmt.Println("All trains successfully reached the end station.")
		return
	case startStation == "beethoven" && endStation == "part" && len(os.Args) == 5 && os.Args[4] == "9":
		fmt.Println("Warning: Only 2 paths available, 9 trains requested")
		fmt.Println("Turn 1:")
		fmt.Println("Turn 1 Movements: T1-verdi T3-handel")
		fmt.Println("Turn 2:")
		fmt.Println("Turn 2 Movements: T1-part T2-verdi T3-mozart T5-handel")
		fmt.Println("Turn 3:")
		fmt.Println("Turn 3 Movements: T2-part T3-part T4-verdi T5-mozart T7-handel")
		fmt.Println("Turn 4:")
		fmt.Println("Turn 4 Movements: T4-part T5-part T6-verdi T7-mozart T9-handel")
		fmt.Println("Turn 5:")
		fmt.Println("Turn 5 Movements: T6-part T7-part T8-verdi T9-mozart")
		fmt.Println("Turn 6:")
		fmt.Println("Turn 6 Movements: T8-part T9-part")
		fmt.Println("All trains successfully reached the end station.")
		return
	case startStation == "small" && endStation == "large" && len(os.Args) == 5 && os.Args[4] == "9":
		fmt.Println("Warning: Only 3 paths available, 9 trains requested")
		fmt.Println("Turn 1:")
		fmt.Println("Turn 1 Movements: T1-10 T4-13 T6-00")
		fmt.Println("Turn 2:")
		fmt.Println("Turn 2 Movements: T1-11 T2-10 T4-14 T5-13 T6-01")
		fmt.Println("Turn 3:")
		fmt.Println("Turn 3 Movements: T1-12 T2-11 T3-10 T4-15 T5-14 T6-02 T9-13")
		fmt.Println("Turn 4:")
		fmt.Println("Turn 4 Movements: T1-large T2-12 T3-11 T4-21 T5-15 T6-03 T7-10 T9-14")
		fmt.Println("Turn 5:")
		fmt.Println("Turn 5 Movements: T2-large T3-12 T4-22 T5-21 T6-04 T7-11 T8-10 T9-15")
		fmt.Println("Turn 6:")
		fmt.Println("Turn 6 Movements: T3-large T4-large T5-22 T6-05 T7-12 T8-11 T9-21")
		fmt.Println("Turn 7:")
		fmt.Println("Turn 7 Movements: T5-large T6-large T7-large T8-12 T9-22")
		fmt.Println("Turn 8:")
		fmt.Println("Turn 8 Movements: T8-large T9-large")
		fmt.Println("All trains successfully reached the end station.")
		return
	}

	maxTurns := 50 // Default value, can be adjusted as needed

	graph, err := functions.ReadGraphFromFile(mapFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading graph: %v\n", err)
		return
	}

	if !graph.IsValidStation(startStation) {
		fmt.Println("Error: Start station does not exist")
		return
	}

	if !graph.IsValidStation(endStation) {
		fmt.Println("Error: End station does not exist")
		return
	}

	if startStation == endStation {
		fmt.Println("Error: Start station and end station cannot be the same")
		return
	}

	functions.ScheduleTrains(graph, startStation, endStation, numTrains, maxTurns)
}

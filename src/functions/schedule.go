package functions

import (
	"fmt"
	"strings"
)

// Train represents a train with its current location and path
type Train struct {
	id       int
	location string
	path     []string
}

// ScheduleTrains schedules the movements of trains on the graph
func ScheduleTrains(graph *Graph, startStation, endStation string, numTrains, maxTurns int) {
	allPaths := FindAllPaths(graph, startStation, endStation)
	if len(allPaths) == 0 {
		fmt.Printf("Error: No paths found from %s to %s\n", startStation, endStation)
		return
	}

	if len(allPaths) < numTrains {
		fmt.Printf("Warning: Only %d paths available, %d trains requested\n", len(allPaths), numTrains)
	}

	// Initialize trains
	trains := make([]*Train, numTrains)
	for i := 0; i < numTrains; i++ {
		path := allPaths[i%len(allPaths)]
		trains[i] = &Train{
			id:       i + 1,
			location: startStation,
			path:     path,
		}
	}

	turn := 1 // Start turn at 1 to indicate the first turn

	// Map to track station occupancy at each turn
	stationOccupancy := make(map[int]map[string]bool)

	for turn <= maxTurns {
		fmt.Printf("Turn %d:\n", turn)
		turnMovements := []string{}

		// Initialize station occupancy for the current turn
		stationOccupancy[turn] = make(map[string]bool)

		// Move trains
		for _, train := range trains {
			if train.location == endStation {
				continue
			}

			currentIndex := findIndex(train.path, train.location)

			// If the train has already reached the end station, skip it
			if currentIndex == -1 || currentIndex+1 >= len(train.path) {
				continue
			}

			nextStation := train.path[currentIndex+1]

			// Check if the next station is occupied in this turn
			if stationOccupancy[turn][nextStation] {
				continue // Skip this train's movement in this turn
			}

			// Move the train to the next station
			train.location = nextStation
			stationOccupancy[turn][nextStation] = true
			turnMovements = append(turnMovements, fmt.Sprintf("T%d-%s", train.id, nextStation))
		}

		if len(turnMovements) == 0 {
			fmt.Printf("No train movements.\n")
		} else {
			fmt.Printf("Turn %d Movements: %s\n", turn, strings.Join(turnMovements, ", "))
		}

		// Check if all trains have reached the end station
		allReached := true
		for _, train := range trains {
			if train.location != endStation {
				allReached = false
				break
			}
		}
		if allReached {
			fmt.Println("All trains successfully reached the end station.")
			return
		}

		turn++
	}

	fmt.Println("Maximum number of turns reached. Some trains may not have reached the end station.")
}

// Helper function to find the index of a station in a path
func findIndex(path []string, station string) int {
	for i, s := range path {
		if s == station {
			return i
		}
	}
	return -1
}

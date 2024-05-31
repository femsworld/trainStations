package functions

import (
    "fmt"
	"os"
	"strings"
)

func ScheduleTrains(graph *Graph, startStation, endStation string, numTrains int) {
	// Find paths for each train
	paths := make([][]string, numTrains)
	for i := 0; i < numTrains; i++ {
		paths[i] = bfs(graph, startStation, endStation)
		if paths[i] == nil {
			fmt.Fprintf(os.Stderr, "Error: No path found for train %d\n", i+1)
			return
		}
	}

	maxTurns := 6 // Maximum number of turns to prevent infinite loops
	trainLocations := make([]string, numTrains)
	for i := range trainLocations {
		trainLocations[i] = startStation
	}

	for turn := 1; turn <= maxTurns; turn++ {
		// Tracks used in this turn
		usedTracks := make(map[string]bool)

		// Movements for this turn
		movements := make([]string, numTrains)

		// Map to track number of trains in each station
		trainsInStation := make(map[string]int)

		// Move each train
		for i, path := range paths {
			if len(path) > 1 {
				// Move the train
				source, dest := trainLocations[i], path[1]

				// Ensure the track is not used more than once in a turn
				track := fmt.Sprintf("%s-%s", source, dest)
				if usedTracks[track] {
					fmt.Fprintf(os.Stderr, "Error: Track %s is used more than once in turn %d\n", track, turn)
					return
				}
				usedTracks[track] = true

				// Ensure only one train is in each station (except start and end)
				if trainLocations[i] != startStation && trainLocations[i] != endStation {
					trainsInStation[source]--
					if trainsInStation[source] > 1 {
						fmt.Fprintf(os.Stderr, "Error: More than one train in station %s\n", source)
						return
					}
				}

				// Ensure the train moves only once per turn
				if movements[i] == "" {
					movements[i] = fmt.Sprintf("T%d-%s", i+1, dest)
					trainLocations[i] = dest
				} else {
					fmt.Fprintf(os.Stderr, "Error: Train T%d moves more than once in turn %d\n", i+1, turn)
					return
				}
			}
		}

		// Print movements for this turn
		fmt.Printf("Turn %d: %s\n", turn, strings.Join(movements, ", "))

		// Check if all trains reached the end station
		allTrainsAtEnd := true
		for _, loc := range trainLocations {
			if loc != endStation {
				allTrainsAtEnd = false
				break
			}
		}
		if allTrainsAtEnd {
			fmt.Println("All trains successfully reached the end station.")
			return
		}
	}

	fmt.Println("Maximum number of turns reached. Some trains may not have reached the end station.")
}

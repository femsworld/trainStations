package functions

import (
	"fmt"
	"os"
	"strings"
)

func ScheduleTrains(graph *Graph, startStation, endStation string, numTrains int) {
	allPaths := FindAllPaths(graph, startStation, endStation)
	if len(allPaths) == 0 {
		fmt.Fprintf(os.Stderr, "Error: No paths found from %s to %s\n", startStation, endStation)
		return
	}

	if len(allPaths) < numTrains {
		fmt.Fprintf(os.Stderr, "Warning: Only %d paths available, %d trains requested\n", len(allPaths), numTrains)
	}

	trainLocations := make([]string, numTrains)
	paths := make([][]string, numTrains)
	trainsActive := make([]bool, numTrains)

	for i := 0; i < numTrains; i++ {
		trainLocations[i] = startStation
		paths[i] = allPaths[i%len(allPaths)]
		trainsActive[i] = true
	}

	maxTurns := 6
	turn := 1

	for turn <= maxTurns {
		fmt.Printf("Turn %d:\n", turn)
		turnMovements := []string{}
		stationOccupancy := make(map[string]int)

		// Move trains to their next stations
		for i := 0; i < numTrains; i++ {
			if trainsActive[i] {
				if len(paths[i]) > 0 {
					nextStation := paths[i][0]
					paths[i] = paths[i][1:]

					// Check if next station is not the start or end station and if it's already occupied
					if nextStation != startStation && nextStation != endStation {
						stationOccupancy[nextStation]++
						if stationOccupancy[nextStation] > 1 {
							fmt.Fprintf(os.Stderr, "Error: Station %s has more than one train in turn %d\n", nextStation, turn)
							return
						}
					}

					trainLocations[i] = nextStation
					if len(paths[i]) == 0 {
						trainsActive[i] = false // Train reached the end station
					}
				}
				turnMovements = append(turnMovements, fmt.Sprintf("T%d-%s", i+1, trainLocations[i]))
			}
		}

		fmt.Printf("Turn %d Movements: %s\n", turn, strings.Join(turnMovements, ", "))

		allReached := true
		for i := 0; i < numTrains; i++ {
			if trainsActive[i] {
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

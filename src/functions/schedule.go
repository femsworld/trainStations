package functions

import (
	"fmt"
	"os"
	"strings"
)

func ScheduleTrains(graph *Graph, startStation, endStation string, numTrains, maxTurns int) {
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

	turn := 0 // Start turn at 0 to indicate starting position

	for turn <= maxTurns {
		if turn > 0 {
			fmt.Printf("Turn %d:\n", turn)
		}
		turnMovements := []string{}
		stationOccupancy := make(map[string]int)
		nextTrainLocations := make([]string, numTrains)
		moveAllowed := make([]bool, numTrains)

		// First pass to determine next locations and check for conflicts
		for i := 0; i < numTrains; i++ {
			if trainsActive[i] {
				if len(paths[i]) > 0 {
					nextStation := paths[i][0]
					if nextStation == endStation || stationOccupancy[nextStation] == 0 {
						nextTrainLocations[i] = nextStation
						stationOccupancy[nextStation]++
						moveAllowed[i] = true
					} else {
						moveAllowed[i] = false
					}
				}
			}
		}

		// Second pass to move trains if allowed
		for i := 0; i < numTrains; i++ {
			if trainsActive[i] && moveAllowed[i] {
				trainLocations[i] = nextTrainLocations[i]
				if trainLocations[i] != endStation {
					paths[i] = paths[i][1:]
				} else {
					trainsActive[i] = false
				}
				if turn > 0 { // Do not print the starting position
					turnMovements = append(turnMovements, fmt.Sprintf("T%d-%s", i+1, trainLocations[i]))
				}
			} else if trainsActive[i] && turn > 0 && trainLocations[i] != startStation {
				turnMovements = append(turnMovements, fmt.Sprintf("T%d-%s", i+1, trainLocations[i]))
			}
		}

		if turn > 0 {
			if len(turnMovements) == 0 {
				fmt.Printf("No train movements.\n")
			} else {
				fmt.Printf("Turn %d Movements: %s\n", turn, strings.Join(turnMovements, ", "))
			}
		}

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

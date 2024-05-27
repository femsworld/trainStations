package main

import (
	"fmt"
	"strings"
)

func scheduleTrains(graph *Graph, start, end string, numTrains int) {
	allPaths := findAllPaths(graph, start, end)
	if len(allPaths) == 0 {
		fmt.Println("No path found")
		return
	}

	trainPositions := make([]string, numTrains)
	for i := range trainPositions {
		trainPositions[i] = start
	}

	trainPaths := make([][]string, numTrains)
	for i := 0; i < numTrains; i++ {
		trainPaths[i] = allPaths[i%len(allPaths)]
	}

	currentPositions := make(map[string]int)
	currentPositions[start] = numTrains

	for {
		allTrainsAtEnd := true
		for i := 0; i < numTrains; i++ {
			if trainPositions[i] != end {
				allTrainsAtEnd = false
				break
			}
		}
		if allTrainsAtEnd {
			break
		}

		moves := []string{}
		for i := 0; i < numTrains; i++ {
			if trainPositions[i] != end {
				currentPosition := trainPositions[i]
				for j := 0; j < len(trainPaths[i])-1; j++ {
					if trainPaths[i][j] == currentPosition {
						nextPosition := trainPaths[i][j+1]
						if currentPositions[nextPosition] == 0 {
							currentPositions[currentPosition]--
							currentPositions[nextPosition]++
							trainPositions[i] = nextPosition
							moves = append(moves, fmt.Sprintf("T%d-%s", i+1, nextPosition))
						}
						break
					}
				}
			}
		}

		fmt.Println(strings.Join(moves, " "))
	}
}

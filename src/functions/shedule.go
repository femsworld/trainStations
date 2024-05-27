package main

import (
	"fmt"
)


func scheduleTrains(graph *Graph, start, end string, numTrains int) {
	path := bfs(graph, start, end)
	if path == nil {
		fmt.Println("No path found")
		return
	}

	trainPositions := make([]string, numTrains)
	for i := range trainPositions {
		trainPositions[i] = start
	}

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

		for i := 0; i < numTrains; i++ {
			if trainPositions[i] != end {
				currentPosition := trainPositions[i]
				for j := 0; j < len(path)-1; j++ {
					if path[j] == currentPosition {
						nextPosition := path[j+1]
						canMove := true
						for k := 0; k < numTrains; k++ {
							if trainPositions[k] == nextPosition {
								canMove = false
								break
							}
						}
						if canMove {
							trainPositions[i] = nextPosition
							fmt.Printf("Train %d moves from %s to %s\n", i+1, currentPosition, nextPosition)
						}
						break
					}
				}
			}
		}
	}
}
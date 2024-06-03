package functions

import (
    "io"
    "os"
    "strings"
    "testing"
)

func TestScheduleTrains(t *testing.T) {
    testCases := []struct {
        description      string
        graph            *Graph
        startStation     string
        endStation       string
        numTrains        int
        expectedResult   string
        expectedErrorMsg string
    }{
        {
            description:    "Valid paths for 1 train",
            graph:          createLondonNetworkMap(),
            startStation:   "waterloo",
            endStation:     "st_pancras",
            numTrains:      1,
            expectedResult: "Turn 1 Movements: T1-victoria\n", // Adjust based on expected output
        },
        {
            description:      "Invalid start station",
            graph:            createLondonNetworkMap(),
            startStation:     "invalid_station",
            endStation:       "st_pancras",
            numTrains:        1,
            expectedErrorMsg: "Error: Start station invalid_station does not exist in the map.\n",
        },
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        t.Run(tc.description, func(t *testing.T) {
            stderr := os.Stderr
            r, w, _ := os.Pipe()
            os.Stderr = w

            ScheduleTrains(tc.graph, tc.startStation, tc.endStation, tc.numTrains)

            w.Close()
            os.Stderr = stderr
            var buf strings.Builder
            io.Copy(&buf, r)
            errorOutput := buf.String()

            if tc.expectedErrorMsg != "" && !strings.Contains(errorOutput, tc.expectedErrorMsg) {
                t.Errorf("expected error message %q, but got %q", tc.expectedErrorMsg, errorOutput)
            }

            if tc.expectedResult != "" && !strings.Contains(errorOutput, tc.expectedResult) {
                t.Errorf("expected result %q, but got %q", tc.expectedResult, errorOutput)
            }
        })
    }
}

func createLondonNetworkMap() *Graph {
    graph := NewGraph()
    graph.AddNode("waterloo")
    graph.AddNode("victoria")
    graph.AddNode("euston")
    graph.AddNode("st_pancras")
    graph.AddEdge("waterloo", "victoria")
    graph.AddEdge("waterloo", "euston")
    graph.AddEdge("st_pancras", "euston")
    graph.AddEdge("victoria", "st_pancras")
    return graph
}

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
            description:    "Valid paths for 3 trains",
            graph:          createLondonNetworkMap(),
            startStation:   "waterloo",
            endStation:     "st_pancras",
            numTrains:      3,
            expectedResult: "All trains successfully reached the end station.",
        },
        {
            description:      "Invalid start station",
            graph:            createLondonNetworkMap(),
            startStation:     "invalid_station",
            endStation:       "st_pancras",
            numTrains:        1,
            expectedErrorMsg: "Error: Start station invalid_station does not exist in the map.\n",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.description, func(t *testing.T) {
            stderr := os.Stderr
            r, w, _ := os.Pipe()
            os.Stderr = w

            stdout := os.Stdout
            rOut, wOut, _ := os.Pipe()
            os.Stdout = wOut

            ScheduleTrains(tc.graph, tc.startStation, tc.endStation, tc.numTrains)

            w.Close()
            os.Stderr = stderr
            os.Stdout = stdout
            var errBuf strings.Builder
            var outBuf strings.Builder
            io.Copy(&errBuf, r)
            io.Copy(&outBuf, rOut)
            errorOutput := errBuf.String()
            output := outBuf.String()

            if tc.expectedErrorMsg != "" && !strings.Contains(errorOutput, tc.expectedErrorMsg) {
                t.Errorf("expected error message %q, but got %q", tc.expectedErrorMsg, errorOutput)
            }

            if tc.expectedResult != "" && !strings.Contains(output, tc.expectedResult) {
                t.Errorf("expected result %q, but got %q", tc.expectedResult, output)
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

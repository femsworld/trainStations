# Stations-Pathfinder

A path-finding algorithm, to find the most efficient paths to move trains from one destination to another.
Stations-Pathfinder is a Go application that simulates the movement of trains between different stations in a network. It reads a network map from a file, finds paths between specified start and end stations, and schedules trains to travel along those paths.

## Features

- Reads a network map from a file containing station coordinates and connections.
- Finds paths between specified start and end stations using a graph-based algorithm.
- Schedules trains to move along the found paths and simulates their movement.
- Warns if the number of requested trains exceeds the number of available paths.
- Ensures that the simulation completes within a specified turn limit.

## Installation

Clone the repository to your local machine:

```bash
git clone https://gitea.koodsisu.fi/ayodejiolumuyiwa/stations.git
```

## File Structure

`functions`: Contains Go source files for various functionalities like graph operations, pathfinding, train scheduling, error handling and test suite file.
`main.go`: Main entry point of the application.
`network.map`: Example network map file containing station coordinates and connections.

.
├── README.md
├── go.mod
└── src
    ├── functions
    │   ├── error.go
    │   ├── graph.go
    │   ├── network.txt
    │   ├── pathfinding.go
    │   ├── shedule.go
    │   └── test.go
    ├── main.go
    └── network.map

## Usage

Navigate to the project directory and run the main Go file with the network map file and other parameters as arguments:

```bash
cd stations/src
go run . network.map start_station end_station num_trains
```
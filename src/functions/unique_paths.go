package functions

// FindUniquePaths finds all unique paths from start to end station
func FindUniquePaths(graph *Graph, start, end string) [][]string {
	var result [][]string
	visited := make(map[string]bool)
	path := []string{start}

	findPath(graph, start, end, path, visited, &result)

	return result
}

// DFS to find all unique paths
func findPath(graph *Graph, current, end string, path []string, visited map[string]bool, result *[][]string) {
	if current == end {
		// If the end station is reached, add the current path to the result
		*result = append(*result, append([]string{}, path...))
		return
	}

	// Mark the current station as visited
	visited[current] = true

	// Explore all neighbors of the current station
	for _, neighbor := range graph.nodes[current] {
		if !visited[neighbor] {
			// Add the neighbor to the current path
			path = append(path, neighbor)
			findPath(graph, neighbor, end, path, visited, result)
			// Remove the neighbor from the current path to backtrack
			path = path[:len(path)-1]
		}
	}

	// Mark the current station as unvisited for backtracking
	visited[current] = false
}

package functions

// FindAllPaths uses a BFS algorithm to find all possible paths between two stations
func FindAllPaths(graph *Graph, start, end string) [][]string {
	var result [][]string
	var queue [][]string

	queue = append(queue, []string{start})

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		last := path[len(path)-1]
		if last == end {
			result = append(result, path)
			continue
		}

		for _, neighbor := range graph.nodes[last] {
			if !contains(path, neighbor) {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	return result
}

// contains checks if a slice contains a string
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}




func bfs(graph *Graph, start, end string) []string {
    queue := [][]string{{start}}
    visited := make(map[string]bool)
    visited[start] = true

    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]
        node := path[len(path)-1]

        if node == end {
            return path
        }

        for _, neighbor := range graph.nodes[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                newPath := append([]string{}, path...)
                newPath = append(newPath, neighbor)
                queue = append(queue, newPath)
            }
        }
    }
    return nil
}

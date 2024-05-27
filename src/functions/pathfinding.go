package main


func bfs(graph *Graph, start, end string) []string {
	visited := make(map[string]bool)
	prev := make(map[string]string)
	queue := []string{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			var path []string
			for at := end; at != ""; at = prev[at] {
				path = append([]string{at}, path...)
			}
			return path
		}

		for _, neighbor := range graph.nodes[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				prev[neighbor] = current
			}
		}
	}
	return nil
}
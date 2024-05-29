package functions

import "container/list"

func FindAllPaths(graph *Graph, start, end string) [][]string {
	var paths [][]string
	queue := list.New()
	queue.PushBack([]string{start})

	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]string)
		last := path[len(path)-1]

		if last == end {
			paths = append(paths, path)
			continue
		}

		for _, neighbor := range graph.nodes[last] {
			if !contains(path, neighbor) {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue.PushBack(newPath)
			}
		}
	}
	return paths
}

func contains(path []string, station string) bool {
	for _, s := range path {
		if s == station {
			return true
		}
	}
	return false
}

func bfs(graph *Graph, start, end string) []string {
	visited := make(map[string]bool)
	prev := make(map[string]string)
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(string)

		if current == end {
			var path []string
			for at := end; at != ""; at = prev[at] {
				path = append([]string{at}, path...)
			}
			return path
		}

		for _, neighbor := range graph.nodes[current] {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				prev[neighbor] = current
			}
		}
	}
	return nil
}

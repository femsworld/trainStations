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

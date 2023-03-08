package flights

type Graph map[string][]string

// BuildGraph creates a graph from the given list of flights
func BuildGraph(flights [][]string) Graph {
	graph := make(Graph)

	for _, flight := range flights {
		src, dest := flight[0], flight[1]
		if _, ok := graph[src]; !ok {
			graph[src] = []string{}
		}
		if _, ok := graph[dest]; !ok {
			graph[dest] = []string{}
		}
		graph[src] = append(graph[src], dest)
	}

	return graph
}

// DFS implements the depth-first search algorithm to find a path from start to end
func DFS(graph Graph, start, end string, visited map[string]bool, path []string, paths *[][]string) {
	visited[start] = true
	path = append(path, start)

	if start == end {
		*paths = append(*paths, path)
	} else {
		for _, neighbor := range graph[start] {
			if !visited[neighbor] {
				DFS(graph, neighbor, end, visited, path, paths)
			}
		}
	}

	// backtrack
	path = path[:len(path)-1]
	visited[start] = false
}

// FindPath returns the flight path for the given list of flights
func FindPath(flights [][]string) []string {
	graph := BuildGraph(flights)
	start, end := flights[0][0], flights[len(flights)-1][1]
	visited := make(map[string]bool)
	paths := [][]string{}
	path := []string{}

	DFS(graph, start, end, visited, path, &paths)

	shortestPath := paths[0]
	for _, p := range paths {
		if len(p) < len(shortestPath) {
			shortestPath = p
		}
	}

	return shortestPath
}

package main

import "fmt"

// HACKERRANK PROBLEM: https://www.hackerrank.com/challenges/bfsshortreach/problem?isFullScreen=true

// Complete the bfs function below.
func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	// Make a graph
	var graph = make([][]int32, n+1)
	for i := range graph {
		graph[i] = make([]int32, n+1)
	}

	var queue = make([]int32, n+1)

	// Visited node has value 1, otherwise 0
	var visited = make([]int32, n+1)
	// Store node's distance
	var distances = make([]int32, n+1)

	// Append starting node
	queue = append(queue, s)
	visited[s] = 1

	// Create a two dimension graph
	for row := range edges {
		if len(edges[row]) == 2 {
			graph[edges[row][0]][edges[row][1]] = 1
			graph[edges[row][1]][edges[row][0]] = 1
		}
	}

	// Traverse through all node
	for len(queue) != 0 {
		// Pop
		node := queue[0]
		queue = queue[1:]

		// Find neighbor
		for neighbor := range graph[node] {
			// If we found a neighbor and it has not been visited yet
			if graph[node][neighbor] == 1 && visited[neighbor] == 0 {
				// Mark it as visited
				visited[neighbor] = 1
				// Append to the queue to find its neighbor on the next loop
				queue = append(queue, int32(neighbor))
				// Calculate distance: distance = last visited node's distance + 6
				distances[neighbor] = distances[node] + 6
			}
		}
	}

	result := make([]int32, 0)
	for index, val := range distances {
		// Mark unreachable node by -1
		if val == 0 {
			distances[index] = -1
		}
		// Remove starting node distance by not appending it to result slice
		if int32(index) != s {
			result = append(result, distances[index])
		}
	}

	// Return the slide without first element
	// Because our node's number starting from 1 (not 0)
	return result[1:]
}

func main() {
	fmt.Println(bfs(5, 3, [][]int32{
		{1, 2}, {1, 3}, {3, 5},
	}, 2))
}

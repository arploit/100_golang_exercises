package main

import "fmt"

func generateAdjacencylist(edges [][]string) map[string][]string {
	var graph = map[string][]string{}

	for idx := range edges {

		for i := range edges[idx] {
			_, ok := graph[edges[idx][i]]
			if !ok {
				graph[edges[idx][i]] = []string{}
			}
		}

	}

	return graph
}

func main() {

	edges := [][]string{{"i", "j"}, {"k", "i"}, {"m", "k"}, {"k", "l"}, {"o", "n"}}

	graph := generateAdjacencylist(edges)

	fmt.Print(graph)

}

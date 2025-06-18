package main

import (
	"fmt"
	paths2 "github.com/joakimbagge/connected-paths/paths"
)

func main() {
	paths := paths2.ConnectedPaths(1, 2)
	for _, path := range paths {
		fmt.Println(path)
	}
	fmt.Println("Antal vägar", len(paths))
}

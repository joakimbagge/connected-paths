package main

import (
	paths2 "connectedpaths/paths"
	"fmt"
)

func main() {
	paths := paths2.ConnectedPaths(1, 2)
	for _, path := range paths {
		fmt.Println(path)
	}
	fmt.Println("Antal vägar", len(paths))
}

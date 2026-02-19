package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run . <maze-file>")
	}

	filename := os.Args[1]

	grid, err := ParseMaze(filename)
	if err != nil {
		log.Fatal(err)
	}

	start, end, err := FindStartAndEnd(grid)
	if err != nil {
		log.Fatal(err)
	}

	path, err := SolveMaze(grid, start, end)
	if err != nil {
		log.Fatal(err)
	}

	// Output with 1-Based Indexing
	for i, p := range path {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("(%d, %d) ", p.R+1, p.C+1)
	}
	fmt.Println()
}

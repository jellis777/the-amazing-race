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

	fmt.Println("Rows:", len(grid))
	fmt.Println("Cols:", len(grid[0]))

	for _, row := range grid {
		fmt.Println(row)
	}

	start, end, err := FindStartAndEnd(grid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start:", start)
	fmt.Println("End:", end)
}

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseMaze(filename string) ([][]int, error) {
	var grid [][]int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // moves line by line
		line := scanner.Text() // returns the full row of the file as a string

		parts := strings.Fields(line) // makes a slice of strings -- splits on any whitespace

		var row []int

		for _, part := range parts {
			num, err := strconv.Atoi(part) // converts each string "0" or "1" to integers
			if err != nil {
				return nil, err
			}
			row = append(row, num) // puts integers in the row
		}

		grid = append(grid, row) // puts each row into the grid

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

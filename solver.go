package main

import "errors"

type Point struct {
	R int
	C int
}

func FindStartAndEnd(grid [][]int) ([2]int, [2]int, error) {

	if len(grid) == 0 {
		return [2]int{}, [2]int{}, errors.New("grid is empty")
	}

	firstRow := grid[0]

	var start [2]int
	countStart := 0

	for col, val := range firstRow {
		if val == 1 {
			start = [2]int{0, col}
			countStart++
		}
	}

	if countStart != 1 {
		return [2]int{}, [2]int{}, errors.New("invalid number of entrances in first row")
	}

	lastRowIndex := len(grid) - 1
	lastRow := grid[lastRowIndex]

	var end [2]int
	countEnd := 0

	for col, val := range lastRow {
		if val == 1 {
			end = [2]int{lastRowIndex, col}
			countEnd++
		}
	}

	if countEnd != 1 {
		return [2]int{}, [2]int{}, errors.New("invalid number of exits in last row")
	}

	return start, end, nil
}

func Neighbors(grid [][]int, p Point) []Point {
	dirs := []Point{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	var result []Point

	for _, d := range dirs {
		nr := p.R + d.R
		nc := p.C + d.C

		if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) { // Verify neighbor is in the grid
			if grid[nr][nc] == 1 { // Check if it is not a wall
				result = append(result, Point{nr, nc})
			}
		}

	}

	return result

}

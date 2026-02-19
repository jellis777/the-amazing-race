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

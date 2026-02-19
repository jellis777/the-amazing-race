package main

import "errors"

type Point struct {
	R int
	C int
}

func FindStartAndEnd(grid [][]int) (Point, Point, error) {

	if len(grid) == 0 {
		return Point{}, Point{}, errors.New("grid is empty")
	}

	firstRow := grid[0]

	var start Point
	countStart := 0

	for col, val := range firstRow {
		if val == 1 {
			start = Point{0, col}
			countStart++
		}
	}

	if countStart != 1 {
		return Point{}, Point{}, errors.New("invalid number of entrances in first row")
	}

	lastRowIndex := len(grid) - 1
	lastRow := grid[lastRowIndex]

	var end Point
	countEnd := 0

	for col, val := range lastRow {
		if val == 1 {
			end = Point{lastRowIndex, col}
			countEnd++
		}
	}

	if countEnd != 1 {
		return Point{}, Point{}, errors.New("invalid number of exits in last row")
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

func SolveMaze(grid [][]int, start Point, end Point) ([]Point, error) {
	queue := []Point{start}
	visited := make(map[Point]bool)
	parent := make(map[Point]Point)

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for _, neighbor := range Neighbors(grid, current) {
			if !visited[neighbor] {
				visited[neighbor] = true //
				parent[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}

	if !visited[end] {
		return nil, errors.New("no path found")
	}

	var path []Point
	current := end

	for current != start {
		path = append(path, current)
		current = parent[current]
	}

	path = append(path, start)

	// reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, nil
}

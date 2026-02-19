package main

import "testing"

func TestFindStartAndEnd(t *testing.T) {
	grid := [][]int{
		{0, 1, 0},
		{1, 1, 0},
		{0, 1, 0},
	}

	start, end, err := FindStartAndEnd(grid)
	if err != nil {
		t.Fatal(err)
	}

	if start != (Point{0, 1}) {
		t.Errorf("expected start (0,1), got %v", start)
	}

	if end != (Point{2, 1}) {
		t.Errorf("expected end (2,1), got %v", end)
	}
}

func TestSolveMaze(t *testing.T) {
	grid := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	start, end, err := FindStartAndEnd(grid)
	if err != nil {
		t.Fatal(err)
	}

	path, err := SolveMaze(grid, start, end)
	if err != nil {
		t.Fatal(err)
	}

	if path[0] != start {
		t.Errorf("path does not start correctly")
	}

	if path[len(path)-1] != end {
		t.Errorf("path does not end correctly")
	}
}

func TestPathMoveOneStep(t *testing.T) {
	grid := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	start, end, _ := FindStartAndEnd(grid)
	path, _ := SolveMaze(grid, start, end)

	for i := 1; i < len(path); i++ {
		dr := path[i].R - path[i-1].R
		dc := path[i].C - path[i-1].C

		if abs(dr)+abs(dc) != 1 {
			t.Errorf("invalid move from %v to %v", path[i-1], path[i])
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

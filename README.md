# The Amazing Race (Go)

## Overview

This program solves a maze represented as a grid of 0s and 1s.

- 1 = walkable path
- 0 = wall
- The entrance is the only 1 in the first row.
- The exit is the only 1 in the last row.
- Movement is allowed up, down, left, and right. 

This program models the maze as a graph and uses Breadth-First Search (BFS) to find the path from the entrance to the exit. 

---

## Requirements 

- Go 1.20 or newer

To check your version:

    go version

---

## How to Run 

From the project directory:

    go run . maze1.txt

Or:

    go run . maze2.txt

---

## Example Output 

    Path: (1,4) (2,4) (3,4) (3,5) ...

Output coordinates are printed using 1-based indexing for readability. 

---

## Running Tests

To run tests:

    go test 

--- 

## Details 

- The maze is parsed into a 2D slice (`[][]int`).
- Each walkable cell is treated as a graph node.
- Adjacent walkable cells (up, down, left, right) form edges.
- BFS is used to traverse the graph layer-by-layer.
- A `visited` map prevents cycles and infinite loops.
- A `parent` map stores traversal history to reconstruct the final path.
- The path is reconstructed from exit to entrance and then reversed for correct order.
- Defensive validation ensures exactly one entrance an one exit. 
- Errors are returned instead of panicking to maintain production-grade safety. 

---

# connected-paths
*ConnectedPaths* is a Go package that calculates all possible paths across a grid from left to right column. Each path is represented as a slice, where each integer indicates the row index at each column step.

## Installation

```bash
go get github.com/joakimbagge/connected-paths
```

## Design
The ConnectedPaths function uses a depth-first search (DFS) starting from each row in the first column and continue with either of the three movement options:

* Up (row - 1)
* Down (row + 1)
* Right (row)

The movement is restricted within the boundaries of the grid.

## Example
Let's say that the grid is given the dimension 2x3. The possible routes from left to right would then start from either of the rows in first column to the right:

[0, 0]
[0, 1]
[1, 0]
[1, 1]
[1, 2]
[2, 1]
[2, 2]

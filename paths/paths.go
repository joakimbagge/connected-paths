package paths

// ConnectedPaths returnerar alla möjliga vägar i ett rutnät från vänster till höger.
// Varje väg representeras som en lista som innehåller rad-position för varje kolumn.
func ConnectedPaths(cols, rows int) [][]int {
	var result [][]int

	var dfs func(col, row int, path []int)
	dfs = func(col, row int, path []int) {

		// För sista kolumnen
		if col == cols {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		for _, d := range []int{-1, 0, 1} {
			newRow := row + d
			// Kontrollera att den nya raden är inom gränserna
			if newRow >= 0 && newRow < rows {
				dfs(col+1, newRow, append(path, newRow))
			}
		}
	}

	// Påbörja från varje rad i första kolumnen
	for r := 0; r < rows; r++ {
		dfs(1, r, []int{r})
	}

	return result
}

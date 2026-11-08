/*
	{
	    "problem_name": "Count Servers that Communicate",
	    "category": "Arrays / Matrix / Counting",
	    "time_complexity": "O(M * N) -> Two ultra-fast linear passes through the grid.",
	    "space_complexity": "O(M + N) -> Two flat 1D frequency arrays. Absolute zero dynamic allocations inside the loop."
	}
*/
package main

func countServers(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	rowCount := make([]int, m)
	colCount := make([]int, n)

	// Range loop bypasses bounds checking in Go
	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	ans := 0
	for i, row := range grid {
		// Micro-optimization: skip entire row if it has no servers
		if rowCount[i] == 0 {
			continue
		}
		for j, val := range row {
			if val == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
				ans++
			}
		}
	}

	return ans
}

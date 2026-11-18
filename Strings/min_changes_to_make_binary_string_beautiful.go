/*
{
    "problem_name": "Minimum Number of Changes to Make Binary String Beautiful",
    "category": "Strings / Greedy Optimization",
    "time_complexity": "O(N) -> Pure linear scan jumping 2 bytes at a time.",
    "space_complexity": "O(1) -> Absolute zero heap allocations, pure CPU register logic."
}
*/
package main

func minChanges(s string) int {
	changes := 0
	n := len(s)

	// HFT Trick: Process strictly in pairs. Jump the pointer by 2.
	// We directly index the string which is a highly optimized O(1) byte-level read in Go.
	for i := 0; i < n; i += 2 {
		if s[i] != s[i+1] {
			changes++
		}
	}

	return changes
}
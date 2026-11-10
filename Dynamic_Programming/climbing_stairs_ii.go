/*
	{
	    "problem_name": "Climbing Stairs II",
	    "category": "Dynamic Programming / State Machine",
	    "time_complexity": "O(N) -> Single linear scan calculating 3 transitions per step.",
	    "space_complexity": "O(1) -> Only 3 integer registers used. Absolute zero heap/slice allocations."
	}
*/
package main

func climbStairs(n int, costs []int) int {
	// Base case: If there is only 1 step
	if n == 1 {
		return costs[0] + 1 // Jump 0 -> 1: cost is costs[0] + (1 - 0)^2
	}

	// Initialize the absolute minimum costs for the first 2 steps
	dp0 := 0            // Cost to be at step 0
	dp1 := costs[0] + 1 // Cost to reach step 1

	dp2 := costs[1] + 4       // Way A: Direct jump 0 -> 2
	if dp1+costs[1]+1 < dp2 { // Way B: Jump 0 -> 1 -> 2
		dp2 = dp1 + costs[1] + 1
	}

	// Execute rolling DP state machine for steps 3 to n
	for i := 3; i <= n; i++ {
		cost := costs[i-1] // Cost of the current target step

		// Transitions from i-1, i-2, and i-3
		minCost := dp2 + cost + 1 // (i - (i-1))^2 = 1

		if dp1+cost+4 < minCost { // (i - (i-2))^2 = 4
			minCost = dp1 + cost + 4
		}

		if dp0+cost+9 < minCost { // (i - (i-3))^2 = 9
			minCost = dp0 + cost + 9
		}

		// Shift the state pointers forward by 1 step
		dp0 = dp1
		dp1 = dp2
		dp2 = minCost
	}

	return dp2
}

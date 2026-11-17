/*
	{
	    "problem_name": "Minimum Changes to Make K Semi-palindromes",
	    "category": "Dynamic Programming / Substring Partitioning",
	    "time_complexity": "O(N^3) -> O(N log N) divisor sieve + O(N^3) palindrome match + O(N^2 * K) partition DP.",
	    "space_complexity": "O(N^2) -> Heap allocation for DP matrices (acceptable for dynamic scaling)."
	}
*/
package main

func minimumChanges(s string, k int) int {
	n := len(s)

	// --- STEP 1: Precompute proper divisors for every string length ---
	// Using a Sieve-like approach to find divisors in O(N log N) time instead of O(N^2).
	// divisors[L] will hold all proper divisors for a substring of length L.
	divisors := make([][]int, n+1)
	for d := 1; d < n; d++ {
		for length := 2 * d; length <= n; length += d {
			divisors[length] = append(divisors[length], d)
		}
	}

	// --- STEP 2: Precompute minimum semi-palindrome conversion costs ---
	// cost[i][j] stores the min changes to make substring s[i:j] a semi-palindrome.
	cost := make([][]int, n+1)
	for i := range cost {
		cost[i] = make([]int, n+1)
		for j := range cost[i] {
			cost[i][j] = n // Initialize with max possible changes (string length)
		}
	}

	for start := 0; start < n; start++ {
		// Minimum length of a semi-palindrome is 2, so end starts at start+2
		for end := start + 2; end <= n; end++ {
			// Iterate over all valid divisors for the current substring length
			for _, d := range divisors[end-start] {
				changes := 0

				// Check palindrome property for each modulo group
				for offset := 0; offset < d; offset++ {
					left := start + offset
					right := end - d + offset

					for left < right {
						if s[left] != s[right] {
							changes++
						}
						// Jump by the divisor 'd' to stay within the same group
						left += d
						right -= d
					}
				}

				// Record the minimum cost found across all valid divisors
				if changes < cost[start][end] {
					cost[start][end] = changes
				}
			}
		}
	}

	// --- STEP 3: Initialize partition DP ---
	// dp[p][end] = min changes to partition the first `end` characters into `p` semi-palindromes.
	const INF = 1000000000
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0 // Base case: 0 cost to partition 0 characters into 0 semi-palindromes

	// --- STEP 4: Execute DP State Transitions ---
	for p := 1; p <= k; p++ {
		// Minimum characters needed for `p` semi-palindromes is 2*p
		for end := 2 * p; end <= n; end++ {
			// Try all valid split points for the last segment
			// The previous p-1 partitions need at least 2*(p-1) characters.
			// The current segment needs at least 2 characters, so start <= end-2.
			for start := 2 * (p - 1); start <= end-2; start++ {
				candidate := dp[p-1][start] + cost[start][end]
				if candidate < dp[p][end] {
					dp[p][end] = candidate
				}
			}
		}
	}

	// --- STEP 5: Return Result ---
	// Return the minimum changes to partition the entire string into k semi-palindromes.
	return dp[k][n]
}

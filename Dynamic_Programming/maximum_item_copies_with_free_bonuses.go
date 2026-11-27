/*
{
    "problem_name": "Maximum Item Copies With Free Bonuses",
    "category": "Dynamic Programming / 0-1 Knapsack & Harmonic Sieve",
    "time_complexity": "O(M log M + N * B) -> Harmonic sieve for divisor counts + bounded 0/1 knapsack.",
    "space_complexity": "O(M + B) -> Strictly stack-allocated fixed arrays ([1501]int). Zero heap allocations."
}
*/
package main

func maximumSaleItems(items [][]int, budget int) int {
	// HFT Trick 1: Pure stack-allocated arrays. Zero GC overhead.
	var freq [1501]int
	var freeCount [1501]int

	maxFactor := 0
	minPrice := 1501

	// Pass 1: Record factor frequencies and track global minPrice
	for _, it := range items {
		f, p := it[0], it[1]
		freq[f]++
		if f > maxFactor {
			maxFactor = f
		}
		if p < minPrice {
			minPrice = p
		}
	}

	// Pass 2: Harmonic Sieve O(M log M) to count free copies for each factor
	for f := 1; f <= maxFactor; f++ {
		if freq[f] == 0 {
			continue
		}
		total := 0
		for mult := f; mult <= maxFactor; mult += f {
			total += freq[mult]
		}
		// Subtract 1 because j != i (an item doesn't give a free copy of itself)
		freeCount[f] = total - 1
	}

	// Pass 3: 0/1 Knapsack for buying at most 1 copy of each item type
	// dp[c] = max total copies obtainable by spending at most c on first copies
	var dp [1501]int
	curLimit := 0

	for _, it := range items {
		p := it[1]
		if p > budget {
			continue
		}
		val := 1 + freeCount[it[0]]

		// Tight upper bound avoids scanning untouched DP states
		curLimit += p
		if curLimit > budget {
			curLimit = budget
		}

		for c := curLimit; c >= p; c-- {
			if cand := dp[c-p] + val; cand > dp[c] {
				dp[c] = cand
			}
		}
	}

	// Pass 4: Spend any remaining budget greedily on the cheapest item (+1 copy each)
	ans := 0
	for c := 0; c <= curLimit; c++ {
		total := dp[c] + (budget-c)/minPrice
		if total > ans {
			ans = total
		}
	}

	return ans
}

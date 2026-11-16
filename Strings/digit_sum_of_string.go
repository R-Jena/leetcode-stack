/*
	{
	    "problem_name": "Calculate Digit Sum of a String",
	    "category": "Strings / Byte Manipulation",
	    "time_complexity": "O(N log N) -> At each round, the string size shrinks. Max rounds is logarithmic.",
	    "space_complexity": "O(N) -> Reusing a single byte slice buffer for the next round. Minimal GC overhead."
	}
*/
package main

func digitSum(s string, k int) string {
	// Convert string to byte slice once to avoid immutable string allocation overheads
	current := []byte(s)

	for len(current) > k {
		var nextRound []byte
		n := len(current)

		// Process groups of size k
		for i := 0; i < n; i += k {
			sum := 0
			end := i + k
			if end > n {
				end = n
			}

			// Calculate sum of digits in the current group
			for j := i; j < end; j++ {
				sum += int(current[j] - '0')
			}

			// Fast int to byte slice conversion without using fmt.Sprintf
			if sum == 0 {
				nextRound = append(nextRound, '0')
			} else {
				// Extract digits in reverse order
				var temp []byte
				for sum > 0 {
					temp = append(temp, byte(sum%10)+'0')
					sum /= 10
				}
				// Append to nextRound in correct order
				for m := len(temp) - 1; m >= 0; m-- {
					nextRound = append(nextRound, temp[m])
				}
			}
		}
		// Update current for the next round
		current = nextRound
	}

	return string(current)
}

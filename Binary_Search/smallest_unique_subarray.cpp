/*
{
    "problem_name": "Smallest Unique Subarray",
    "category": "Binary Search / Double Rolling Hash",
    "time_complexity": "O(N log^2 N) -> Precomputed O(N) prefix hashes + Binary Search with flat uint64_t sort.",
    "space_complexity": "O(N) -> Pre-allocated flat arrays and single reused hash buffer. Zero unordered_map overhead."
}
*/
#include <vector>
#include <algorithm>
#include <cstdint>

using namespace std;

class Solution {
public:
    int smallestUniqueSubarray(vector<int>& nums) {
        int n = nums.size();
        if (n <= 1) return n;

        // HFT Fast-Path: Check L = 1 in O(N) using a flat frequency array
        int maxVal = *max_element(nums.begin(), nums.end());
        vector<int> freq(maxVal + 1, 0);
        for (int x : nums) freq[x]++;
        for (int x : nums) {
            if (freq[x] == 1) return 1;
        }

        // Precompute Double Rolling Hash (2 independent 32-bit mods packed into uint64_t)
        constexpr uint64_t MOD1 = 1000000007ULL;
        constexpr uint64_t MOD2 = 1000000009ULL;
        constexpr uint64_t BASE1 = 100003ULL;
        constexpr uint64_t BASE2 = 100019ULL;

        vector<uint64_t> h1(n + 1, 0), h2(n + 1, 0);
        vector<uint64_t> p1(n + 1, 1), p2(n + 1, 1);

        for (int i = 0; i < n; ++i) {
            h1[i + 1] = (h1[i] * BASE1 + nums[i]) % MOD1;
            h2[i + 1] = (h2[i] * BASE2 + nums[i]) % MOD2;
            p1[i + 1] = (p1[i] * BASE1) % MOD1;
            p2[i + 1] = (p2[i] * BASE2) % MOD2;
        }

        // Pre-allocate hash buffer once to avoid repeated heap allocations inside binary search
        vector<uint64_t> hashes(n);

        auto hasUnique = [&](int L) -> bool {
            int m = n - L + 1;
            for (int i = 0; i < m; ++i) {
                uint64_t x1 = (h1[i + L] + MOD1 - (h1[i] * p1[L]) % MOD1) % MOD1;
                uint64_t x2 = (h2[i + L] + MOD2 - (h2[i] * p2[L]) % MOD2) % MOD2;
                hashes[i] = (x1 << 32) | x2;
            }

            sort(hashes.begin(), hashes.begin() + m);

            // Check if any hash appears strictly once
            for (int i = 0; i < m; ++i) {
                bool diffLeft = (i == 0 || hashes[i] != hashes[i - 1]);
                bool diffRight = (i == m - 1 || hashes[i] != hashes[i + 1]);
                if (diffLeft && diffRight) return true;
            }
            return false;
        };

        // Binary search the minimum valid subarray length in [2, n]
        int lo = 2, hi = n;
        while (lo < hi) {
            int mid = lo + ((hi - lo) >> 1);
            if (hasUnique(mid)) {
                hi = mid;
            } else {
                lo = mid + 1;
            }
        }

        return lo;
    }
};

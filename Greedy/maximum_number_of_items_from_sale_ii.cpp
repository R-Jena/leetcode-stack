/*
{
    "problem_name": "Maximum Number of Items From Sale II",
    "category": "Greedy / Ordered Map & Multiple Counting",
    "time_complexity": "O(N log N) -> Multiple counting bounded by maxFact + ordered map aggregation by price.",
    "space_complexity": "O(N) -> Frequency array sized to maxFact and ordered map for valid bonus deals."
}
*/
#include <vector>
#include <map>
#include <algorithm>
#include <climits>

using namespace std;

class Solution {
public:
    int maximumSaleItems(vector<vector<int>>& items, int budget) {
        int n = items.size();
        int maxFact = 0, minPrice = INT_MAX;

        // Pass 1: Track maximum factor (for frequency array sizing) and global minimum price
        for (auto& item : items) {
            maxFact = max(maxFact, item[0]);
            minPrice = min(minPrice, item[1]);
        }

        // Build frequency table of all factors
        vector<int> freq(maxFact + 1, 0);
        for (auto& item : items) {
            freq[item[0]]++;
        }

        // Ordered map: price -> total bonus copies available at this price
        // Automatically keeps deals sorted by ascending price for greedy purchasing
        map<int, long long> mp;
        for (auto& item : items) {
            int price = item[1];
            int f = item[0];

            // Skip items where buying 1 bonus copy (2 items) costs more than 2 cheapest regular copies
            if (price > 2 * minPrice) continue;

            // Count how many items have a factor divisible by f
            long long cnt = 0;
            for (int i = f; i <= maxFact; i += f) {
                cnt += freq[i];
            }

            // Exclude the item itself (since i != j)
            cnt--;
            if (cnt > 0) {
                mp[price] += cnt;
            }
        }

        // Pass 2: Greedily buy the cheapest bonus copies first (each yields 1 purchased + 1 free = 2 copies)
        long long bonus = 0;
        long long cur = budget;
        for (auto& [price, cnt] : mp) {
            if (cur < price) break;
            long long buy = min(cnt, 1LL * cur / price);
            cur -= buy * price;
            bonus += buy;
        }

        // Pass 3: Spend remaining budget on the cheapest regular item (+1 copy each)
        long long ans = cur / minPrice;
        ans += 2 * bonus;
        return ans;
    }
};

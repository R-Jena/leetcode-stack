/*
{
    "problem_name": "Find Subtree Sizes After Changes",
    "category": "Trees / Single-Pass Post-Order DFS",
    "time_complexity": "O(N) -> Single DFS pass with O(1) ancestor lookup via 26-element stack array.",
    "space_complexity": "O(N) -> Flat forward-star arrays (head/nxt) bypassing vector<vector<int>> heap fragmentation."
}
*/
#include <vector>
#include <string>
#include <cstring>

using namespace std;

class Solution {
public:
    vector<int> findSubtreeSizes(vector<int>& parent, string s) {
        int n = parent.size();

        // HFT Trick 1: Flat Forward-Star representation instead of vector<vector<int>>
        // Zero per-node heap allocations, maximum L1/L2 cache locality.
        vector<int> head(n, -1);
        vector<int> nxt(n, -1);
        for (int i = 1; i < n; ++i) {
            nxt[i] = head[parent[i]];
            head[parent[i]] = i;
        }

        vector<int> ans(n, 1); // Every node starts with subtree size 1 (itself)
        int last[26];
        memset(last, -1, sizeof(last)); // Stack-allocated closest ancestor tracker

        // HFT Trick 2: Single-Pass Post-Order DFS
        // Since any re-attached node moves strictly to an original ancestor,
        // ans[u] is 100% finalized once all original children of u return.
        auto dfs = [&](auto& self, int u) -> void {
            int c = s[u] - 'a';
            int prev = last[c];
            last[c] = u;

            for (int v = head[u]; v != -1; v = nxt[v]) {
                self(self, v);
            }

            // Restore ancestor state before bubbling up
            last[c] = prev;

            // Directly push finalized subtree size to the new parent
            if (u != 0) {
                int newParent = (prev != -1) ? prev : parent[u];
                ans[newParent] += ans[u];
            }
        };

        dfs(dfs, 0);
        return ans;
    }
};

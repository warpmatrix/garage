#include <algorithm>
#include <cstddef>
#include <queue>
#include <unordered_set>
#include <vector>

class Solution {
   public:
    std::vector<std::vector<int>> getAncestors(
        int n, std::vector<std::vector<int>>& edges) {
        std::vector<std::vector<int>> e(n);
        std::vector<int> indeg(n, 0);
        for (auto&& edge : edges) {
            e[edge[0]].push_back(edge[1]);
            indeg[edge[1]]++;
        }
        std::queue<int> q;
        for (size_t node = 0; node < e.size(); node++) {
            if (indeg[node] == 0) {
                q.push(node);
            }
        }
        std::vector<std::unordered_set<int>> ancs(n);
        while (q.size()) {
            int node = q.front();
            q.pop();
            for (auto&& adjNode : e[node]) {
                // can be optimized to O(1) using bitset
                ancs[adjNode].insert(ancs[node].cbegin(), ancs[node].cend());
                ancs[adjNode].insert(node);
                indeg[adjNode]--;
                if (indeg[adjNode] == 0) {
                    q.push(adjNode);
                }
            }
        }
        std::vector<std::vector<int>> ret(n);
        for (size_t i = 0; i < n; i++) {
            for (size_t j = 0; j < n; j++) {
                if (ancs[i].find(j) != ancs[i].end()) {
                    ret[i].push_back(j);
                }
            }
        }
        return ret;
    }
};

int main(int argc, char const* argv[]) {
    return 0;
}

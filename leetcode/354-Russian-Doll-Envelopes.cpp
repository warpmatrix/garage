#include <algorithm>
#include <iostream>
#include <vector>
using std::vector;

bool cmp(const vector<int> &lhs, const vector<int> &rhs) {
    return lhs[0] < rhs[0] || lhs[0] == rhs[0] && lhs[1] > rhs[1];
}

int maxEnvelopes(vector<vector<int>> &envelopes) {
    if (envelopes.empty()) return 0;
    std::sort(envelopes.begin(), envelopes.end(), cmp);
    std::vector<int> incSeqMin(1, envelopes[0][1]);
    for (int i = 1; i < envelopes.size(); i++) {
        std::vector<int>::iterator lowBnd =
            std::lower_bound(incSeqMin.begin(), incSeqMin.end(), envelopes[i][1]);
        if (lowBnd == incSeqMin.end()) {
            incSeqMin.push_back(envelopes[i][1]);
        } else {
            *lowBnd = envelopes[i][1];
        }
    }
    return incSeqMin.size();
}

// [[5,4],[6,4],[6,7],[2,3]]
// [[4,5],[4,6],[6,7],[2,3],[1,1]]

#include <algorithm>
#include <iostream>
#include <queue>
#include <vector>

using std::vector;

int findMaximizedCapital(int k, int w, vector<int>& profits, vector<int>& capital) {
    typedef std::pair<int, int> Proj;
    std::vector<Proj> projs(profits.size());
    for (int i = 0; i < projs.size(); i++) {
        projs[i] = {capital[i], profits[i]};
    }
    std::sort(projs.begin(), projs.end());
    std::priority_queue<int> pq;
    for (int i = 0, idx = 0; i < k; i++) {
        while (idx < projs.size() && projs[idx].first <= w) {
            pq.push(projs[idx].second);
            idx++;
        }
        if (pq.empty()) break;
        w += pq.top();
        pq.pop();
    }
    return w;
}

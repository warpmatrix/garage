#include <cmath>
#include <unordered_map>
#include <vector>

using std::vector;

int calcDist(int x0, int x1, int y0, int y1) {
    return (x0 - x1) * (x0 - x1) + (y0 - y1) * (y0 - y1);
}

int numberOfBoomerangs(vector<vector<int>> &points) {
    int ret = 0;
    for (auto &&origin : points) {
        std::unordered_map<int, int> cnts;
        for (auto &&point : points) {
            int dist = calcDist(origin[0], point[0], origin[1], point[1]);
            cnts[dist]++;
        }
        for (auto &&cnt : cnts) {
            ret += cnt.second * (cnt.second - 1);
        }
    }
    return ret;
}

#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

bool isPossible(vector<int>& nums) {
    unordered_map<int, int> cnts, endArrCnts;
    for (auto&& num : nums) {
        cnts[num]++;
    }
    for (auto&& num : nums) {
        if (cnts[num] == 0) continue;
        if (!endArrCnts[num - 1]) {
            if (!cnts[num] || !cnts[num + 1] || !cnts[num + 2]) return false;
            cnts[num + 1]--, cnts[num + 2]--;
            endArrCnts[num + 2]++;
        } else {
            endArrCnts[num - 1]--;
            endArrCnts[num]++;
        }
        cnts[num]--;
    }
    return true;
}

int main(int argc, char const* argv[]) {
    std::vector<int> v;
    v = std::vector<int>({1, 2, 3, 3, 4, 5});
    std::cout << isPossible(v) << '\n';
    v = std::vector<int>({1, 2, 3, 3, 4, 4, 5, 5});
    std::cout << isPossible(v) << '\n';
    v = std::vector<int>({1, 2, 3, 4, 4, 5});
    std::cout << isPossible(v) << '\n';
    return 0;
}

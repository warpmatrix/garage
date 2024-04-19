#include <cstddef>
#include <iostream>
#include <map>
#include <vector>

class Solution {
   public:
    std::vector<int> findOriginalArray(std::vector<int>& changed) {
        std::vector<int> ret;
        if (changed.size() % 2 != 0) {
            return ret;
        }
        std::map<int, int> cnts;
        for (auto&& num : changed) {
            cnts[num]++;
        }
        while (cnts.size() > 0) {
            int key = cnts.begin()->first;
            int target = key * 2;
            if (key == target) {
                if (cnts[key] % 2 != 0) {
                    break;
                }
                cnts[key] /= 2;
            }
            if (cnts[key] > cnts[target]) {
                break;
            }
            cnts[target] -= cnts[key];
            for (size_t i = 0; i < cnts[key]; i++) {
                ret.push_back(key);
            }
            if (cnts[target] == 0) {
                cnts.erase(target);
            } else {
                cnts.erase(key);
            }
        }
        if (cnts.size() > 0) {
            return std::vector<int>();
        }
        return ret;
    }
};

int main(int argc, char const* argv[]) {
    auto s = Solution();
    // std::vector<int> input({1, 3, 4, 2, 6, 8});
    // auto res = s.findOriginalArray(input);
    // for (auto&& num : res) {
    //     std::cout << num << '\n';
    // }
    std::vector<int> input({0, 0, 0, 0});
    auto res = s.findOriginalArray(input);
    for (auto&& num : res) {
        std::cout << num << '\n';
    }
    return 0;
}

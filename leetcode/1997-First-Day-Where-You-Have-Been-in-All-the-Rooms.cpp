#include <iostream>
#include <vector>
using namespace std;

class Solution {
   public:
    int firstDayBeenInAllRooms(vector<int>& nextVisit) {
        const int MOD = (int)1e9 + 7;
        vector<int> day{0};
        for (size_t i = 1; i < nextVisit.size(); i++) {
            int revisit = (1 + day[i - 1] - day[nextVisit[i - 1]] + MOD) % MOD;
            day.push_back((day[i - 1] + revisit + 1) % MOD);
        }
        return day.back();
    }
};

int main(int argc, char const* argv[]) {
    auto s = Solution();
    auto input = vector({0, 0, 2, 1});
    std::cout << s.firstDayBeenInAllRooms(input) << '\n';
    return 0;
}

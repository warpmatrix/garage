#include <unordered_map>
#include <vector>
using namespace std;

int fourSumCount(vector<int>& A, vector<int>& B, vector<int>& C,
                 vector<int>& D) {
    std::unordered_map<int, int> abSum;
    int res = 0;
    for (auto&& a : A) {
        for (auto&& b : B) {
            abSum[a + b]++;
        }
    }
    for (auto&& c : C) {
        for (auto&& d : D) {
            res += abSum[-(c + d)];
        }
    }
    return res;
}

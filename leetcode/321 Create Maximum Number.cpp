#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

std::vector<int> findMaxSubSeq(std::vector<int>& nums, int k) {
    int left = nums.size() - k;
    std::vector<int> res;
    for (int i = 0; i < nums.size(); i++) {
        while (res.size() && nums[i] > res.back() && left) {
            res.pop_back();
            left--;
        }
        res.push_back(nums[i]);
    }
    res.resize(k);
    return res;
}

vector<int> maxNumber(vector<int>& nums1, vector<int>& nums2, int k) {
    int m = nums1.size(), n = nums2.size();
    std::vector<int> res;
    for (int i = std::max(0, k - n); i <= std::min(k, m); i++) {
        std::vector<int> seq1 = findMaxSubSeq(nums1, i);
        std::vector<int> seq2 = findMaxSubSeq(nums2, k - i);
        // lexicographical_compare
        auto it1 = seq1.begin(), it2 = seq2.begin();
        std::vector<int> tmp;
        while (it1 != seq1.end() || it2 != seq2.end()) {
            if (std::lexicographical_compare(it1, seq1.end(), it2, seq2.end())) {
                tmp.push_back(*it2++);
            } else {
                tmp.push_back(*it1++);
            }
        }
        if (std::lexicographical_compare(res.begin(), res.end(), tmp.begin(), tmp.end())) res = tmp;
    }
    return res;
}

int main(int argc, char const *argv[])
{
    vector<int> num1({6, 7, 5});
    vector<int> num2({4, 8, 1});
    int k = 3;
    std::cout << maxNumber(num1, num2, k).size() << '\n';
    return 0;
}


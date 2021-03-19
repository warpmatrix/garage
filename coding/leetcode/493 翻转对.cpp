#include <iostream>
#include <vector>
using namespace std;

int getImptRevPair(int l, int r, const vector<int>& nums) {
    int ret = 0;
    int mid = (l + r) / 2, p2 = mid;
    for (int p1 = l; p1 < mid; p1++) {
        while (p2 < r && (long long) nums[p1] > (long long) nums[p2] * 2) p2++;
        ret += p2 - mid;
    }
    return ret;
}

int mergeSort(int l, int r, vector<int>& nums) {
    if (l + 1 == r || l == r) return 0;
    int mid = (l + r) / 2;
    int res = 0;
    res += mergeSort(l, mid, nums);
    res += mergeSort(mid, r, nums);
    res += getImptRevPair(l, r, nums);
    int tmp[r - l];
    int p1 = l, p2 = mid, pTmp = 0;
    while (p1 < mid || p2 < r) {
        if (p1 < mid && p2 < r) {
            if (nums[p1] <= nums[p2]) {
                tmp[pTmp++] = nums[p1++];
            } else {
                tmp[pTmp++] = nums[p2++];
            }
        } else {
            while (p1 < mid) tmp[pTmp++] = nums[p1++];
            while (p2 < r) tmp[pTmp++] = nums[p2++];
        }
    }
    for (int i = l; i < r; i++) {
        nums[i] = tmp[i - l];
    }
    return res;
}

int reversePairs(vector<int>& nums) { return mergeSort(0, nums.size(), nums); }

int main(int argc, char const* argv[]) {
    vector<int> nums({-5, -5});
    std::cout << reversePairs(nums) << '\n';
    return 0;
}

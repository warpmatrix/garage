#include <iostream>
#include <vector>
using namespace std;

int searchHead(std::vector<int>& nums, int target) {
    if (nums[0] == target) return 0;
    int l = 0, r = nums.size();
    while (l + 1 < r) {
        int mid = (l + r) / 2;
        if (nums[mid] >= target) {
            r = mid;
        } else {
            l = mid;
        }
    }
    if (r == nums.size() || nums[r] != target)
        return -1;
    else
        return r;
}

int searchTail(vector<int>& nums, int target) {
    int l = 0, r = nums.size();
    while (l + 1 < r) {
        int mid = (l + r) / 2;
        if (nums[mid] > target) {
            r = mid;
        } else {
            l = mid;
        }
    }
    return l;
}

vector<int> searchRange(vector<int>& nums, int target) {
    std::vector<int> ret(2, -1);
    if (nums.size() == 0) return ret;
    ret[0] = searchHead(nums, target);
    if (ret[0] == -1) return ret;
    ret[1] = searchTail(nums, target);
    return ret;
}

int main(int argc, char const* argv[]) {
    std::vector<int> v;
    std::vector<int> res;
    v = std::vector<int>({5, 7, 7, 8, 8, 10});
    res = searchRange(v, 8);
    std::cout << res[0] << ' ' << res[1] << '\n';
    v = std::vector<int>({5, 7, 7, 8, 8, 10});
    res = searchRange(v, 6);
    std::cout << res[0] << ' ' << res[1] << '\n';
    v = std::vector<int>({});
    res = searchRange(v, 0);
    std::cout << res[0] << ' ' << res[1] << '\n';
    v = std::vector<int>({1});
    res = searchRange(v, 0);
    std::cout << res[0] << ' ' << res[1] << '\n';
    v = std::vector<int>({1});
    res = searchRange(v, 1);
    std::cout << res[0] << ' ' << res[1] << '\n';
    v = std::vector<int>({2, 2});
    res = searchRange(v, 3);
    std::cout << res[0] << ' ' << res[1] << '\n';
    return 0;
}

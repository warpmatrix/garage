#include <algorithm>
#include <functional>
#include <vector>
using namespace std;

int largestPerimeter(vector<int>& A) {
    std::sort(A.begin(), A.end(), std::greater<int>());
    for (size_t minIdx = 2; minIdx < A.size(); minIdx++) {
        if (A[minIdx] + A[minIdx - 1] > A[minIdx - 2]) {
            return A[minIdx] + A[minIdx - 1] + A[minIdx - 2];
        }
    }
    return 0;
}

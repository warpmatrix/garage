#include <vector>
using std::vector;

vector<int> spiralOrder(vector<vector<int>>& matrix) {
    std::vector<int> ret;
    ret.resize(matrix.size() * matrix[0].size());
    int index = 0;
    for (int top = 0, bottom = matrix.size() - 1, left = 0,
             right = matrix[0].size() - 1;
         top <= bottom && left <= right;
         top++, bottom--, left++, right--) {
        for (int col = left; col <= right; col++) {
            ret[index] = matrix[top][col];
            index++;
        }
        for (int row = top + 1; row <= bottom; row++) {
            ret[index] = matrix[row][right];
            index++;
        }
        for (int col = right - 1; bottom > top && col >= left; col--) {
            ret[index] = matrix[bottom][col];
            index++;
        }
        for (int row = bottom - 1; right > left && row >= top + 1; row--) {
            ret[index] = matrix[row][left];
            index++;
        }
    }
    return ret;
}
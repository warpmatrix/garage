#include <vector>
using std::vector;

bool searchMatrix(vector<vector<int>>& matrix, int target) {
    if (matrix.size() == 0) return false;
    int row = matrix.size() - 1;
    int col = 0;
    while (row >= 0 && col < matrix[0].size()) {
        if (matrix[row][col] == target) return true;
        if (matrix[row][col] > target)
            row--;
        else if (matrix[row][col] < target)
            col++;
    }
    return false;
}
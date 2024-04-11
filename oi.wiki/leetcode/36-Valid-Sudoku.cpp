#include <vector>
using std::vector;

class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        bool rows[9][9] = {};
        bool cols[9][9] = {};
        bool subboxes[3][3][9] = {};
        for (int r = 0; r < board.size(); r++) {
            for (int c = 0; c < board[0].size(); c++) {
                if (board[r][c] == '.') continue;
                int &&num = board[r][c] - '0' - 1;
                if (rows[r][num] || cols[c][num] || subboxes[r / 3][c / 3][num]) {
                    return false;
                }
                rows[r][num] = cols[c][num] = subboxes[r / 3][c / 3][num] = true;
            }
        }
        return true;
    }
};

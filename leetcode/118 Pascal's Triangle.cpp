#include <vector>
using namespace std;

vector<vector<int>> generate(int numRows) {
    vector<vector<int>> ret;
    for (int row = 0; row < numRows; row++) {
        ret.push_back(vector<int>({1}));
        for (int col = 1; col < row - 1; col++) {
            ret[row].push_back(ret[row - 1][col - 1] + ret[row - 1][col]);
        }
        if (row > 0) ret[row].push_back(1);
    }
    return ret;
}

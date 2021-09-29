#include <memory.h>

#include <vector>
using namespace std;

int matrixScore(vector<vector<int>>& A) {
    int rowNum = A.size(), colNum = A[0].size();
    int ret = rowNum * (1 << (colNum - 1));
    bool filp[rowNum];
    memset(filp, false, sizeof(filp));
    for (size_t i = 0; i < rowNum; i++) {
        if (A[i][0] == '0') filp[i] = true;
    }
    for (size_t col = 0; col < colNum; col++) {
        int cnt = 0;
        for (size_t row = 0; row < rowNum; row++) {
            if (A[row][col] ^ filp[col] == 1) cnt++;
        }
        ret += max(cnt, colNum - cnt) * (1 << (colNum - col));
    }
    return ret;
}
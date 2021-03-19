#include <vector>
using std::vector;

vector<int> countBits(int num) {
    // most siginificant bit
    int msb = 0;
    std::vector<int> ret;
    ret.push_back(0);
    for (int i = 1; i <= num; i++) {
        if (i == 1 << msb) msb++;
        ret.push_back(ret[i - (1 << (msb - 1))] + 1);
    }
    return ret;
}

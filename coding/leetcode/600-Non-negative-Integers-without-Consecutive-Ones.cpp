#include <cmath>

class Solution {
public:
    int findIntegers(int n) {
        if (n == 1) return 2;
        int len = log2(n) + 1;
        int fibo[len];
        fibo[0] = 1, fibo[1] = 2;
        for (int i = 2; i < len; i++) {
            fibo[i] = fibo[i - 1] + fibo[i - 2];
        }
        int ret = 0;
        for (int i = len - 1; i >= 0; i--) {
            if (n & (1 << i)) {
                ret += fibo[i];
                i--;
                if (i >= 0 && n & (1 << i)) {
                    ret += fibo[i];
                    return ret;
                }
            }
        }
        if ((n & 1) == 0 || (n & 2) == 0) {
            ret++;
        }
        return ret;
    }
};

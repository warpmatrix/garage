#include <memory.h>

#include <vector>
using namespace std;

bool lemonadeChange(vector<int>& bills) {
    int cnts[3];
    memset(cnts, 0, sizeof(cnts));
    for (auto&& bill : bills) {
        switch (bill) {
            case 5:
                cnts[0]++;
                break;
            case 10:
                if (cnts[0]) {
                    cnts[0]--;
                    cnts[1]++;
                } else
                    return false;
                break;
            case 20:
                if (cnts[1] && cnts[0]) {
                    cnts[0]--, cnts[1]--;
                    cnts[2]++;
                } else if (cnts[0] >= 3) {
                    cnts[0] -= 3;
                    cnts[2]++;
                } else
                    return false;
            default:
                break;
        }
    }
    return true;
}

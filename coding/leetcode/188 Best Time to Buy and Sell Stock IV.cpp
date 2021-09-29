#include <memory.h>

#include <algorithm>
#include <vector>
using namespace std;

int maxProfit(int k, vector<int>& prices) {
    if (k == 0 || prices.empty()) return 0;
    int n = prices.size();
    k = min(k, n / 2);
    int buy[k + 1], sell[k + 1];
    memset(buy, 0xc0, sizeof(buy));
    memset(sell, 0xc0, sizeof(sell));
    buy[0] = -prices[0], sell[0] = 0;
    for (int i = 1; i < n; i++) {
        buy[0] = max(buy[0], sell[0] - prices[i]);
        for (int j = 1; j <= k; j++) {
            buy[j] = max(buy[j], sell[j] - prices[i]);
            sell[j] = max(buy[j - 1] + prices[i], sell[j]);
        }
    }
    int maxProfit = 0;
    for (int i = 0; i <= k; i++) {
        maxProfit = max(maxProfit, sell[i]);
    }
    return maxProfit;
}

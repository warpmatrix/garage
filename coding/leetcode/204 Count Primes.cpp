#include <memory.h>

#include <iostream>

int countPrimes(int n) {
    if (n <= 2) return 0;
    int cnt = 0;
    bool isPrime[n];
    memset(isPrime, true, sizeof(isPrime));
    for (int i = 2; i < n; i++) {
        if (!isPrime[i]) continue;
        cnt++;
        for (int j = (i << 1); j < n; j += i) {
            isPrime[j] = false;
        }
    }
    return cnt;
}

int main(int argc, char const *argv[]) {
    std::cout << countPrimes(10) << '\n';
    std::cout << countPrimes(0) << '\n';
    std::cout << countPrimes(1) << '\n';
    std::cout << countPrimes(2) << '\n';
    return 0;
}

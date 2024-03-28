#include <iostream>
#include <vector>
#include <cmath>

long long int power10[11] = {
    int(1e0), int(1e1), int(1e2), int(1e3), int(1e4),
    int(1e5), int(1e6), int(1e7), int(1e8), int(1e9), (long long)(1e10),
};

int countNumInDigit(int k, int num, int digit) {
    int highDigit = num / power10[digit + 1];
    int weight = power10[digit];
    int ret = highDigit * weight;
    int digit_num = num / power10[digit] % 10;
    if (digit_num > k) {
        ret += weight;
    } else if (digit_num == k) {
        int lowDigit = num % power10[digit];
        ret += lowDigit + 1;
    }
    return ret;
}

int digitOneInNumber(int num) {
    int totalDigit = 1;
    if (num > 0) {
        totalDigit += (int)log10(num);
    }
    int sum = 0;
    for (size_t digit = 0; digit < totalDigit; digit++) {
        sum += countNumInDigit(1, num, digit);
    }
    return sum;
}

int main(int argc, char const *argv[]) {
    std::cout << digitOneInNumber(13) << '\n';
    std::cout << digitOneInNumber(0) << '\n';
    std::cout << digitOneInNumber(1) << '\n';
    std::cout << digitOneInNumber(123) << '\n';
    std::cout << digitOneInNumber(1'000'000'000) << '\n';
    return 0;
}

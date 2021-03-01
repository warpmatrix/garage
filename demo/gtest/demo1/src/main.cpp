#include <iostream>
#include "factorial.h"

int main(int argc, char const *argv[]) {
    int n;
    std::cin >> n;
    std::cout << factorial(n) << '\n';
    return 0;
}

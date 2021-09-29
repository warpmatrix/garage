#include "factorial.h"
#include <stdexcept>

int factorial(int num) {
    if (num < 0) throw num;
    if (num < 1) return 1;
    return num * factorial(num-1);
}

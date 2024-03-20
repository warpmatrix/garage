#include "m1/f1.hpp"

#include <iostream>

void f1() {
    std::cout << "hello in f1" << '\n';
    f2();
}

void f2() {
    std::cout << "hello in f2" << '\n';
}

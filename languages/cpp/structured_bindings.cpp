#include <iostream>
#include <tuple>

auto f() {
    return std::tuple(1, 2);
}

int main() {
    // auto [a, b] = std::make_tuple(1,2);
    // auto [a, b] = f();
    auto [a, b] = []() -> auto {
        return std::tuple(1, 2);
    }();
    std::cout << a << " " << b << "\n";
    return 0;
}

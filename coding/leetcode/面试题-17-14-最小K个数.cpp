#include <vector>

int main(int argc, char const *argv[]) {
    std::vector<int> a;
	std::vector<int> b(a.begin(), a.end());
    return 0;
}

#include "sparsity-table.hpp"

#include <cmath>

template <typename T>
SparsityTable<T>::SparsityTable(
    const std::vector<T> &nums, op_type f
){
    this->op = f;
    int len = nums.size();
    const int LOG = std::ceil(std::log2(len)) + 1;
    st.assign(len, std::vector<int>(LOG, 0));
    for (size_t i = 0; i < len; i++) {
        st[i][0] = nums[i];
    }
    for (size_t j = 0; j < LOG; j++) {
        int rngLen = 1 << j - 1;
        for (size_t i = 0; i + rngLen < len; i++) {
            st[i][j] = op(st[i][j - 1], st[i + rngLen][j - 1]);
        }
    }
}

template <typename T>
T SparsityTable<T>::query(int l, int r) {
    int len = r - l + 1;
    int s = std::floor(std::log2(len));
    T ret = op(st[l][s], st[r - (1 << s) + 1][s]);
    return ret;
}

#include <cmath>
#include <functional>
#include <vector>

template <typename T>
class SparsityTable {
    std::vector<std::vector<T>> st;
    using op_type = std::function<T(const T &, const T &)>;
    static T default_op(const T &lhs, const T &rhs) { return std::max(lhs, rhs); }
    op_type op;

   public:
    SparsityTable(const std::vector<T> &nums, op_type f = default_op) {
        this->op = f;
        int len = nums.size();
        const int LOG = std::ceil(std::log2(len)) + 1;
        st.assign(len, std::vector<int>(LOG, 0));
        for (size_t i = 0; i < len; i++) {
            st[i][0] = nums[i];
        }
        for (size_t j = 0; j < LOG; j++) {
            int rngLen = 1 << (j - 1);
            for (size_t i = 0; i + rngLen < len; i++) {
                st[i][j] = op(st[i][j - 1], st[i + rngLen][j - 1]);
            }
        }
    }

    T query(int l, int r) {
        int len = r - l + 1;
        int s = std::floor(std::log2(len));
        T ret = op(st[l][s], st[r - (1 << s) + 1][s]);
        return ret;
    }
};

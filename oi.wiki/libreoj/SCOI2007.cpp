#include <algorithm>
#include <cmath>
#include <cstdio>
#include <functional>
#include <map>
#include <string>
#include <tuple>
#include <vector>


template <typename T>
class SparsityTable {
    std::vector<std::vector<T>> st;
    using op_type = std::function<T(const T &, const T &)>;
    static T default_op(const T &lhs, const T &rhs) { return std::max(lhs, rhs); }
    op_type op;
    std::vector<int> LOG2;

   public:
    SparsityTable(const std::vector<T> &nums, op_type f = default_op) : op(f) {
        int len = nums.size();
        LOG2.push_back(-1);
        for (size_t i = 1; i <= len; i++) {
            LOG2.push_back(LOG2[i / 2] + 1);
        }
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
        int s = LOG2[len];
        T ret = op(st[l][s], st[r - (1 << s) + 1][s]);
        return ret;
    }

    T &operator[](int index) {
        return st[index][0];
    }
};


std::string judgeStatement(
    std::vector<int> &ys, SparsityTable<int> &rs, int x, int y
) {
    auto year2Idx = [&ys](int year) {
        auto iter = std::lower_bound(ys.begin(), ys.end(), year);
        bool found = *iter == year;
        auto idx = iter - ys.begin();
        return std::make_pair(found, idx);
    };
    auto [foundY, idxY] = year2Idx(y);
    auto [foundX, idxX] = year2Idx(x);
    if (foundX && foundY && rs[idxX] > rs[idxY]) {
        return "false";
    }
    int z_l = y + 1, z_r = x - 1;
    auto l = std::lower_bound(ys.begin(), ys.end(), z_l) - ys.begin();
    auto r = std::upper_bound(ys.begin(), ys.end(), z_r) - ys.begin() - 1;
    if (l <= r) {
        int r_z = rs.query(l, r);
        if (foundY && rs[idxY] <= r_z) {
            return "false";
        }
        if (foundX && rs[idxX] <= r_z) {
            return "false";
        }
    }
    bool fullZ = (r - l) == (z_r - z_l);
    if (foundX && foundY && fullZ) {
        return "true";
    }
    return "maybe";
}

int main(int argc, char const *argv[]) {
    int n;
    scanf("%d", &n);
    std::vector<int> ys;
    std::vector<int> rs;
    for (size_t i = 0; i < n; i++) {
        int y, r;
        scanf("%d %d", &y, &r);
        ys.push_back(y);
        rs.push_back(r);
    }
    SparsityTable<int> st(rs);
    int m;
    scanf("%d", &m);
    for (size_t i = 0; i < m; i++) {
        int x, y;
        scanf("%d %d", &y, &x);
        std::string res = judgeStatement(ys, st, x, y);
        printf("%s\n", res.c_str());
    }
    return 0;
}

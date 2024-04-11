#include <cmath>
#include <cstdio>
#include <vector>

class RngMax {
   private:
    // sparsity table, st[i][j] = max(nums[i : i + 2^j - 1])
    std::vector<std::vector<int>> st;
   public:
    RngMax(std::vector<int> &nums);
    int query(int l, int r);
};

RngMax::RngMax(std::vector<int> &nums) {
    int len = nums.size();
    const int LOG = std::ceil(std::log2(len)) + 1;
    st.assign(len, std::vector<int>(LOG, 0));
    for (size_t i = 0; i < len; i++) {
        st[i][0] = nums[i];
    }
    for (size_t j = 1; j < LOG; j++) {
        int rngLen = 1 << j - 1;
        for (size_t i = 0; i + rngLen < len; i++) {
            st[i][j] = std::max(st[i][j - 1], st[i + rngLen][j - 1]);
        }
    }
}

int RngMax::query(int l, int r) {
    int len = r - l + 1;
    int s = std::floor(std::log2(len));
    int ret = std::max(st[l][s], st[r - (1 << s) + 1][s]);
    return ret;
}

int main(int argc, char const *argv[]) {
    int n, m;
    scanf("%d %d", &n, &m);
    std::vector<int> a(n);
    for (size_t i = 0; i < n; i++) {
        scanf("%d", &a[i]);
    }
    RngMax st(a);
    for (size_t i = 0; i < m; i++) {
        int l, r;
        scanf("%d %d", &l, &r);
        int res = st.query(l - 1, r - 1);
        printf("%d\n", res);
    }
    return 0;
}

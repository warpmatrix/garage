int fib(int n) {
    if (n == 0) { return 0; }
    const int MOD = 1e9 + 7;
    int f1 = 0, f2 = 1;
    for (int i = 2; i <= n; i++) {
        int f0 = f1;
        f1 = f2;
        f2 = (f1 + f0) % MOD;
    }
    return f2;
}

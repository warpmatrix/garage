// The rand7() API is already defined for you.
// int rand7();
// @return a random integer in the range 1 to 7

int rand10() {
    int oddEven = rand7();
    while (oddEven == 7) {
        oddEven = rand7();
    }
    int base = rand7();
    while (base > 5) {
        base = rand7();
    }
    return base * 2 - (oddEven % 2);
}

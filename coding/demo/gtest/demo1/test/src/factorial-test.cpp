#include <gtest/gtest.h>
#include "factorial.h"

TEST(FactorialTest, HandlesNegetiveInput) {
    EXPECT_ANY_THROW(factorial(-1));
}

TEST(FactorialTest, HandlesZeroInput) {
    EXPECT_EQ(factorial(0), 1);
}

TEST(FactorialTest, HandlesPositiveInput) {
    EXPECT_EQ(factorial(1), 1);
    EXPECT_EQ(factorial(2), 2);
    EXPECT_EQ(factorial(3), 6);
    EXPECT_EQ(factorial(8), 40320);
}

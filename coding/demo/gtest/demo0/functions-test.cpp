#include <gtest/gtest.h>
#include "functions.h"

TEST(AddTest, AddTestCase) {
    ASSERT_EQ(2, add(1, 1));
}

TEST(DivideTest, DivideTestCase) {
    ASSERT_EQ(2, divide(7, 3));
}

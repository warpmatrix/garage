<!-- omit in toc -->
# Leetcode Note

- [131. Palindrome Partitioning](#131-palindrome-partitioning)
- [132. Palindrome Partitioning II](#132-palindrome-partitioning-ii)
- [224 & 227. Basic Calculator](#224--227-basic-calculator)
- [338. Counting Bits](#338-counting-bits)
- [354. Russian Doll Envelopes](#354-russian-doll-envelopes)
- [503. Next Greater Element II](#503-next-greater-element-ii)
- [1047. Remove All Adjacent Duplicates In String](#1047-remove-all-adjacent-duplicates-in-string)

## 131. Palindrome Partitioning

划分方法数（解空间）：$\sum\limits_{i=0}^{n-1} C_{n-1}^i = 2^{n-1}$

存在最坏情况下，解空间中的每一个解都接受，时间复杂度 $O(2^n)$

判断回文方法：动态规划，双指针判断

## 132. Palindrome Partitioning II

在 [131](#131-palindrome-partitioning) 的判断字串回文的基础上，使用动态规划。若字串 [i, j] 为回文串，则 $minCut[j] = \min(minCut[i] + 1)$。枚举所有可能的前缀 i，时间复杂度为 O(n^2)。

## 224 & 227. Basic Calculator

- 224 只有 '+' 和 '-' 两种运算，针对正负两种情况设计符号栈，直接展开表达式。
- 227 只有四则运算，没有括号可以使用栈，先计算乘除，加减号将操作数入栈

显式后缀转中缀表达式，大量字符串操作容易超时，需要特别留意（尤其子字符串的传递），使用双栈（数字栈、符号栈）操作无需进行后缀转中缀的操作。

## 338. Counting Bits

- `x = x & (x−1)`，将 x 最低位的 1 变成 0
  - 若不返回数组，空间复杂度为 O(1)
- `bits[i] = bits[i − highBit] + 1`
- `bits[x] = bits[x >> 1] + (x & 1)`
- `bits[x] = bits[x & (x − 1)] + 1`

## 354. Russian Doll Envelopes

二维字典序严格递增最长序列 -> 最长递增子序列

问题转换：先对第一维进行字典序排序，再对最后一个维度求最长递增子序列

- 避免前面维数不是严格递增，最后一维使用倒序排序

## 503. Next Greater Element II

单调栈维护序列的单调性，两个循环实现循环队列的需求，也可以在第一个循环增加循环长度下标取模实现

## 1047. Remove All Adjacent Duplicates In String

使用栈（string 也可以作为容器）记录移出相邻重复的字符。主要看清题意，每次移出两个字符。
